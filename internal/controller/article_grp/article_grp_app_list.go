package article_grp

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article_grp/app"
)

func (c *ControllerApp) List(ctx context.Context, req *app.ListReq) (res *app.ListRes, err error) {
	list, err := service.ArticleGrp().List(ctx)
	var listSafe []model.ArticleGrpListSafe
	if err == nil {
		countList, _ := service.ArticleGrp().ListArticleCount(ctx)
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
