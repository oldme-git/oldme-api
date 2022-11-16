package controller

import (
	"context"
	v1 "oldme-api/api/v1"
	"oldme-api/internal/model"
	"oldme-api/internal/service"
)

var Admin = &cAdmin{}

type cAdmin struct {
}

func (c cAdmin) Login(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	result, _ := service.Admin().Login(ctx, model.Login{
		Username: req.Username,
		Password: req.Password,
	})
	res = &v1.AdminLoginRes{
		Result: result,
	}
	err = nil
	return
}
