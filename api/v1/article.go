package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"oldme-api/internal/model"
)

type ArticleCreReq struct {
	g.Meta `path:"article/create" method:"post" sm:"新增" tags:"文章"`
	model.ArticleInput
}

type ArticleCreRes struct {
}
