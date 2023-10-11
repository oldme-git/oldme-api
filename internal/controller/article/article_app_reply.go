package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article/app"
)

func (c *ControllerApp) Reply(ctx context.Context, req *app.ReplyReq) (res *app.ReplyRes, err error) {
	total, list, err := service.Reply().ListForAid(ctx, req.Id)
	if err == nil {
		res = &app.ReplyRes{
			List:  list,
			Total: total,
		}
	}
	return
}
