package saying

import (
	"context"

	"github.com/oldme-git/oldme-api/api/saying/v1"
	"github.com/oldme-git/oldme-api/internal/logic/saying"
)

func (c *ControllerV1) SayingUpd(ctx context.Context, req *v1.SayingUpdReq) (res *v1.SayingUpdRes, err error) {
	err = saying.Upd(ctx, req.Id, req.Saying)
	return
}
