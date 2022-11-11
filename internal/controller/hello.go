package controller

import (
	"context"
	"oldme-api/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	res = &v1.HelloRes{
		Name: "我的名字",
	}
	return
}
