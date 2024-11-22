package tag

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/tag"

	"github.com/oldme-git/oldme-api/api/tag/v1"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	info, err := tag.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ShowRes{
		Tag: info,
	}, nil
}
