package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type CreReq struct {
	g.Meta `path:"tag/group/create" method:"post" sm:"新增" tags:"标签分类"`
	Name   string `json:"name" v:"required|length:1, 30"`
}

type CreRes struct {
}

type UpdReq struct {
	g.Meta `path:"tag/group/update/{id}" method:"post" sm:"修改" tags:"标签分类"`
	*model.IdInput
	Name string `json:"name" v:"required|length:1, 30"`
}

type UpdRes struct {
}

type DelReq struct {
	g.Meta `path:"tag/group/delete/{id}" method:"post" sm:"删除" tags:"标签分类"`
	*model.IdInput
}

type DelRes struct {
}

type ShowReq struct {
	g.Meta `path:"tag/group/show/{id}" method:"get" sm:"查询详情" tags:"标签分类"`
	*model.IdInput
}

type ShowRes struct {
	*entity.TagGrp
}

type ListReq struct {
	g.Meta `path:"tag/group/list" method:"get" sm:"查询列表" tags:"标签分类"`
}

type ListRes struct {
	List []List `json:"list"`
}
