package tag_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/do"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

// Cre 创建标签分类
func Cre(ctx context.Context, name string) (err error) {
	_, err = dao.TagGrp.Ctx(ctx).Data(do.TagGrp{
		Name: name,
	}).Insert()
	if err != nil {
		return utility.Err.Sys(err)
	}
	return
}

// Upd 更新标签分类
func Upd(ctx context.Context, id model.Id, name string) (err error) {
	_, err = dao.TagGrp.Ctx(ctx).Data(do.TagGrp{
		Name: name,
	}).Where("id", id).Update()
	if err != nil {
		return utility.Err.Sys(err)
	}
	return
}

// Del 删除标签分类
func Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.TagGrp.Ctx(ctx).Where("id", id).Delete()
	return
}

// List 读取标签分类列表
func List(ctx context.Context) (list []entity.TagGrp, err error) {
	list = make([]entity.TagGrp, 0)
	err = dao.TagGrp.Ctx(ctx).Scan(&list)
	return
}

// Show 读取详情
func Show(ctx context.Context, id model.Id) (info *entity.TagGrp, err error) {
	err = dao.TagGrp.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = utility.Err.Skip(10504)
	}
	return
}
