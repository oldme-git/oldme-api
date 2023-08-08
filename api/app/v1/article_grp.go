package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/oldme-api/internal/model"
)

type ArticleGrpListReq struct {
	g.Meta `path:"article/group/list" method:"get" sm:"查询列表" tags:"app"`
}

type ArticleGrpListRes struct {
	List []model.ArticleGrpListSafe `json:"list"`
}
