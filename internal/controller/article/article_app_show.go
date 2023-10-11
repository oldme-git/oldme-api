package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article/app"
)

func (c *ControllerApp) Show(ctx context.Context, req *app.ShowReq) (res *app.ShowRes, err error) {
	info, err := service.Article().Show(ctx, req.Id)
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
			_ = service.Article().UpdLastedAt(ctx, model.Id(info.Id))
		}
	}
	return
}
