package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type ArticleCreReq struct {
	g.Meta `path:"article/create" method:"post" sm:"新增" tags:"文章"`
	*model.ArticleInput
}

type ArticleCreRes struct {
	LastId model.Id `json:"lastId"`
}

type ArticleUpdReq struct {
	g.Meta `path:"article/update/{id}" method:"post" sm:"修改" tags:"文章"`
	*model.IdInput
	*model.ArticleInput
}

type ArticleUpdRes struct {
}

type ArticleDelReq struct {
	g.Meta `path:"article/delete/{id}" method:"post" sm:"删除" tags:"文章"`
	*model.IdInput
	IsReal bool `dc:"是否彻底删除"`
}

type ArticleDelRes struct {
}

type ArticleListReq struct {
	g.Meta `path:"article/list/*grpId" method:"get" sm:"查询列表" tags:"文章"`
	*model.ArticleQuery
}

type ArticleListRes struct {
	List  []model.ArticleList `json:"list"`
	Total uint                `json:"total"`
}

type ArticleShowReq struct {
	g.Meta `path:"article/show/{id}" method:"get" sm:"查询详情" tags:"文章"`
	*model.IdInput
}

type ArticleShowRes struct {
	*entity.Article
}

type ArticleReCreReq struct {
	g.Meta `path:"article/recreate/{id}" method:"post" sm:"找回文章" tags:"文章"`
	*model.IdInput
}

type ArticleReCreRes struct {
}
