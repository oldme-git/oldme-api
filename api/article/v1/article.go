package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/oldme-git/oldme-api/internal/model"
)

type (
	CreReq struct {
		g.Meta      `path:"article/create" method:"post" sm:"新增" tags:"文章"`
		GrpId       model.Id `json:"grpId" v:"required|integer|between:1,4294967295"`
		Title       string   `json:"title" v:"required|length:2, 100"`
		Author      string   `json:"author" v:"length:2, 30"`
		Thumb       string   `json:"thumb" v:"length:2, 200"`
		Tags        string   `json:"tags" v:"length:2, 200"`
		Description string   `json:"description" v:"length:2, 200"`
		Content     string   `json:"content" v:"length:2, 100000"`
		Order       int      `json:"order" v:"integer|between:-9999,9999"`
		Ontop       uint     `json:"ontop" v:"boolean"`
		Onshow      uint     `json:"onshow" v:"boolean"`
		Hist        uint     `json:"hist" v:"integer|between:0,999999"`
		Post        uint     `json:"post" v:"integer|between:0,999999"`
	}

	CreRes struct {
		LastId model.Id `json:"lastId"`
	}
)

type (
	UpdReq struct {
		g.Meta `path:"article/update/{id}" method:"post" sm:"修改" tags:"文章"`
		*model.IdInput
		GrpId       model.Id `json:"grpId" v:"required|integer|between:1,4294967295"`
		Title       string   `json:"title" v:"required|length:2, 100"`
		Author      string   `json:"author" v:"length:2, 30"`
		Thumb       string   `json:"thumb" v:"length:2, 200"`
		Tags        string   `json:"tags" v:"length:2, 200"`
		Description string   `json:"description" v:"length:2, 200"`
		Content     string   `json:"content" v:"length:2, 100000"`
		Order       int      `json:"order" v:"integer|between:-9999,9999"`
		Ontop       uint     `json:"ontop" v:"boolean"`
		Onshow      uint     `json:"onshow" v:"boolean"`
		Hist        uint     `json:"hist" v:"integer|between:0,999999"`
		Post        uint     `json:"post" v:"integer|between:0,999999"`
	}

	UpdRes struct{}
)

type (
	DelReq struct {
		g.Meta `path:"article/delete/{id}" method:"post" sm:"删除" tags:"文章"`
		*model.IdInput
		IsReal bool `dc:"是否彻底删除"`
	}

	DelRes struct{}
)

type (
	ListReq struct {
		g.Meta `path:"article/list/*grpId" method:"get" sm:"查询列表" tags:"文章"`
		model.Paging
		GrpId  model.Id `v:"integer|between:1,4294967295" json:"grpId"`
		Search string   `v:"length: 1,30" json:"search" dc:"查询文本，会检索标题、标签、简介"`
		Onshow bool     `json:"onshow" dc:"是否查询只发布的文章"`
		IsDel  bool     `json:"isDel" dc:"是否查询删除掉的文章"`
	}

	ListRes struct {
		List  []List `json:"list"`
		Total uint   `json:"total"`
	}
)

type (
	ShowReq struct {
		g.Meta `path:"article/show/{id}" method:"get" sm:"查询详情" tags:"文章"`
		*model.IdInput
	}

	ShowRes struct {
		*One
	}
)

type (
	ReCreReq struct {
		g.Meta `path:"article/recreate/{id}" method:"post" sm:"找回文章" tags:"文章"`
		*model.IdInput
	}

	ReCreRes struct {
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
		Order       int         `json:"order"       description:"排序，越大越靠前"`
		Ontop       uint        `json:"ontop"       description:"是否置顶"`
		Onshow      uint        `json:"onshow"      description:"是否显示"`
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
		Order       int         `json:"order"       description:"排序，越大越靠前"`
		Ontop       uint        `json:"ontop"       description:"是否置顶"`
		Onshow      uint        `json:"onshow"      description:"是否显示"`
		Hist        uint        `json:"hist"        description:"点击数"`
		Post        uint        `json:"post"        description:"评论数"`
		CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
		UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
		DeletedAt   *gtime.Time `json:"deletedAt"   description:"删除时间"`
		LastedAt    *gtime.Time `json:"lastedAt"    description:"最后浏览时间"`
	}
)
