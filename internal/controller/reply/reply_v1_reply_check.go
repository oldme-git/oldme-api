package reply

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reply/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reply"
)

func (c *ControllerV1) ReplyCheck(ctx context.Context, req *v1.ReplyCheckReq) (res *v1.ReplyCheckRes, err error) {
	err = reply.Check(ctx, req.Id, req.Result, req.Reason)
	return
}
