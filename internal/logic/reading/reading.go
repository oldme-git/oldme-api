package reading

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/do"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

// Cre 创建阅读
func Cre(ctx context.Context, in *entity.Reading) (err error) {
	_, err = dao.Reading.Ctx(ctx).Data(do.Reading{
		Name:       in.Name,
		Author:     in.Author,
		Status:     in.Status,
		FinishedAt: in.FinishedAt,
	}).Insert()
	if err != nil {
		return utility.Err.Sys(err)
	}
	return err
}

// Upd 更新阅读
func Upd(ctx context.Context, in *entity.Reading) (err error) {
	_, err = dao.Reading.Ctx(ctx).Data(do.Reading{
		Name:       in.Name,
		Author:     in.Author,
		Status:     in.Status,
		FinishedAt: in.FinishedAt,
	}).Where("id", in.Id).Update()
	if err != nil {
		return utility.Err.Sys(err)
	}

	return
}

// List 读取列表
func List(ctx context.Context) (list []entity.Reading, err error) {
	res, err := dao.Reading.Ctx(ctx).
		Order("status desc").
		Order("finished_at desc").
		All()
	if err != nil {
		return nil, utility.Err.Sys(err)
	}
	_ = res.Structs(&list)
	return
}

// Show 读取阅读详情
func Show(ctx context.Context, id model.Id) (info *entity.Reading, err error) {
	info = &entity.Reading{}
	err = dao.Reading.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = utility.Err.TableNotData(err)
	}
	return
}

// Del 删除阅读
func Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Reading.Ctx(ctx).Where("id", id).Delete()
	return
}
