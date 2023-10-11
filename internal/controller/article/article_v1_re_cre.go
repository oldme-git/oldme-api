package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article/v1"
)

func (c *ControllerV1) ReCre(ctx context.Context, req *v1.ReCreReq) (res *v1.ReCreRes, err error) {
	err = service.Article().ReCre(ctx, req.Id)
	return
}
