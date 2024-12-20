package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/app"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerApp) Show(ctx context.Context, req *app.ShowReq) (res *app.ShowRes, err error) {
	info, err := article.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if info == nil {
		return nil, err
	}

	if info.Onshow != 0 {
		_ = article.UpdLastedAt(ctx, model.Id(info.Id))
	}

	return &app.ShowRes{
		One: &app.One{
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
	}, nil
}
