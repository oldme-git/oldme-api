package login

import (
	"context"

	"github.com/oldme-git/oldme-api/api/login/v1"
	"github.com/oldme-git/oldme-api/internal/logic/account"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := account.Login(ctx, &model.Login{
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
