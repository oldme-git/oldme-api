package reply

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/reply/v1"
)

func (c *ControllerV1) ReplyCre(ctx context.Context, req *v1.ReplyCreReq) (res *v1.ReplyCreRes, err error) {
	if req.Status == 0 {
		req.Status = model.SuccessStatus
	}
	_, err = service.Reply().Cre(ctx, &entity.Reply{
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
