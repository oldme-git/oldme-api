package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type CreReq struct {
	g.Meta     `path:"reading/create" method:"post" sm:"新增" tags:"阅读"`
	Name       string              `json:"name" v:"required"`
	Author     string              `json:"author" v:"required"`
	Status     model.ReadingStatus `json:"status" v:"required"`
	FinishedAt *gtime.Time         `json:"finished_at" v:"required"`
}

type CreRes struct {
}

type UpdReq struct {
	g.Meta `path:"reading/update/{id}" method:"post" sm:"修改" tags:"阅读"`
	*model.IdInput
	Name       string              `json:"name" v:"required"`
	Author     string              `json:"author" v:"required"`
	Status     model.ReadingStatus `json:"status" v:"required"`
	FinishedAt *gtime.Time         `json:"finished_at" v:"required"`
}

type UpdRes struct {
}

type ShowReq struct {
	g.Meta `path:"reading/show/{id}" method:"get" sm:"查询详情" tags:"阅读"`
	*model.IdInput
}

type ShowRes struct {
	*entity.Reading
}

type DelReq struct {
	g.Meta `path:"reading/delete/{id}" method:"post" sm:"删除" tags:"阅读"`
	*model.IdInput
}

type DelRes struct {
}

type ListReq struct {
	g.Meta `path:"reading/list" method:"get" sm:"查询阅读列表" tags:"reading"`
	Search string `json:"search"  dc:"搜索文本"`
}

type ListRes struct {
	List []entity.Reading `json:"list"`
}
