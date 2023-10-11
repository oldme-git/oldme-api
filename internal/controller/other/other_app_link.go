package other

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/other/app"
)

func (c *ControllerApp) Link(ctx context.Context, req *app.LinkReq) (res *app.LinkRes, err error) {
	list, err := service.Link().List(ctx)
	if err == nil {
		res = &app.LinkRes{
			List:  list,
			Total: uint(len(list)),
		}
	}
	return
}
