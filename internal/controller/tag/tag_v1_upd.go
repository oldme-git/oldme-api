package tag

import (
	"context"

	"github.com/oldme-git/oldme-api/api/tag/v1"
	"github.com/oldme-git/oldme-api/internal/logic/tag"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = tag.Upd(ctx, req.Id, req.GrpId, req.Name)
	return nil, err
}
