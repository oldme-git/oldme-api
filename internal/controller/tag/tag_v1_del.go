package tag

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/tag"

	"github.com/oldme-git/oldme-api/api/tag/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = tag.Del(ctx, req.Id)
	return nil, err
}
