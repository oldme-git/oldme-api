package article_grp

import (
	"context"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/service"
)

type sLink struct {
}

func init() {
	service.RegisterLink(&sLink{})
}

// Cre 创建友链
func (s *sLink) Cre(ctx context.Context, in *model.LinkInput) (err error) {
	_, err = dao.Link.Ctx(ctx).Data(do.Link{
		Name:        in.Name,
		Description: in.Description,
		Link:        in.Link,
	}).Insert()
	return
}

// Upd 更新友链
func (s *sLink) Upd(ctx context.Context, id uint, in *model.LinkInput) (err error) {
	_, err = dao.Link.Ctx(ctx).Data(do.Link{
		Name:        in.Name,
		Description: in.Description,
		Link:        in.Link,
	}).Where("id", id).Update()
	return
}

// Del 删除友链
func (s *sLink) Del(ctx context.Context, id uint) (err error) {
	_, err = dao.Link.Ctx(ctx).Where("id", id).Delete()
	return
}

// List 读取友链列表
func (s *sLink) List(ctx context.Context) (list *[]entity.Link, err error) {
	list = &[]entity.Link{}
	res, err := dao.Link.Ctx(ctx).All()
	_ = res.Structs(list)
	return
}
