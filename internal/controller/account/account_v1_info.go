package account

import (
	"context"

	"github.com/oldme-git/oldme-api/api/account/v1"
	"github.com/oldme-git/oldme-api/internal/logic/account"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	admin, err := account.Info(ctx)
	if err == nil {
		res = &v1.InfoRes{
			Username:  admin.Username,
			Nickname:  admin.Nickname,
			Avatar:    admin.Avatar,
			Register:  admin.Register,
			LastLogin: admin.LastLogin,
		}
	}
	return
}
