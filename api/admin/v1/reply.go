package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type ReplyCreReq struct {
	g.Meta `path:"reply/create" method:"post" sm:"新增" tags:"文章回复"`
	*model.ReplyInput
}

type ReplyCreRes struct {
}

type ReplyUpdReq struct {
	g.Meta `path:"reply/update/{id}" method:"post" sm:"修改" tags:"文章回复"`
	*model.IdInput
	*model.ReplyBody
}

type ReplyUpdRes struct {
}

type ReplyDelReq struct {
	g.Meta `path:"reply/delete/{id}" method:"post" sm:"删除" tags:"文章回复"`
	*model.IdInput
}

type ReplyDelRes struct {
}

type ReplyListReq struct {
	g.Meta `path:"reply/list" method:"get" sm:"查询列表" tags:"文章回复"`
	*model.ReplyQuery
}

type ReplyListRes struct {
	List  *[]entity.Reply `json:"list"`
	Total uint            `json:"total"`
}

type ReplyShowReq struct {
	g.Meta `path:"reply/show/{id}" method:"get" sm:"查询详情" tags:"文章回复"`
	*model.IdInput
}

type ReplyShowRes struct {
	*model.ReplyShow
}

type ReplyCheckReq struct {
	g.Meta `path:"reply/check/{id}" method:"post" sm:"审核" tags:"文章回复"`
	*model.IdInput
	Result bool   `json:"result" v:"required" dc:"审核结果：true成功，false失败"`
	Reason string `json:"reason" v:"length:1, 200" dc:"失败原因"`
}

type ReplyCheckRes struct {
}
