package admin

import (
	"context"
	"oldme-api/api/admin/v1"
	"oldme-api/internal/service"
)

var ArticleGrp = cArticleGrp{}

type cArticleGrp struct {
}

func (c *cArticleGrp) Cre(ctx context.Context, req *v1.ArticleGrpCreReq) (res *v1.ArticleGrpCreRes, err error) {
	err = service.ArticleGrp().Cre(ctx, req.ArticleGrpInput)
	return
}

func (c *cArticleGrp) Upd(ctx context.Context, req *v1.ArticleGrpUpdReq) (res *v1.ArticleGrpUpdRes, err error) {
	err = service.ArticleGrp().Upd(ctx, req.Id, req.ArticleGrpInput)
	return
}

func (c *cArticleGrp) Del(ctx context.Context, req *v1.ArticleGrpDelReq) (res *v1.ArticleGrpDelRes, err error) {
	err = service.ArticleGrp().Del(ctx, req.Id)
	return
}

func (c *cArticleGrp) List(ctx context.Context, req *v1.ArticleGrpListReq) (res *v1.ArticleGroListRes, err error) {
	list, err := service.ArticleGrp().List(ctx)
	if err == nil {
		res = &v1.ArticleGroListRes{
			List:  list,
			Total: uint(len(*list)),
		}
	}
	return
}

func (c *cArticleGrp) Show(ctx context.Context, req *v1.ArticleGrpShowReq) (res *v1.ArticleGrpShowRes, err error) {
	info, err := service.ArticleGrp().Show(ctx, req.Id)
	if err == nil {
		res = &v1.ArticleGrpShowRes{
			ArticleGrp: info,
		}
	}
	return
}
