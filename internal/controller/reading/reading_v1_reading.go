package reading

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/reading"

	"github.com/oldme-git/oldme-api/api/reading/v1"
)

func (c *ControllerV1) Reading(ctx context.Context, req *v1.ReadingReq) (res *v1.ReadingRes, err error) {
	list, err := reading.List(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ReadingRes{
		List: list,
	}, nil
}
