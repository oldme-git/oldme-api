package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article/v1"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	LastId, err := service.Article().Cre(ctx, req.ArticleInput)
	if err == nil {
		res = &v1.CreRes{LastId: LastId}
	}
	return
}
