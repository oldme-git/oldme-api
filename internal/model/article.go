package model

import "github.com/gogf/gf/v2/os/gtime"

type ArticleInput struct {
	GrpId       uint   `json:"grpId" v:"required|integer|between:1,4294967295"`
	Title       string `json:"title" v:"required|length:2, 100"`
	Author      string `json:"author" v:"length:2, 30"`
	Thumb       string `json:"thumb" v:"length:2, 200"`
	Tags        string `json:"tags" v:"length:2, 200"`
	Description string `json:"description" v:"length:2, 200"`
	Content     string `json:"content" v:"length:2, 10000"`
	Order       int    `json:"order" v:"integer|between:0,9999"`
	Ontop       uint   `json:"ontop" v:"boolean"`
	Onshow      uint   `json:"onshow" v:"boolean"`
	Hist        uint   `json:"hist" v:"integer|between:0,999999"`
	Post        uint   `json:"post" v:"integer|between:0,999999"`
}

type ArticleList []struct {
	Id          uint        `json:"id"          description:""`
	GrpId       uint        `json:"grpId"       description:"分组id"`
	Title       string      `json:"title"       description:"标题"`
	Author      string      `json:"author"      description:"作者"`
	Thumb       string      `json:"thumb"       description:"图片地址"`
	Tags        string      `json:"tags"        description:"标签，依英文逗号隔开"`
	Description string      `json:"description" description:"简介"`
	Order       int         `json:"order"       description:"排序，越大越靠前"`
	Ontop       uint        `json:"ontop"       description:"是否置顶"`
	Onshow      uint        `json:"onshow"      description:"是否显示"`
	Hist        uint        `json:"hist"        description:"点击数"`
	Post        uint        `json:"post"        description:"评论数"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
}

type ArticleQuery struct {
	GrpId  uint   `v:"integer|between:1,999999999" json:"grpId"`
	Page   int    `v:"integer|between:1,999999999" json:"page" dc:"查询分页：页码，默认1"`
	Size   int    `v:"integer|between:1,999999999" json:"size" dc:"查询分页：条数，默认15"`
	Search string `v:"length: 1,30" json:"search" dc:"查询文本，会检索标题、标签、简介"`
	IsDel  bool   ` json:"isDel" dc:"是否查询删除掉的文章"`
}
