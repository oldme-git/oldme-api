package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, total, err := article.List(ctx, &model.ArticleQuery{
		Paging: model.Paging{
			Page: req.Page,
			Size: req.Size,
		},
		GrpId:  req.GrpId,
		Search: req.Search,
		Onshow: req.Onshow,
		IsDel:  req.IsDel,
	})

	if err != nil {
		return nil, err
	}
	var listOut []v1.List
	for _, v := range list {
		listOut = append(listOut, v1.List{
			Id:          v.Id,
			GrpId:       v.GrpId,
			Title:       v.Title,
			Author:      v.Author,
			Thumb:       v.Thumb,
			Tags:        v.Tags,
			Description: v.Description,
			Order:       v.Order,
			Ontop:       v.Ontop,
			Onshow:      v.Onshow,
			Hist:        v.Hist,
			Post:        v.Post,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			LastedAt:    v.LastedAt,
		})
	}
	// 查询数据表里总共的数据条数
	return &v1.ListRes{
		List:  listOut,
		Total: total,
	}, nil
}
