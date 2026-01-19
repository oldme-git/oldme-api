package reading

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reading/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reading"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = reading.Upd(ctx, &entity.Reading{
		Id:         uint(req.Id),
		Name:       req.Name,
		Author:     req.Author,
		Status:     uint(req.Status),
		FinishedAt: req.FinishedAt,
	})
	return nil, err
}
