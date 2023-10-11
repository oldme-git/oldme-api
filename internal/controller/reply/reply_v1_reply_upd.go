package reply

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/reply/v1"
)

func (c *ControllerV1) ReplyUpd(ctx context.Context, req *v1.ReplyUpdReq) (res *v1.ReplyUpdRes, err error) {
	err = service.Reply().Upd(ctx, req.Id, req.ReplyBody)
	return
}
