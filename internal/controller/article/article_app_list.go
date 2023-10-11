package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article/app"
)

func (c *ControllerApp) List(ctx context.Context, req *app.ListReq) (res *app.ListRes, err error) {
	query := &model.ArticleQuery{
		ArticleQuerySafe: *req.ArticleQuerySafe,
		Onshow:           true,
		IsDel:            false,
	}
	list, total, err := service.Article().List(ctx, query)
	if err == nil {
		var listOut []model.ArticleListSafe
		for _, v := range list {
			listOut = append(listOut, model.ArticleListSafe{Article: v})
		}
		res = &app.ListRes{
			List:  listOut,
			Total: total,
		}
	}
	return
}
