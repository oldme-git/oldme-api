package article_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article_grp/app"
	"github.com/oldme-git/oldme-api/internal/logic/article_grp"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerApp) List(ctx context.Context, req *app.ListReq) (res *app.ListRes, err error) {
	list, err := article_grp.List(ctx)
	var listSafe []model.ArticleGrpListSafe
	if err == nil {
		countList, _ := article_grp.ListArticleCount(ctx)
		for _, v := range list {
			listSafe = append(listSafe, model.ArticleGrpListSafe{
				ArticleGrp:   v,
				ArticleCount: countList[v.Id],
			})
		}
		res = &app.ListRes{
			List: listSafe,
		}
	}
	return
}
