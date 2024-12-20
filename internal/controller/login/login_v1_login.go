package login

import (
	"context"

	"github.com/oldme-git/oldme-api/api/login/v1"
	"github.com/oldme-git/oldme-api/internal/logic/account"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRes, error) {
	token, err := account.Login(ctx, &model.LoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{
		Token: token,
	}, nil
}
