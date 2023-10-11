package saying

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/saying/v1"
)

func (c *ControllerV1) SayingUpd(ctx context.Context, req *v1.SayingUpdReq) (res *v1.SayingUpdRes, err error) {
	err = service.Saying().Upd(ctx, req.Id, req.Saying)
	return
}
