package saying

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/saying/v1"
)

func (c *ControllerV1) SayingDel(ctx context.Context, req *v1.SayingDelReq) (res *v1.SayingDelRes, err error) {
	err = service.Saying().Del(ctx, req.Id)
	return
}
