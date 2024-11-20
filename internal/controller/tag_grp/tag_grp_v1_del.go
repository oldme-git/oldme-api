package tag_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/tag_grp"

	"github.com/oldme-git/oldme-api/api/tag_grp/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = tag_grp.Del(ctx, req.Id)
	return nil, err
}
