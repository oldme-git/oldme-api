package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
)

type ReplyReq struct {
	g.Meta `path:"reply" method:"post" sm:"文章回复" tags:"app"`
	*model.ReplyInputApp
}

type ReplyRes struct {
}
