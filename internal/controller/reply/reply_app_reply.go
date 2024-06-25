package reply

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reply/app"
	"github.com/oldme-git/oldme-api/internal/logic/reply"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

func (c *ControllerApp) Reply(ctx context.Context, req *app.ReplyReq) (res *app.ReplyRes, err error) {
	_, err = reply.Cre(ctx, &entity.Reply{
		Aid:     int(req.Aid),
		Pid:     int(req.Pid),
		Email:   req.Email,
		Site:    req.Site,
		Name:    req.Name,
		Content: req.Content,
		Status:  model.CheckStatus,
	})
	return
}
