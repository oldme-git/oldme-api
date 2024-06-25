package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/v1"
	"github.com/oldme-git/oldme-api/internal/logic/article"
	"github.com/oldme-git/oldme-api/internal/logic/article_grp"
	"github.com/oldme-git/oldme-api/internal/utility"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	// 判断该分类是否存在
	if ok := article_grp.IsExist(ctx, req.ArticleInput.GrpId); !ok {
		err = utility.Err.Skip(10101)
		return
	}
	err = article.Upd(ctx, req.Id, req.ArticleInput)
	return
}
