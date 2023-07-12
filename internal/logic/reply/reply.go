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
)

type sReply struct {
}

func init() {
	service.RegisterReply(&sReply{})
}

// Cre 新增回复
func (s *sReply) Cre(ctx context.Context, in *model.ReplyInput) (err error) {
	if !strings.HasPrefix(in.Site, "http") {
		return packed.Err.Skip(10301)
	}
	// 判断该文章是否存在
	if ok := service.Article().IsExist(ctx, in.Aid); !ok {
		return packed.Err.Skip(10201)
	}

	// 判断回复的父级内容是否属于该文章
	if err = replyCheck(ctx, in.Aid, in.Pid); err != nil {
		return
	}

	_, err = dao.Reply.Ctx(ctx).Data(do.Reply{
		Aid:     in.Aid,
		Pid:     in.Pid,
		Email:   in.Email,
		Name:    in.Name,
		Site:    in.Site,
		Content: in.Content,
		Status:  in.Status,
		Reason:  in.Reason,
	}).Insert()
	return
}

// Upd 更新文章回复
func (s *sReply) Upd(ctx context.Context, id model.Id, in *model.ReplyBody) (err error) {
	if !strings.HasPrefix(in.Site, "http") {
		return packed.Err.Skip(10301)
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

// Show 读取文章回复详情
func (s *sReply) Show(ctx context.Context, id model.Id) (info *model.ReplyShow, err error) {
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

// replyCheck 判断回复内容是否在正确的文章下
// aid 回复的文章id
// pid 回复的父级回复id
func replyCheck(ctx context.Context, aid model.Id, pid model.Id) error {
	// 判断回复的父级内容是否属于该文章
	if pid != 0 {
		parent := &entity.Reply{}
		err := dao.Reply.Ctx(ctx).Where("id", pid).Where("status", model.SuccessStatus).Scan(parent)
		if err != nil {
			return packed.Err.Skip(10303)
		}
		if parent.Aid != int(aid) {
			return packed.Err.Skip(10302)
		}
	}
	return nil
}
