package article_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article_grp/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article_grp"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = article_grp.Upd(ctx, req.Id, &model.ArticleGrpInput{
		Name:        req.Name,
		Tags:        req.Tags,
		Description: req.Description,
		Onshow:      req.Onshow,
		Order:       req.Order,
	})
	return
}
