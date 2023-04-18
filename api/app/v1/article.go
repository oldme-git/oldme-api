package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"oldme-api/internal/model"
)

type ArticleListReq struct {
	g.Meta `path:"article/list/*grpId" method:"get" sm:"查询列表" tags:"app"`
	*model.ArticleQueryApp
}

type ArticleListRes struct {
	List  *[]model.ArticleListApp `json:"list"`
	Total uint                    `json:"total"`
}

type ArticleRankReq struct {
	g.Meta `path:"article/rank" method:"get" sm:"查询文章排行" tags:"app"`
	Basis  int `v:"required|in:1,2" dc:"1-热门文章 2-最新文章"`
}

type ArticleRankRes struct {
	List *[]model.ArticleListApp `json:"list"`
}

type ArticleShowReq struct {
	g.Meta `path:"article/show/{id}" method:"get" sm:"查询详情" tags:"app"`
	Id     uint `v:"integer|between:1,999999999"`
}

type ArticleShowRes struct {
	*model.ArticleShowApp
}

type AboutShowReq struct {
	g.Meta `path:"about" method:"get" sm:"查询关于我们" tags:"app"`
}

type ArticleHistReq struct {
	g.Meta `path:"article/hist" method:"post" sm:"增加一个点击数" tags:"app"`
	Id     uint `v:"required|integer|between:1,999999999"`
}

type ArticleHistRes struct {
}
