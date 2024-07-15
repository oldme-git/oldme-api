package reading

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

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
