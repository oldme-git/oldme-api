package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/app"
	"github.com/oldme-git/oldme-api/internal/logic/article"
)

func (c *ControllerApp) About(ctx context.Context, req *app.AboutReq) (res *app.AboutRes, err error) {
	info, err := article.Show(ctx, 1)

	if err == nil {
		res = &app.AboutRes{
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
		}
	}
	return
}
