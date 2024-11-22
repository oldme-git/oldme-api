package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type (
	ReadingReq struct {
		g.Meta `path:"reading" method:"get" sm:"查询阅读列表" tags:"reading"`
	}

	ReadingRes struct {
		List []entity.Reading `json:"list"`
	}
)
