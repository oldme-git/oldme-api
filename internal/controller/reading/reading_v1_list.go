package reading

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reading/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reading"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, err := reading.List(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListRes{
		List: list,
	}, nil
}
