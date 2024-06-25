package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, total, err := article.List(ctx, req.ArticleQuery)
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
