package controller

import (
	"context"
	v1 "oldme-api/api/v1"
	"oldme-api/internal/model"
	"oldme-api/internal/service"
)

var Account = &cAccount{}

type cAccount struct {
}

// Login 登录
func (c cAccount) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := service.Account().Login(ctx, model.Login{
		Username: req.Username,
		Password: req.Password,
	})
	if err == nil {
		res = &v1.LoginRes{
			Token: token,
		}
	}
	return
}

// Logout 退出登录
func (c cAccount) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	err = service.Account().Logout(ctx)
	return
}

// Info 获取账户信息
func (c cAccount) Info(ctx context.Context, req *v1.AccountReq) (res *v1.AccountRes, err error) {
	admin, err := service.Account().Info(ctx)
	if err == nil {
		res = &v1.AccountRes{
			Username:  admin.Username,
			Nickname:  admin.Nickname,
			Avatar:    admin.Avatar,
			Register:  admin.Register,
			LastLogin: admin.LastLogin,
		}
	}
	return
}
