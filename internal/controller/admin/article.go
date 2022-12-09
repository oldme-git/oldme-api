package admin

import (
	"context"
	"fmt"
	v1 "oldme-api/api/v1"
	"oldme-api/internal/service"
)

var Article = cArticle{}

type cArticle struct {
}

func (c *cArticle) Cre(ctx context.Context, req *v1.ArticleCreReq) (res *v1.ArticleCreRes, err error) {
	LastId, err := service.Article().Cre(ctx, req.ArticleInput)
	if err == nil {
		res = &v1.ArticleCreRes{LastId: LastId}
	}
	return
}

func (c *cArticle) Upt(ctx context.Context, req *v1.ArticleUptReq) (res *v1.ArticleUptRes, err error) {
	err = service.Article().Upt(ctx, req.Id, req.ArticleInput)
	return
}

func (c *cArticle) List(ctx context.Context, req *v1.ArticleListReq) (res *v1.ArticleListRes, err error) {
	fmt.Println(req.GrpId)
	return
}
