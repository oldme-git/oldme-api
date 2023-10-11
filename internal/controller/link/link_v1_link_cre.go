package link

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/link/v1"
)

func (c *ControllerV1) LinkCre(ctx context.Context, req *v1.LinkCreReq) (res *v1.LinkCreRes, err error) {
	err = service.Link().Cre(ctx, req.LinkInput)
	return
}
