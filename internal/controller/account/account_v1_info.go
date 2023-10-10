package account

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/account/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	admin, err := service.Account().Info(ctx)
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
