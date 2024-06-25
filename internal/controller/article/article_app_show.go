package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/app"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerApp) Show(ctx context.Context, req *app.ShowReq) (res *app.ShowRes, err error) {
	info, err := article.Show(ctx, req.Id)
	if err == nil {
		if info == nil || info.Onshow == 0 {
			res = &app.ShowRes{
				ArticleSafe: nil,
			}
		} else {
			res = &app.ShowRes{
				ArticleSafe: &model.ArticleSafe{
					Article: info,
				},
			}
			_ = article.UpdLastedAt(ctx, model.Id(info.Id))
		}
	}
	return
}
