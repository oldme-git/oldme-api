package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	info, err := article.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.ShowRes{
		One: &v1.One{
			Id:          info.Id,
			GrpId:       info.GrpId,
			Title:       info.Title,
			Author:      info.Author,
			Thumb:       info.Thumb,
			Tags:        info.Tags,
			Description: info.Description,
			Content:     info.Content,
			Order:       info.Order,
			Ontop:       info.Ontop,
			Onshow:      info.Onshow,
			Hist:        info.Hist,
			Post:        info.Post,
			CreatedAt:   info.CreatedAt,
			UpdatedAt:   info.UpdatedAt,
			DeletedAt:   info.DeletedAt,
			LastedAt:    info.LastedAt,
		},
	}, nil
}
