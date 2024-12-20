package app

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"article/group/list" method:"get" sm:"查询列表" tags:"app"`
}

type ListRes struct {
	List []List `json:"list"`
}

type List struct {
	Id           uint   `json:"id"          description:""`
	Name         string `json:"name"        description:"名称"`
	Tags         string `json:"tags"        description:"标签，依英文逗号隔开"`
	Description  string `json:"description" description:"简介"`
	ArticleCount uint   `json:"article_count"`
}
