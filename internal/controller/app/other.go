package app

import (
	"context"
	v1 "oldme-api/api/app/v1"
	"oldme-api/internal/service"
)

var Other = cOther{}

type cOther struct {
}

// List 查询友链列表
func (c *cOther) List(ctx context.Context, req *v1.LinkReq) (res *v1.LinkRes, err error) {
	list, err := service.Link().List(ctx)
	if err == nil {
		res = &v1.LinkRes{
			List:  list,
			Total: uint(len(*list)),
		}
	}
	return
}

// Saying 查询一条随机句子
func (c *cOther) Saying(ctx context.Context, req *v1.SayingReq) (res *v1.SayingRes, err error) {
	saying, err := service.Saying().Show(ctx)
	if err == nil {
		res = &v1.SayingRes{Saying: saying}
	}
	return
}
