package reply

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reply/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reply"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

func (c *ControllerV1) ReplyCre(ctx context.Context, req *v1.ReplyCreReq) (res *v1.ReplyCreRes, err error) {
	if req.Status == 0 {
		req.Status = model.SuccessStatus
	}
	_, err = reply.Cre(ctx, &entity.Reply{
		Aid:     int(req.Aid),
		Pid:     int(req.Pid),
		Email:   req.Email,
		Site:    req.Site,
		Name:    req.Name,
		Content: req.Content,
		Status:  int(req.Status),
	})
	return
}
