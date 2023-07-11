package model

import "oldme-api/internal/model/entity"

const (
	AllStatus = iota
	CheckStatus
	SuccessStatus
	FailStatus
)

type ReplyInput struct {
	Aid     Id          `json:"aid" v:"required|integer|between:1,4294967295"`
	Pid     Id          `json:"pid" v:"required|integer|between:0,4294967295"`
	Name    string      `json:"name" v:"required|length:1, 20"`
	Email   string      `json:"email" v:"required|email|length:1, 100"`
	Content string      `json:"content" v:"required|length:2, 100000"`
	Status  ReplyStatus `json:"status" v:"required|in:1,2,3"`
	Site    string      `json:"site" v:"url"`
	Reason  string      `json:"reason" v:"length:1,50"`
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
