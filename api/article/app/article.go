package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/oldme-git/oldme-api/internal/model"
)

type (
	ListReq struct {
		g.Meta `path:"article/list/*grpId" method:"get" sm:"查询列表" tags:"app"`
		model.Paging
		GrpId  model.Id `v:"integer|between:1,4294967295" json:"grpId"`
		Search string   `v:"length: 1,30" json:"search" dc:"查询文本，会检索标题、标签、简介"`
	}

	ListRes struct {
		List  []List `json:"list"`
		Total uint   `json:"total"`
	}
)

type (
	RankReq struct {
		g.Meta `path:"article/rank" method:"get" sm:"查询文章排行" tags:"app"`
		Basis  int `v:"required|in:1,2" dc:"1-热门文章 2-最新文章"`
	}

	RankRes struct {
		List []List `json:"list"`
	}
)

type (
	ShowReq struct {
		g.Meta `path:"article/show/{id}" method:"get" sm:"查询详情" tags:"app"`
		*model.IdInput
	}

	ShowRes struct {
		*One
	}
)

type (
	AboutReq struct {
		g.Meta `path:"about" method:"get" sm:"查询关于我们" tags:"app"`
	}

	AboutRes struct {
		*One
	}
)

type (
	HistReq struct {
		g.Meta `path:"article/hist" method:"post" sm:"增加一个点击数" tags:"app"`
		*model.IdInput
	}

	HistRes struct {
	}
)

type (
	ReplyReq struct {
		g.Meta `path:"article/reply/{id}" method:"get" sm:"查看该文章的回复" tags:"app"`
		*model.IdInput
	}

	ReplyRes struct {
		List  []model.ReplyFloorApp `json:"list"`
		Total uint                  `json:"total"`
	}
)

type (
	List struct {
		Id          uint        `json:"id"          description:""`
		GrpId       uint        `json:"grpId"       description:"分组id"`
		Title       string      `json:"title"       description:"标题"`
		Author      string      `json:"author"      description:"作者"`
		Thumb       string      `json:"thumb"       description:"图片地址"`
		Tags        string      `json:"tags"        description:"标签，依英文逗号隔开"`
		Description string      `json:"description" description:"简介"`
		Hist        uint        `json:"hist"        description:"点击数"`
		Post        uint        `json:"post"        description:"评论数"`
		CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
		UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
		LastedAt    *gtime.Time `json:"lastedAt"    description:"最后浏览时间"`
	}

	One struct {
		Id          uint        `json:"id"          description:""`
		GrpId       uint        `json:"grpId"       description:"分组id"`
		Title       string      `json:"title"       description:"标题"`
		Author      string      `json:"author"      description:"作者"`
		Thumb       string      `json:"thumb"       description:"图片地址"`
		Tags        string      `json:"tags"        description:"标签，依英文逗号隔开"`
		Description string      `json:"description" description:"简介"`
		Content     string      `json:"content"     description:"内容"`
		Hist        uint        `json:"hist"        description:"点击数"`
		Post        uint        `json:"post"        description:"评论数"`
		CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
		UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
		LastedAt    *gtime.Time `json:"lastedAt"    description:"最后浏览时间"`
	}
)
