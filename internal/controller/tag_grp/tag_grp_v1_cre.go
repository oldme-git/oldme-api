package tag_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/tag_grp"

	"github.com/oldme-git/oldme-api/api/tag_grp/v1"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = tag_grp.Cre(ctx, req.Name)
	return nil, err
}
