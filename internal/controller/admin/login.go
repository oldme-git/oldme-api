package admin

import (
	"context"
	"github.com/oldme-git/oldme-api/api/admin/v1"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/service"
)

var Login = &cLogin{}

type cLogin struct {
}

// Login 登录
func (c cLogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := service.Account().Login(ctx, &model.Login{
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
