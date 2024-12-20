package model

type ArticleInput struct {
	GrpId       Id
	Title       string
	Author      string
	Thumb       string
	Tags        string
	Description string
	Content     string
	Order       int
	Ontop       uint
	Onshow      uint
	Hist        uint
	Post        uint
}

type ArticleQuery struct {
	Paging
	GrpId  Id
	Search string
	Onshow bool
	IsDel  bool
}
