package admin

import (
	"context"
	v1 "oldme-api/api/v1"
	"oldme-api/internal/service"
)

var Article = cArticle{}

type cArticle struct {
}

func (c *cArticle) Cre(ctx context.Context, req *v1.ArticleCreReq) (res *v1.ArticleCreRes, err error) {
	err = service.Article().Cre(ctx, req.ArticleInput)
	return
}
