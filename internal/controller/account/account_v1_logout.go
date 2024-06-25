package account

import (
	"context"

	"github.com/oldme-git/oldme-api/api/account/v1"
	"github.com/oldme-git/oldme-api/internal/logic/account"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	err = account.Logout(ctx)
	return
}
