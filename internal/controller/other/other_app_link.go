package other

import (
	"context"

	"github.com/oldme-git/oldme-api/api/other/app"
	"github.com/oldme-git/oldme-api/internal/logic/link"
)

func (c *ControllerApp) Link(ctx context.Context, req *app.LinkReq) (res *app.LinkRes, err error) {
	list, err := link.List(ctx)
	if err == nil {
		res = &app.LinkRes{
			List:  list,
			Total: uint(len(list)),
		}
	}
	return
}
