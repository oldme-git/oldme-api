package admin

import (
	"context"
	v1 "github.com/oldme-git/oldme-api/api/admin/v1"
	"github.com/oldme-git/oldme-api/internal/service"
)

var Saying = cSaying{}

type cSaying struct {
}

func (c *cSaying) Cre(ctx context.Context, req *v1.SayingCreReq) (res *v1.SayingCreRes, err error) {
	err = service.Saying().Cre(ctx, req.Saying)
	return
}

func (c *cSaying) Upd(ctx context.Context, req *v1.SayingUpdReq) (res *v1.SayingUpdRes, err error) {
	err = service.Saying().Upd(ctx, req.Id, req.Saying)
	return

}

func (c *cSaying) Del(ctx context.Context, req *v1.SayingDelReq) (res *v1.SayingDelRes, err error) {
	err = service.Saying().Del(ctx, req.Id)
	return
}

func (c *cSaying) List(ctx context.Context, req *v1.SayingListReq) (res *v1.SayingListRes, err error) {
	list, err := service.Saying().List(ctx)
	if err == nil {
		res = &v1.SayingListRes{
			List:  list,
			Total: uint(len(list)),
		}
	}
	return
}
