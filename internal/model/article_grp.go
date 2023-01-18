package model

type ArticleGrpInput struct {
	Name        string `v:"required|length:2, 30"`
	Tags        string `v:"length:1, 200"`
	Description string `v:"length:2, 200"`
	Onshow      bool   `v:"required"`
}

type ArticleGrpListApp struct {
	Id          uint   `json:"id"          description:""`
	Name        string `json:"name"        description:"名称"`
	Tags        string `json:"tags"        description:"标签，依英文逗号隔开"`
	Description string `json:"description" description:"简介"`
}
