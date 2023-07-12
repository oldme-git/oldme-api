package model

import "oldme-api/internal/model/entity"

const (
	AllStatus = iota
	CheckStatus
	SuccessStatus
	FailStatus
)

// ReplyInput 新增回复
type ReplyInput struct {
	ReplyBody
	Aid    Id          `json:"aid" v:"required|integer|between:1,4294967295"`
	Pid    Id          `json:"pid" v:"required|integer|between:0,4294967295"`
	Status ReplyStatus `json:"status" v:"in:1,2,3" dc:"审核状态，1待审核 2审核通过 3审核失败，默认2"`
	Reason string      `json:"reason" v:"length:1,50"`
}

// ReplyInputApp App新增回复
type ReplyInputApp struct {
	ReplyBody
	Aid Id `json:"aid" v:"required|integer|between:1,4294967295"`
	Pid Id `json:"pid" v:"required|integer|between:0,4294967295"`
}

// ReplyBody 回复内容主体
type ReplyBody struct {
	Name    string `json:"name" v:"required|length:1,20"`
	Email   string `json:"email" v:"required|email|length:1,100"`
	Site    string `json:"site" v:"url|length:1,50"`
	Content string `json:"content" v:"required|length:2, 100000"`
}

type ReplyQuery struct {
	Paging
	Status ReplyStatus `json:"status" v:"in:1,2,3" dc:"查询状态，1待审核 2审核通过 3审核失败	，不传则查询所有"`
	Search string      `v:"length: 1,30" json:"search" dc:"查询文本，会检索名称，邮箱，内容"`
}

type ReplyShow struct {
	ArticleTitle string       `json:"articleTitle" dc:"回复的文章标题"`
	ParentReply  entity.Reply `json:"parentReply" dc:"父级回复"`
	entity.Reply
}
