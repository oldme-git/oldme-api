package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/logic/article_grp"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/utility"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (*v1.CreRes, error) {
	// 判断该分类是否存在
	if ok := article_grp.IsExist(ctx, req.GrpId); !ok {
		return nil, utility.Err.Skip(10101)
	}

	LastId, err := article.Cre(ctx, &model.ArticleInput{
		GrpId:       req.GrpId,
		Title:       req.Title,
		Author:      req.Author,
		Thumb:       req.Thumb,
		Tags:        req.Tags,
		Description: req.Description,
		Content:     req.Content,
		Order:       req.Order,
		Ontop:       req.Ontop,
		Onshow:      req.Onshow,
		Hist:        req.Hist,
		Post:        req.Post,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreRes{LastId: LastId}, nil
}
