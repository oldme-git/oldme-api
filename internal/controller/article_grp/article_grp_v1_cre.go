package article_grp

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article_grp/v1"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = service.ArticleGrp().Cre(ctx, req.ArticleGrpInput)
	return
}
