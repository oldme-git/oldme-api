package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type CreReq struct {
	g.Meta   `path:"sentence/create" method:"post" sm:"新增" tags:"句子"`
	BookId   model.Id   `json:"book_id" v:"required"`
	TagIds   []model.Id `json:"tag_ids"`
	Sentence string     `json:"sentence" v:"required"`
}

type CreRes struct {
}

type UpdReq struct {
	g.Meta `path:"sentence/update/{id}" method:"post" sm:"修改" tags:"句子"`
	*model.IdInput
	BookId   model.Id   `json:"book_id" v:"required"`
	TagIds   []model.Id `json:"tag_ids"`
	Sentence string     `json:"sentence" v:"required"`
}

type UpdRes struct {
}

type ShowReq struct {
	g.Meta `path:"sentence/show/{id}" method:"get" sm:"查询详情" tags:"句子"`
	*model.IdInput
}

type ShowRes struct {
	*entity.Sentence
	TagList []entity.Tag `json:"tag_list"`
}

type DelReq struct {
	g.Meta `path:"sentence/delete/{id}" method:"post" sm:"删除" tags:"句子"`
	*model.IdInput
}

type DelRes struct {
}

type ListReq struct {
	g.Meta `path:"sentence/list" method:"get" sm:"查询列表" tags:"句子"`
	*model.Paging
	BookId model.Id   `json:"book_id" dc:"书籍id"`
	TagIds []model.Id `json:"tag_ids" dc:"标签id"`
	Search string     `json:"search"  dc:"搜索文本"`
}

type ListRes struct {
	List  []List `json:"list"`
	Total uint   `json:"total"`
}
