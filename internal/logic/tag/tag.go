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
