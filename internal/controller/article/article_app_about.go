package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/app"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerApp) About(ctx context.Context, req *app.AboutReq) (res *app.AboutRes, err error) {
	info, err := article.Show(ctx, 1)

	if err == nil {
		res = &app.AboutRes{
			ArticleSafe: &model.ArticleSafe{
				Article: info,
			},
		}
	}
	return
}
