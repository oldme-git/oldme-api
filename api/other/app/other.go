package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type (
	SayingReq struct {
		g.Meta `path:"saying" method:"get" sm:"随机查询一条句子" tags:"app"`
	}

	SayingRes struct {
		Saying string
	}
)

type (
	LinkReq struct {
		g.Meta `path:"link" method:"get" sm:"查询友链" tags:"app"`
	}

	LinkRes struct {
		List  []entity.Link `json:"list"`
		Total uint          `json:"total"`
	}
)

type (
	ReadingReq struct {
		g.Meta `path:"reading" method:"get" sm:"查询阅读列表" tags:"app"`
	}

	ReadingRes struct {
		List []ReadingList `json:"list"`
	}

	ReadingList struct {
		Id         uint   `json:"id"`
		Name       string `json:"name" sm:"书名"`
		Author     string `json:"author" sm:"作者"`
		Status     uint   `json:"status" sm:"状态 1-完结 2-在读 3-弃读"`
		FinishedAt string `json:"finish_at" sm:"完结时间"`
	}
)
