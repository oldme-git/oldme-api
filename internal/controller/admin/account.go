package admin

import (
	"context"
	"github.com/oldme-git/oldme-api/api/admin/v1"
	"github.com/oldme-git/oldme-api/internal/service"
)

var Account = &cAccount{}

type cAccount struct {
}

// Info 获取账户信息
func (c cAccount) Info(ctx context.Context, req *v1.AccountInfoReq) (res *v1.AccountInfoRes, err error) {
	admin, err := service.Account().Info(ctx)
	if err == nil {
		res = &v1.AccountInfoRes{
			Username:  admin.Username,
			Nickname:  admin.Nickname,
			Avatar:    admin.Avatar,
			Register:  admin.Register,
			LastLogin: admin.LastLogin,
		}
	}
	return
}

// Logout 退出登录
func (c cAccount) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	err = service.Account().Logout(ctx)
	return
}
