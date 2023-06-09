package app

import (
	"context"
	"oldme-api/api/app/v1"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/service"
)

var Article = cArticle{}

type cArticle struct {
}

func (c *cArticle) List(ctx context.Context, req *v1.ArticleListReq) (res *v1.ArticleListRes, err error) {
	query := &model.ArticleQuery{
		GrpId:  req.ArticleQueryApp.GrpId,
		Page:   req.ArticleQueryApp.Page,
		Size:   req.ArticleQueryApp.Size,
		Search: req.ArticleQueryApp.Search,
		Onshow: true,
		IsDel:  false,
	}
	list, total, err := service.Article().List(ctx, query)
	var listApp []model.ArticleListApp
	for _, v := range *list {
		listApp = append(listApp, model.ArticleListApp{
			Id:          v.Id,
			GrpId:       v.GrpId,
			Title:       v.Title,
			Author:      v.Author,
			Thumb:       v.Thumb,
			Tags:        v.Tags,
			Description: v.Description,
			Hist:        v.Hist,
			Post:        v.Post,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			LastedAt:    v.LastedAt,
		})
	}
	if err == nil {
		res = &v1.ArticleListRes{
			List:  &listApp,
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
	list := &[]model.ArticleList{}
	_ = data.Structs(list)
	var listApp []model.ArticleListApp
	for _, v := range *list {
		listApp = append(listApp, model.ArticleListApp{
			Id:          v.Id,
			GrpId:       v.GrpId,
			Title:       v.Title,
			Author:      v.Author,
			Thumb:       v.Thumb,
			Tags:        v.Tags,
			Description: v.Description,
			Hist:        v.Hist,
			Post:        v.Post,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			LastedAt:    v.LastedAt,
		})
	}
	if err == nil {
		res = &v1.ArticleRankRes{
			List: &listApp,
		}
	}
	return
}

func (c *cArticle) Show(ctx context.Context, req *v1.ArticleShowReq) (res *v1.ArticleShowRes, err error) {
	info, err := service.Article().Show(ctx, req.Id)
	if err == nil {
		if info == nil || info.Onshow == 0 {
			res = &v1.ArticleShowRes{
				ArticleShowApp: nil,
			}
		} else {
			res = &v1.ArticleShowRes{
				ArticleShowApp: &model.ArticleShowApp{
					Id:          info.Id,
					GrpId:       info.GrpId,
					Title:       info.Title,
					Author:      info.Author,
					Thumb:       info.Thumb,
					Tags:        info.Tags,
					Description: info.Description,
					Content:     info.Content,
					Hist:        info.Hist,
					Post:        info.Post,
					CreatedAt:   info.CreatedAt,
					UpdatedAt:   info.UpdatedAt,
					LastedAt:    info.LastedAt,
				},
			}
			_ = service.Article().UpdLastedAt(ctx, info.Id)
		}
	}
	return
}

func (c *cArticle) About(ctx context.Context, req *v1.AboutShowReq) (res *v1.ArticleShowRes, err error) {
	info, err := service.Article().Show(ctx, 1)
	if err == nil {
		res = &v1.ArticleShowRes{
			ArticleShowApp: &model.ArticleShowApp{
				Id:          info.Id,
				GrpId:       info.GrpId,
				Title:       info.Title,
				Author:      info.Author,
				Thumb:       info.Thumb,
				Tags:        info.Tags,
				Description: info.Description,
				Content:     info.Content,
				Hist:        info.Hist,
				Post:        info.Post,
				CreatedAt:   info.CreatedAt,
				UpdatedAt:   info.UpdatedAt,
				LastedAt:    info.LastedAt,
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
