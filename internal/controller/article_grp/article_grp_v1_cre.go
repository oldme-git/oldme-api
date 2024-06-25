package article_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article_grp/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article_grp"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = article_grp.Cre(ctx, req.ArticleGrpInput)
	return
}
