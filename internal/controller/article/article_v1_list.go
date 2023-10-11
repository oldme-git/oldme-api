package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, total, err := service.Article().List(ctx, req.ArticleQuery)
	if err == nil {
		var listOut []model.ArticleList
		for _, v := range list {
			listOut = append(listOut, model.ArticleList{Article: v})
		}
		// 查询数据表里总共的数据条数
		res = &v1.ListRes{
			List:  listOut,
			Total: total,
		}
	}
	return
}
