package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"oldme-api/internal/model"
)

type ArticleCreReq struct {
	g.Meta `path:"article/create" method:"post" sm:"新增" tags:"文章"`
	*model.ArticleInput
}

type ArticleCreRes struct {
	LastId uint `json:"lastId"`
}

type ArticleUptReq struct {
	g.Meta `path:"article/update/{id}" method:"post" sm:"修改" tags:"文章"`
	Id     uint `v:"integer|between:1,999999999"`
	*model.ArticleInput
}

type ArticleUptRes struct {
}

type ArticleDelReq struct {
	g.Meta `path:"article/delete" method:"post" sm:"删除" tags:"文章"`
}

type ArticleDelRes struct {
}

type ArticleListReq struct {
	g.Meta `path:"article/list/*grpId" method:"get" sm:"查询列表" tags:"文章"`
	GrpId  uint `v:"integer|between:1,999999999"`
}

type ArticleListRes struct {
}

type ArticleShowReq struct {
	g.Meta `path:"article/show/{id}" method:"get" sm:"查询详情" tags:"文章"`
	Id     uint `v:"integer|between:1,999999999"`
}

type ArticleShowRes struct {
}
