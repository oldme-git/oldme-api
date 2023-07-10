package model

type ReplyInput struct {
	Aid     Id          `json:"aid" v:"required|integer|between:1,4294967295"`
	Pid     Id          `json:"pid" v:"required|integer|between:0,4294967295"`
	Name    string      `json:"name" v:"required|length:1, 20"`
	Email   string      `json:"email" v:"required|email"`
	Content string      `json:"content" v:"required|length:2, 100000"`
	Status  ReplyStatus `json:"status" v:"required|in:0,1,2"`
	Site    string      `json:"site" v:"url"`
	Reason  string      `json:"reason" v:"length:1,50"`
}

type ReplyQuery struct {
	Paging
	Search string `v:"length: 1,30" json:"search" dc:"查询文本，会检索名称，邮箱，内容"`
}
