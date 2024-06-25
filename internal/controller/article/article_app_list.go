package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/app"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerApp) List(ctx context.Context, req *app.ListReq) (res *app.ListRes, err error) {
	query := &model.ArticleQuery{
		ArticleQuerySafe: *req.ArticleQuerySafe,
		Onshow:           true,
		IsDel:            false,
	}
	list, total, err := article.List(ctx, query)
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
