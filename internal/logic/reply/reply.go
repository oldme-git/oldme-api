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
		err = packed.Err.Skip(10201)
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
func (s *sReply) Upd(ctx context.Context, id model.Id, in *model.ReplyInput) (err error) {
	// 判断该文章是否存在
	if ok := service.Article().IsExist(ctx, in.Aid); !ok {
		err = packed.Err.Skip(10101)
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
	}).Where("id", id).Update()
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
	db = db.Order("created_at desc, id desc").Page(query.Page, query.Size)

	data, err := db.All()
	if err != nil {
		return
	}
	list = &[]entity.Reply{}
	_ = data.Structs(list)
	// 查询总数据条数
	totalInt, err := db.Count()
	if err != nil {
		return
	}
	total = uint(totalInt)
	return
}

// Show 读取文章回复详情
func (s *sReply) Show(ctx context.Context, id model.Id) (info *entity.Reply, err error) {
	err = dao.Reply.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = packed.Err.Skip(10100)
	}
	return
}
