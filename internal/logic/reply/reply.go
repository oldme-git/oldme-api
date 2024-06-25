package reply

import (
	"context"
	"strings"
	"sync"

	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/do"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

// Cre 新增回复
func Cre(ctx context.Context, in *entity.Reply) (lastId model.Id, err error) {
	if in.Site != "" {
		if !strings.HasPrefix(in.Site, "http") {
			return 0, utility.Err.Skip(10301)
		}
	}

	// 植入aid
	aid, err := GetAid(ctx, model.Id(in.Pid))
	if err != nil {
		return
	}
	if aid != 0 {
		in.Aid = int(aid)
	}

	// 判断该文章是否存在
	if ok := article.IsExist(ctx, model.Id(in.Aid)); !ok {
		return 0, utility.Err.Skip(10201)
	}

	// 植入pName
	if in.Pid != 0 {
		pName, _ := dao.Reply.Ctx(ctx).Where("id", in.Pid).Fields("name").Value()
		in.PName = pName.String()
	}

	// 植入根id
	rid, err := GetRid(ctx, model.Id(in.Pid))
	if err != nil {
		return
	}
	in.Rid = int(rid)

	res, err := dao.Reply.Ctx(ctx).Data(in).Insert()
	if err != nil {
		return 0, utility.Err.Sys(err)
	}
	resId, _ := res.LastInsertId()
	return model.Id(resId), nil
}

// Upd 更新文章回复
func Upd(ctx context.Context, id model.Id, in *model.ReplyBody) (err error) {
	if in.Site != "" {
		if !strings.HasPrefix(in.Site, "http") {
			return utility.Err.Skip(10301)
		}
	}
	_, err = dao.Reply.Ctx(ctx).Data(do.Reply{
		Email:   in.Email,
		Name:    in.Name,
		Site:    in.Site,
		Content: in.Content,
	}).Where("id", id).Update()
	if err != nil {
		return utility.Err.Sys(err)
	}
	return
}

// Del 删除文章回复
func Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Reply.Ctx(ctx).Where("id", id).Delete()
	return
}

// List 读取文章回复列表
func List(ctx context.Context, query *model.ReplyQuery) (list []entity.Reply, total uint, err error) {
	if query == nil {
		query = &model.ReplyQuery{}
	}
	// 对于查询初始值的处理
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Size == 0 {
		query.Size = 15
	}
	// 组成查询链
	db := dao.Reply.Ctx(ctx)
	// 搜索文本
	if len(query.Search) != 0 {
		db = db.Where(db.Builder().Where("name like ?", "%"+query.Search+"%").
			WhereOr("email like ?", "%"+query.Search+"%").
			WhereOr("content like ?", "%"+query.Search+"%"))
	}
	if query.Status != 0 {
		db = db.Where("status", query.Status)
	}
	if query.Aid != 0 {
		db = db.Where("aid", query.Aid)
	}
	db = db.Order("created_at desc, id desc").Page(query.Page, query.Size)

	data, totalInt, err := db.AllAndCount(true)
	if err != nil {
		err = utility.Err.Sys(err)
		return
	}
	list = []entity.Reply{}
	_ = data.Structs(&list)

	total = uint(totalInt)

	return
}

// Show 读取文章回复详情，携带父级信息
func Show(ctx context.Context, id model.Id) (info *model.ReplyShow, err error) {
	info = &model.ReplyShow{}
	err = dao.Reply.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = utility.Err.Skip(10100)
	}

	// 读取所属文章名称
	articleTitle, _ := dao.Article.Ctx(ctx).Where("id", info.Aid).Fields("title").Value()
	info.ArticleTitle = articleTitle.String()

	// 读取父级回复内容
	if info.Pid != 0 {
		parentReply := &entity.Reply{}
		_ = dao.Reply.Ctx(ctx).Where("id", info.Pid).Scan(parentReply)
		info.ParentReply = *parentReply
	}
	return
}

// Details 读取文章回复详情
func Details(ctx context.Context, id model.Id) (info *entity.Reply, err error) {
	info = &entity.Reply{}
	err = dao.Reply.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = utility.Err.Skip(10100)
	}
	return
}

// Check 审核
func Check(ctx context.Context, id model.Id, result bool, reasonSlice ...string) error {
	var (
		status model.ReplyStatus
		reason string
	)
	if result {
		status = model.SuccessStatus
	} else {
		status = model.FailStatus
	}

	if len(reasonSlice) > 0 {
		reason = reasonSlice[0]
	}

	_, err := dao.Reply.Ctx(ctx).Data(do.Reply{
		Status: status,
		Reason: reason,
	}).Where("id", id).Update()
	if err != nil {
		return utility.Err.Sys(err)
	}
	return err
}

// GetRid 根据父id获取根id，如果父回复已经是根回复了，则rid=0，否则rid=pid
func GetRid(ctx context.Context, pid model.Id) (model.Id, error) {
	var rid model.Id
	if pid == 0 {
		rid = 0
	} else {
		details, err := Details(ctx, pid)
		if err != nil {
			return 0, err
		}
		if details.Rid != 0 {
			rid = model.Id(details.Rid)
		} else {
			rid = model.Id(details.Id)
		}
	}
	return rid, nil
}

// GetAid 根据父id获取aid，如果父回复已经是根回复了，则aid返回0
func GetAid(ctx context.Context, pid model.Id) (model.Id, error) {
	var aid model.Id
	if pid != 0 {
		parent := &entity.Reply{}
		err := dao.Reply.Ctx(ctx).Where("id", pid).Where("status", model.SuccessStatus).Scan(parent)
		if err != nil {
			return 0, utility.Err.Skip(10303)
		}
		aid = model.Id(parent.Aid)
	}
	return aid, nil
}

// ListForAid 根据Aid读取回复列表
func ListForAid(ctx context.Context, aid model.Id) (uint, []model.ReplyFloorApp, error) {
	data, err := dao.Reply.Ctx(ctx).Where("aid", aid).Where("status", model.SuccessStatus).All()
	if err != nil {
		return 0, nil, utility.Err.Sys(err)
	}
	var (
		list      []entity.Reply
		listFloor []model.ReplyFloorApp
		wg        = &sync.WaitGroup{}
		lock      = &sync.RWMutex{}
		total     = 0
	)
	_ = data.Structs(&list)

	listFloor = GetReplyFloor(list, 0)
	listLen := len(listFloor)
	total += listLen
	wg.Add(listLen)

	for k, v := range listFloor {
		go func(k int, v model.ReplyFloorApp) {
			listFloor[k].List = GetReplyFloor(list, v.Id)
			lock.Lock()
			total += len(listFloor[k].List)
			lock.Unlock()
			wg.Done()
		}(k, v)
	}
	wg.Wait()

	return uint(total), listFloor, nil
}

// GetReplyFloor 根据rid读取本层楼的所有回复信息，id
func GetReplyFloor(list []entity.Reply, rid model.Id) (reply []model.ReplyFloorApp) {
	for _, v := range list {
		if v.Rid == int(rid) {
			reply = append(reply, model.ReplyFloorApp{
				Id:        model.Id(v.Id),
				Pid:       model.Id(v.Pid),
				Email:     v.Email,
				PName:     v.PName,
				Name:      v.Name,
				Site:      v.Site,
				Content:   v.Content,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}
	}
	return reply
}
