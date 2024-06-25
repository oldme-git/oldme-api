package reply

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reply/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reply"
)

func (c *ControllerV1) ReplyUpd(ctx context.Context, req *v1.ReplyUpdReq) (res *v1.ReplyUpdRes, err error) {
	err = reply.Upd(ctx, req.Id, req.ReplyBody)
	return
}
