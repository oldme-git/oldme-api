package model

type Id uint32
type ReplyStatus uint8

// IdInput 公共Id input，一般用作api中的upd和del
type IdInput struct {
	Id Id `json:"aid" v:"required|integer|between:1,4294967295"`
}

// Paging 分页查询条件
type Paging struct {
	Page int `v:"integer|between:1,9999999" json:"page" dc:"查询分页：页码，默认1"`
	Size int `v:"integer|between:1,100" json:"size" dc:"查询分页：条数，默认15"`
}
