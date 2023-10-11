package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article/app"
)

func (c *ControllerApp) About(ctx context.Context, req *app.AboutReq) (res *app.AboutRes, err error) {
	info, err := service.Article().Show(ctx, 1)

	if err == nil {
		res = &app.AboutRes{
			ArticleSafe: &model.ArticleSafe{
				Article: info,
			},
		}
	}
	return
}
