package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = article.Del(ctx, req.Id, req.IsReal)
	return
}
