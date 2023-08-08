package model

import (
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type ArticleInput struct {
	GrpId       Id     `json:"grpId" v:"required|integer|between:1,4294967295"`
	Title       string `json:"title" v:"required|length:2, 100"`
	Author      string `json:"author" v:"length:2, 30"`
	Thumb       string `json:"thumb" v:"length:2, 200"`
	Tags        string `json:"tags" v:"length:2, 200"`
	Description string `json:"description" v:"length:2, 200"`
	Content     string `json:"content" v:"length:2, 100000"`
	Order       int    `json:"order" v:"integer|between:-9999,9999"`
	Ontop       uint   `json:"ontop" v:"boolean"`
	Onshow      uint   `json:"onshow" v:"boolean"`
	Hist        uint   `json:"hist" v:"integer|between:0,999999"`
	Post        uint   `json:"post" v:"integer|between:0,999999"`
}

type ArticleQuery struct {
	ArticleQuerySafe
	Onshow bool `json:"onshow" dc:"是否查询只发布的文章"`
	IsDel  bool `json:"isDel" dc:"是否查询删除掉的文章"`
}

type ArticleQuerySafe struct {
	Paging
	GrpId  Id     `v:"integer|between:1,4294967295" json:"grpId"`
	Search string `v:"length: 1,30" json:"search" dc:"查询文本，会检索标题、标签、简介"`
}

type ArticleList struct {
	entity.Article
	Content   Out `json:"content,omitempty"`
	DeletedAt Out `json:"deletedAt,omitempty"`
}

type ArticleListSafe struct {
	entity.Article
	ArticleSafe
	DeletedAt Out `json:"deletedAt,omitempty"`
}

type ArticleSafe struct {
	*entity.Article
	Order     Out `json:"order,omitempty"`
	Ontop     Out `json:"ontop,omitempty"`
	Onshow    Out `json:"onshow,omitempty"`
	DeletedAt Out `json:"deletedAt,omitempty"`
}
