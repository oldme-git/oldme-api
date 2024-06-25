package article_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article_grp/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article_grp"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, err := article_grp.List(ctx)
	if err == nil {
		res = &v1.ListRes{
			List:  list,
			Total: uint(len(list)),
		}
	}
	return
}
