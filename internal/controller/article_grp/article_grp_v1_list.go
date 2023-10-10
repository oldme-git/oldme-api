package article_grp

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article_grp/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, err := service.ArticleGrp().List(ctx)
	if err == nil {
		res = &v1.ListRes{
			List:  list,
			Total: uint(len(list)),
		}
	}
	return
}
