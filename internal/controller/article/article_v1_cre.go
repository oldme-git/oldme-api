package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/logic/article_grp"
	"github.com/oldme-git/oldme-api/internal/utility"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	// 判断该分类是否存在
	if ok := article_grp.IsExist(ctx, req.ArticleInput.GrpId); !ok {
		err = utility.Err.Skip(10101)
		return
	}
	LastId, err := article.Cre(ctx, req.ArticleInput)
	if err == nil {
		res = &v1.CreRes{LastId: LastId}
	}
	return
}
