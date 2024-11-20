package tag

import (
	"context"

	"github.com/oldme-git/oldme-api/api/tag/v1"
	"github.com/oldme-git/oldme-api/internal/logic/tag"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = tag.Cre(ctx, req.GrpId, req.Name)
	return nil, err
}
