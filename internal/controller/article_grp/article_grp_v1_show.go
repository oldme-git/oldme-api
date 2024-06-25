package article_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article_grp/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article_grp"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	info, err := article_grp.Show(ctx, req.Id)
	if err == nil {
		res = &v1.ShowRes{
			ArticleGrp: info,
		}
	}
	return
}
