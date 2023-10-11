package saying

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/saying/v1"
)

func (c *ControllerV1) SayingCre(ctx context.Context, req *v1.SayingCreReq) (res *v1.SayingCreRes, err error) {
	err = service.Saying().Cre(ctx, req.Saying)
	return
}
