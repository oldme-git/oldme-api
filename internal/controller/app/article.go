package app

import (
	"context"
	"github.com/oldme-git/oldme-api/api/app/v1"
	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/service"
)

var Article = cArticle{}

type cArticle struct {
}

func (c *cArticle) List(ctx context.Context, req *v1.ArticleListReq) (res *v1.ArticleListRes, err error) {
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
		res = &v1.ArticleListRes{
			List:  listOut,
			Total: total,
		}
	}
	return
}

func (c *cArticle) ArticleRank(ctx context.Context, req *v1.ArticleRankReq) (res *v1.ArticleRankRes, err error) {
	db := dao.Article.Ctx(ctx)
	if req.Basis == 1 {
		db = db.Order("ontop desc,order desc,post desc,hist desc")
	} else {
		db = db.Order("created_at desc")
	}
	db = db.Where("onshow", true)
	data, err := db.Limit(0, 10).All()
	if err != nil {
		return
	}
	var (
		list    []entity.Article
		listOut []model.ArticleListSafe
	)
	_ = data.Structs(&list)
	for _, v := range list {
		listOut = append(listOut, model.ArticleListSafe{Article: v})
	}
	if err == nil {
		res = &v1.ArticleRankRes{
			List: listOut,
		}
	}
	return
}

func (c *cArticle) Show(ctx context.Context, req *v1.ArticleShowReq) (res *v1.ArticleShowRes, err error) {
	info, err := service.Article().Show(ctx, req.Id)
	if err == nil {
		if info == nil || info.Onshow == 0 {
			res = &v1.ArticleShowRes{
				ArticleSafe: nil,
			}
		} else {
			res = &v1.ArticleShowRes{
				ArticleSafe: &model.ArticleSafe{
					Article: info,
				},
			}
			_ = service.Article().UpdLastedAt(ctx, model.Id(info.Id))
		}
	}
	return
}

func (c *cArticle) About(ctx context.Context, req *v1.AboutShowReq) (res *v1.ArticleShowRes, err error) {
	info, err := service.Article().Show(ctx, 1)

	if err == nil {
		res = &v1.ArticleShowRes{
			ArticleSafe: &model.ArticleSafe{
				Article: info,
			},
		}
	}
	return
}

func (c *cArticle) Hist(ctx context.Context, req *v1.ArticleHistReq) (res *v1.ArticleHistRes, err error) {
	err = service.Article().Hist(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cArticle) ReplyList(ctx context.Context, req *v1.ArticleReplyReq) (res *v1.ArticleReplyRes, err error) {
	total, list, err := service.Reply().ListForAid(ctx, req.Id)
	if err == nil {
		res = &v1.ArticleReplyRes{
			List:  list,
			Total: total,
		}
	}
	return
}
