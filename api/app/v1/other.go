package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"oldme-api/internal/model/entity"
)

type SayingReq struct {
	g.Meta `path:"saying" method:"get" sm:"随机查询一条句子" tags:"app"`
}

type SayingRes struct {
	Saying string
}

type LinkReq struct {
	g.Meta `path:"link" method:"get" sm:"查询友链" tags:"app"`
}

type LinkRes struct {
	List  *[]entity.Link `json:"list"`
	Total uint           `json:"total"`
}
