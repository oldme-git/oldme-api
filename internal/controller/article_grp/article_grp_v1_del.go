package article_grp

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article_grp/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.ArticleGrp().Del(ctx, req.Id)
	return
}
