package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
)

type CreReq struct {
	g.Meta `path:"tag/create" method:"post" sm:"新增" tags:"标签"`
	GrpId  model.Id `json:"grpId" v:"required|integer|between:1,4294967295"`
	Name   string   `json:"name" v:"required|length:1, 30"`
}

type CreRes struct {
}

type UpdReq struct {
	g.Meta `path:"tag/update/{id}" method:"post" sm:"修改" tags:"标签"`
	*model.IdInput
	GrpId model.Id `json:"grpId" v:"required|integer|between:1,4294967295"`
	Name  string   `json:"name" v:"required|length:1, 30"`
}

type UpdRes struct {
}

type DelReq struct {
	g.Meta `path:"tag/delete/{id}" method:"post" sm:"删除" tags:"标签"`
	*model.IdInput
}

type DelRes struct {
}

type ListReq struct {
	g.Meta `path:"tag/list" method:"get" sm:"查询列表" tags:"标签"`
	GrpId  model.Id `json:"grpId" v:"required|integer|between:1,4294967295"`
}

type ListRes struct {
	List []List `json:"list"`
}
