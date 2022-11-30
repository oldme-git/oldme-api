package model

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
