package reading

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reading/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reading"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = reading.Cre(ctx, &entity.Reading{
		Name:       req.Name,
		Author:     req.Author,
		Status:     uint(req.Status),
		FinishedAt: req.FinishedAt,
	})
	return nil, err
}
