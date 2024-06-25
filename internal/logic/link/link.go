package link

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/do"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

// Cre 创建友链
func Cre(ctx context.Context, in *model.LinkInput) (err error) {
	_, err = dao.Link.Ctx(ctx).Data(do.Link{
		Name:        in.Name,
		Description: in.Description,
		Link:        in.Link,
	}).Insert()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// Upd 更新友链
func Upd(ctx context.Context, id model.Id, in *model.LinkInput) (err error) {
	_, err = dao.Link.Ctx(ctx).Data(do.Link{
		Name:        in.Name,
		Description: in.Description,
		Link:        in.Link,
	}).Where("id", id).Update()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// Del 删除友链
func Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Link.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// List 读取友链列表
func List(ctx context.Context) (list []entity.Link, err error) {
	res, err := dao.Link.Ctx(ctx).All()
	if err != nil {
		return nil, utility.Err.Sys(err)
	}
	_ = res.Structs(&list)
	return
}
