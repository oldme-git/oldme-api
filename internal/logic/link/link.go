package article_grp

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/do"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/packed"
	"github.com/oldme-git/oldme-api/internal/service"
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
	if err != nil {
		err = packed.Err.SysDb("insert", "link")
	}
	return
}

// Upd 更新友链
func (s *sLink) Upd(ctx context.Context, id model.Id, in *model.LinkInput) (err error) {
	_, err = dao.Link.Ctx(ctx).Data(do.Link{
		Name:        in.Name,
		Description: in.Description,
		Link:        in.Link,
	}).Where("id", id).Update()
	if err != nil {
		err = packed.Err.SysDb("update", "link")
	}
	return
}

// Del 删除友链
func (s *sLink) Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Link.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		err = packed.Err.SysDb("delete", "link")
	}
	return
}

// List 读取友链列表
func (s *sLink) List(ctx context.Context) (list []entity.Link, err error) {
	res, err := dao.Link.Ctx(ctx).All()
	if err != nil {
		return nil, packed.Err.SysDb("select", "link")
	}
	_ = res.Structs(&list)
	return
}
