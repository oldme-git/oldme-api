package reply

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/reply/v1"
)

func (c *ControllerV1) ReplyShow(ctx context.Context, req *v1.ReplyShowReq) (res *v1.ReplyShowRes, err error) {
	info, err := service.Reply().Show(ctx, req.Id)
	if err == nil {
		res = &v1.ReplyShowRes{
			ReplyShow: info,
		}
	}
	return
}
