package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "oldme-api/api/v1"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/service"
)

var Demo = cDemo{}

type cDemo struct {
}

func (c *cDemo) Create(ctx context.Context, req *v1.DemoCreateReq) (res *v1.DemoCreateRes, err error) {
	err = service.Admin().Create(ctx, model.AdminInput{
		Username: req.Username,
		Password: req.Password,
	})
	return
}

func (c *cDemo) Update(ctx context.Context, req *v1.DemoUpdateReq) (res *v1.DemoUpdateRes, err error) {
	_, err = dao.Admin.Ctx(ctx).Where("username", "admin").Data(do.Admin{
		Password:  "456",
		LastLogin: gtime.Now(),
	}).Update()
	return
}

func (c *cDemo) Read(ctx context.Context, req *v1.DemoReadReq) (res *v1.DemoReadRes, err error) {
	var admin *entity.Admin
	err = dao.Admin.Ctx(ctx).Where("username", req.Username).Scan(&admin)
	if err != nil {
		return nil, err
	}
	res = &v1.DemoReadRes{Admin: admin}
	return
}

func (c *cDemo) Del(ctx context.Context, req *v1.DemoDelReq) (res *v1.DemoDelRes, err error) {
	_, err = dao.Admin.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}
