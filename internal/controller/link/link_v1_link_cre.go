package link

import (
	"context"

	"github.com/oldme-git/oldme-api/api/link/v1"
	"github.com/oldme-git/oldme-api/internal/logic/link"
)

func (c *ControllerV1) LinkCre(ctx context.Context, req *v1.LinkCreReq) (res *v1.LinkCreRes, err error) {
	err = link.Cre(ctx, req.LinkInput)
	return
}
