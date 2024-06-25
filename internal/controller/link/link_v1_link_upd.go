package link

import (
	"context"

	"github.com/oldme-git/oldme-api/api/link/v1"
	"github.com/oldme-git/oldme-api/internal/logic/link"
)

func (c *ControllerV1) LinkUpd(ctx context.Context, req *v1.LinkUpdReq) (res *v1.LinkUpdRes, err error) {
	err = link.Upd(ctx, req.Id, req.LinkInput)
	return
}
