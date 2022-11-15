package admin

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/service"
)

type sAdmin struct {
}

func init() {
	service.RegisterAdmin(&sAdmin{})
}

func (s sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (err error) {
	_, err = dao.Admin.Ctx(ctx).Data(do.Admin{
		Username:  in.Username,
		Password:  in.Password,
		LastLogin: gtime.Now(),
	}).Insert()
	return
}

func (s sAdmin) Update(ctx context.Context) {

}

func (s sAdmin) Read(ctx context.Context) {

}

func (s sAdmin) Del(ctx context.Context) {

}
