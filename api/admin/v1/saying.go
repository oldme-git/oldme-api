package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type SayingCreReq struct {
	g.Meta `path:"saying/create" method:"post" sm:"新增" tags:"句子"`
	Saying string `json:"saying" v:"required|length:1, 200"`
}

type SayingCreRes struct {
	LastId model.Id `json:"lastId"`
}

type SayingUpdReq struct {
	g.Meta `path:"saying/update/{id}" method:"post" sm:"修改" tags:"句子"`
	*model.IdInput
	Saying string `json:"saying" v:"required|length:1, 200"`
}

type SayingUpdRes struct {
}

type SayingDelReq struct {
	g.Meta `path:"saying/delete/{id}" method:"post" sm:"删除" tags:"句子"`
	*model.IdInput
}

type SayingDelRes struct {
}

type SayingListReq struct {
	g.Meta `path:"saying/list" method:"get" sm:"查询列表" tags:"句子"`
}

type SayingListRes struct {
	List  []entity.Saying `json:"list"`
	Total uint            `json:"total"`
}
