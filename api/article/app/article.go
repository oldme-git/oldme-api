package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
)

type ListReq struct {
	g.Meta `path:"article/list/*grpId" method:"get" sm:"查询列表" tags:"app"`
	*model.ArticleQuerySafe
}

type ListRes struct {
	List  []model.ArticleListSafe `json:"list"`
	Total uint                    `json:"total"`
}

type RankReq struct {
	g.Meta `path:"article/rank" method:"get" sm:"查询文章排行" tags:"app"`
	Basis  int `v:"required|in:1,2" dc:"1-热门文章 2-最新文章"`
}

type RankRes struct {
	List []model.ArticleListSafe `json:"list"`
}

type ShowReq struct {
	g.Meta `path:"article/show/{id}" method:"get" sm:"查询详情" tags:"app"`
	*model.IdInput
}

type ShowRes struct {
	*model.ArticleSafe
}

type AboutReq struct {
	g.Meta `path:"about" method:"get" sm:"查询关于我们" tags:"app"`
}

type AboutRes struct {
	*model.ArticleSafe
}

type HistReq struct {
	g.Meta `path:"article/hist" method:"post" sm:"增加一个点击数" tags:"app"`
	*model.IdInput
}

type HistRes struct {
}

type ReplyReq struct {
	g.Meta `path:"article/reply/{id}" method:"get" sm:"查看该文章的回复" tags:"app"`
	model.IdInput
}

type ReplyRes struct {
	List  []model.ReplyFloorApp `json:"list"`
	Total uint                  `json:"total"`
}
