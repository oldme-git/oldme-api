package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"

	"github.com/oldme-git/oldme-api/api/article/app"
)

func (c *ControllerApp) Rank(ctx context.Context, req *app.RankReq) (res *app.RankRes, err error) {
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
		res = &app.RankRes{
			List: listOut,
		}
	}
	return
}
