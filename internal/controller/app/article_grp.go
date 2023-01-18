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
	var listApp []model.ArticleGrpListApp
	for _, v := range *list {
		listApp = append(listApp, model.ArticleGrpListApp{
			Id:          v.Id,
			Name:        v.Name,
			Tags:        v.Tags,
			Description: v.Description,
		})
	}
	if err == nil {
		res = &v1.ArticleGrpListRes{
			List: &listApp,
		}
	}
	return
}
