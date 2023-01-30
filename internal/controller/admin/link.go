package admin

import (
	"context"
	v1 "oldme-api/api/admin/v1"
	"oldme-api/internal/service"
)

var Link = cLink{}

type cLink struct {
}

func (c *cLink) Cre(ctx context.Context, req *v1.LinkCreReq) (res *v1.LinkCreRes, err error) {
	err = service.Link().Cre(ctx, req.LinkInput)
	return
}

func (c *cLink) Upd(ctx context.Context, req *v1.LinkUpdReq) (res *v1.LinkUpdRes, err error) {
	err = service.Link().Upd(ctx, req.Id, req.LinkInput)
	return

}

func (c *cLink) Del(ctx context.Context, req *v1.LinkDelReq) (res *v1.LinkDelRes, err error) {
	err = service.Link().Del(ctx, req.Id)
	return
}

func (c *cLink) List(ctx context.Context, req *v1.LinkListReq) (res *v1.LinkListRes, err error) {
	list, err := service.Link().List(ctx)
	if err == nil {
		res = &v1.LinkListRes{
			List:  list,
			Total: uint(len(*list)),
		}
	}
	return
}
