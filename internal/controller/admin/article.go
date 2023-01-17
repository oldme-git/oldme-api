package admin

import (
	"context"
	"oldme-api/api/admin/v1"
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
	list, total, err := service.Article().List(ctx, req.ArticleQuery)
	if err == nil {
		// 查询数据表里总共的数据条数
		res = &v1.ArticleListRes{
			List:  list,
			Total: total,
		}
	}
	return
}

func (c *cArticle) Show(ctx context.Context, req *v1.ArticleShowReq) (res *v1.ArticleShowRes, err error) {
	info, err := service.Article().Show(ctx, req.Id)
	if err == nil {
		res = &v1.ArticleShowRes{
			Article: info,
		}
	}
	return
}

func (c *cArticle) Del(ctx context.Context, req *v1.ArticleDelReq) (res *v1.ArticleDelRes, err error) {
	err = service.Article().Del(ctx, req.Id, req.IsReal)
	return
}

func (c *cArticle) ReCre(ctx context.Context, req *v1.ArticleReCreReq) (res *v1.ArticleReCreRes, err error) {
	err = service.Article().ReCre(ctx, req.Id)
	return
}
