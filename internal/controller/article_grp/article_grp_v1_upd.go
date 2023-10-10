package article_grp

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article_grp/v1"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = service.ArticleGrp().Upd(ctx, req.Id, req.ArticleGrpInput)
	return
}
