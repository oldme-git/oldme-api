package app

import (
	"context"
	v1 "oldme-api/api/app/v1"
	"oldme-api/internal/model"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/service"
)

var Reply = &cReply{}

type cReply struct {
}

func (c *cReply) Reply(ctx context.Context, req *v1.ReplyReq) (res *v1.ReplyRes, err error) {
	_, err = service.Reply().Cre(ctx, &entity.Reply{
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
