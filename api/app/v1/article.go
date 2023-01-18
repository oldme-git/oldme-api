package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"oldme-api/internal/model"
)

type ArticleListReq struct {
	g.Meta `path:"article/list/*grpId" method:"get" sm:"查询列表" tags:"文章app"`
	*model.ArticleQueryApp
}

type ArticleListRes struct {
	List  *[]model.ArticleListApp `json:"list"`
	Total uint                    `json:"total"`
}

type ArticleShowReq struct {
	g.Meta `path:"article/show/{id}" method:"get" sm:"查询详情" tags:"文章app"`
	Id     uint `v:"integer|between:1,999999999"`
}

type ArticleShowRes struct {
	*model.ArticleShowApp
}
