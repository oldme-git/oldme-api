package link

import (
	"context"

	"github.com/oldme-git/oldme-api/api/link/v1"
	"github.com/oldme-git/oldme-api/internal/logic/link"
)

func (c *ControllerV1) LinkDel(ctx context.Context, req *v1.LinkDelReq) (res *v1.LinkDelRes, err error) {
	err = link.Del(ctx, req.Id)
	return
}
