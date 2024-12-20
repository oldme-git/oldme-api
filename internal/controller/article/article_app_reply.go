package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/app"
	"github.com/oldme-git/oldme-api/internal/logic/reply"
)

func (c *ControllerApp) Reply(ctx context.Context, req *app.ReplyReq) (res *app.ReplyRes, err error) {
	total, list, err := reply.ListForAid(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &app.ReplyRes{
		List:  list,
		Total: total,
	}, nil
}
