package other

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/other/app"
)

func (c *ControllerApp) Saying(ctx context.Context, req *app.SayingReq) (res *app.SayingRes, err error) {
	saying, err := service.Saying().Show(ctx)
	if err == nil {
		res = &app.SayingRes{Saying: saying}
	}
	return
}
