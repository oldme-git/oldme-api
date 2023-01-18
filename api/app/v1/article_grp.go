package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"oldme-api/internal/model"
)

type ArticleGrpListReq struct {
	g.Meta `path:"article/group/list" method:"get" sm:"查询列表" tags:"文章分类"`
	*model.ArticleQueryApp
}

type ArticleGrpListRes struct {
	List *[]model.ArticleGrpListApp `json:"list"`
}
