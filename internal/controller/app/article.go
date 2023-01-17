package app

import (
	"context"
	"oldme-api/api/app/v1"
)

var Article = cArticle{}

type cArticle struct {
}

func (c *cArticle) List(ctx context.Context, req *v1.ArticleListReq) (res *v1.ArticleListRes, err error) {
	//query := &model.ArticleQuery{
	//	GrpId:  req.ArticleQueryApp.GrpId,
	//	Page:   req.ArticleQueryApp.Page,
	//	Size:   req.ArticleQueryApp.Size,
	//	Search: req.ArticleQueryApp.Search,
	//	IsDel:  false,
	//}
	//list, total, err := service.Article().List(ctx, query)
	//var listApp model.ArticleListApp
	//for _, v := range *list {
	//	listApp = append(listApp)
	//	fmt.Println(v)
	//}
	//if err == nil {
	//	// 查询数据表里总共的数据条数
	//	res = &v1.ArticleListRes{
	//		List:  &listApp,
	//		Total: total,
	//	}
	//}
	return
}
