package tag_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/tag_grp"

	"github.com/oldme-git/oldme-api/api/tag_grp/v1"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = tag_grp.Upd(ctx, req.Id, req.Name)
	return nil, err
}
