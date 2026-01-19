package reading

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reading/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reading"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	info, err := reading.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ShowRes{
		Reading: info,
	}, nil
}
