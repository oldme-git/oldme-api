package link

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/link/v1"
)

func (c *ControllerV1) LinkDel(ctx context.Context, req *v1.LinkDelReq) (res *v1.LinkDelRes, err error) {
	err = service.Link().Del(ctx, req.Id)
	return
}
