package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"oldme-api/internal/model"
)

type ArticleListReq struct {
	g.Meta `path:"article/list/*grpId" method:"get" sm:"查询列表" tags:"文章"`
	*model.ArticleQueryApp
}

type ArticleListRes struct {
	List  *model.ArticleListApp `json:"list"`
	Total uint                  `json:"total"`
}
