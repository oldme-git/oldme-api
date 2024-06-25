package article_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article_grp/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article_grp"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = article_grp.Del(ctx, req.Id)
	return
}
