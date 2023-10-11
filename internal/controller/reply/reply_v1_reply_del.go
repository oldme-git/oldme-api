package reply

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/reply/v1"
)

func (c *ControllerV1) ReplyDel(ctx context.Context, req *v1.ReplyDelReq) (res *v1.ReplyDelRes, err error) {
	err = service.Reply().Del(ctx, req.Id)
	return
}
