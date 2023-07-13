package reply

import (
	"context"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/packed"
	"oldme-api/internal/service"
	"strings"
	"sync"
)

type sReply struct {
}

func init() {
	service.RegisterReply(&sReply{})
}

// Cre 新增回复
func (s *sReply) Cre(ctx context.Context, in *entity.Reply) (lastId model.Id, err error) {
	if in.Site != "" {
		if !strings.HasPrefix(in.Site, "http") {
			return 0, packed.Err.Skip(10301)
		}
	}

	// 植入aid
	aid, err := s.GetAid(ctx, model.Id(in.Pid))
	if err != nil {
		return
	}
	if aid != 0 {
		in.Aid = int(aid)
	}

	// 判断该文章是否存在
	if ok := service.Article().IsExist(ctx, model.Id(in.Aid)); !ok {
		return 0, packed.Err.Skip(10201)
	}

	// 植入pName
	if in.Pid != 0 {
		pName, _ := dao.Reply.Ctx(ctx).Where("id", in.Pid).Fields("name").Value()
		in.PName = pName.String()
	}

	// 植入根id
	rid, err := s.GetRid(ctx, model.Id(in.Pid))
	if err != nil {
		return
	}
	in.Rid = int(rid)

	res, err := dao.Reply.Ctx(ctx).Data(in).Insert()
	if err != nil {
		return 0, packed.Err.SysDb("insert", "reply")
	}
	resId, _ := res.LastInsertId()
	return model.Id(resId), nil
}

// Upd 更新文章回复
func (s *sReply) Upd(ctx context.Context, id model.Id, in *model.ReplyBody) (err error) {
	if in.Site != "" {
		if !strings.HasPrefix(in.Site, "http") {
			return packed.Err.Skip(10301)
		}
	}
	_, err = dao.Reply.Ctx(ctx).Data(do.Reply{
		Email:   in.Email,
		Name:    in.Name,
		Site:    in.Site,
		Content: in.Content,
	}).Where("id", id).Update()
	if err != nil {
		return packed.Err.SysDb("update", "reply")
	}
	return
}

// Del 删除文章回复
func (s *sReply) Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Reply.Ctx(ctx).Where("id", id).Delete()
	return
}

// List 读取文章回复列表
func (s *sReply) List(ctx context.Context, query *model.ReplyQuery) (list *[]entity.Reply, total uint, err error) {
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

	data, err := db.All()
	if err != nil {
		err = packed.Err.SysDb("select", "relay")
		return
	}
	list = &[]entity.Reply{}
	_ = data.Structs(list)
	// 查询总数据条数
	totalInt, err := db.Count()
	if err != nil {
		err = packed.Err.SysDb("select", "relay")
		return
	}
	total = uint(totalInt)
	return
}

// Show 读取文章回复详情，携带父级信息
func (s *sReply) Show(ctx context.Context, id model.Id) (info *model.ReplyShow, err error) {
	info = &model.ReplyShow{}
	err = dao.Reply.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = packed.Err.Skip(10100)
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
func (s *sReply) Details(ctx context.Context, id model.Id) (info *entity.Reply, err error) {
	info = &entity.Reply{}
	err = dao.Reply.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = packed.Err.Skip(10100)
	}
	return
}

// Check 审核
func (s *sReply) Check(ctx context.Context, id model.Id, result bool, reasonSlice ...string) error {
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
		return packed.Err.SysDb("update", "reply")
	}
	return err
}

// GetRid 根据父id获取根id，如果父回复已经是根回复了，则rid=0，否则rid=pid
func (s *sReply) GetRid(ctx context.Context, pid model.Id) (model.Id, error) {
	var rid model.Id
	if pid == 0 {
		rid = 0
	} else {
		details, err := s.Details(ctx, pid)
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
func (s *sReply) GetAid(ctx context.Context, pid model.Id) (model.Id, error) {
	var aid model.Id
	if pid != 0 {
		parent := &entity.Reply{}
		err := dao.Reply.Ctx(ctx).Where("id", pid).Where("status", model.SuccessStatus).Scan(parent)
		if err != nil {
			return 0, packed.Err.Skip(10303)
		}
		aid = model.Id(parent.Aid)
	}
	return aid, nil
}

// ListForAid 根据Aid读取回复列表
func (s *sReply) ListForAid(ctx context.Context, aid model.Id) ([]model.ReplyFloorApp, error) {
	data, err := dao.Reply.Ctx(ctx).Where("aid", aid).All()
	if err != nil {
		return nil, packed.Err.SysDb("select", "reply")
	}
	var (
		list      []entity.Reply
		listFloor []model.ReplyFloorApp
		wg        = &sync.WaitGroup{}
	)
	_ = data.Structs(&list)

	listFloor = s.GetReplyFloor(list, 0)
	wg.Add(len(listFloor))

	for k, v := range listFloor {
		go func(k int, v model.ReplyFloorApp) {
			listFloor[k].List = s.GetReplyFloor(list, v.Id)
			wg.Done()
		}(k, v)
	}
	wg.Wait()

	return listFloor, nil
}

// GetReplyFloor 根据rid读取本层楼的所有回复信息，id
func (s *sReply) GetReplyFloor(list []entity.Reply, rid model.Id) (reply []model.ReplyFloorApp) {
	for _, v := range list {
		if v.Rid == int(rid) {
			reply = append(reply, model.ReplyFloorApp{
				Id:        model.Id(v.Id),
				Pid:       model.Id(v.Pid),
				Email:     v.Email,
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
