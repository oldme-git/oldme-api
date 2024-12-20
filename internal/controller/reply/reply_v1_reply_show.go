package reply

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reply/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reply"
)

func (c *ControllerV1) ReplyShow(ctx context.Context, req *v1.ReplyShowReq) (*v1.ReplyShowRes, error) {
	info, err := reply.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ReplyShowRes{
		ReplyShow: info,
	}, nil
}
