package app

import (
	"context"
	v1 "oldme-api/api/app/v1"
	"oldme-api/internal/model"
	"oldme-api/internal/service"
)

var ArticleGrp = cArticleGrp{}

type cArticleGrp struct {
}

func (c *cArticleGrp) List(ctx context.Context, req *v1.ArticleGrpListReq) (res *v1.ArticleGrpListRes, err error) {
	list, err := service.ArticleGrp().List(ctx)
	var listSafe []model.ArticleGrpListSafe
	if err == nil {
		for _, v := range list {
			listSafe = append(listSafe, model.ArticleGrpListSafe{ArticleGrp: v})
		}
		res = &v1.ArticleGrpListRes{
			List: listSafe,
		}
	}
	return
}
