package tag

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/do"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

// Cre 创建标签
func Cre(ctx context.Context, grpId model.Id, name string) (err error) {
	if err := check(ctx, 0, name); err != nil {
		return err
	}
	_, err = dao.Tag.Ctx(ctx).Data(do.Tag{
		GrpId: grpId,
		Name:  name,
	}).Insert()
	if err != nil {
		return utility.Err.Sys(err)
	}
	return err
}

// Upd 更新标签
func Upd(ctx context.Context, id, grpId model.Id, name string) (err error) {
	if err := check(ctx, id, name); err != nil {
		return err
	}

	_, err = dao.Tag.Ctx(ctx).Data(do.Tag{
		GrpId: grpId,
		Name:  name,
	}).Where("id", id).Update()
	if err != nil {
		return utility.Err.Sys(err)
	}

	return
}

// List 读取标签列表
func List(ctx context.Context, grpId model.Id) (list []entity.Tag, err error) {
	list = make([]entity.Tag, 0)
	err = dao.Tag.Ctx(ctx).Where("grp_id", grpId).Scan(&list)
	return
}

// Del 删除标签
func Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Tag.Ctx(ctx).Where("id", id).Delete()
	return
}

// Show 读取详情
func Show(ctx context.Context, id model.Id) (info *entity.Tag, err error) {
	err = dao.Tag.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = utility.Err.Skip(10504)
	}
	return
}

// check 检查标签是否存在
func check(ctx context.Context, id model.Id, name string) error {
	db := dao.Tag.Ctx(ctx).Where("name", name)
	if id > 0 {
		db = db.WhereNot("id", id)
	}
	count, err := db.Count()
	if err != nil {
		return utility.Err.Sys(err)
	}
	if count > 0 {
		return utility.Err.Skip(10601)
	}
	return nil
}
