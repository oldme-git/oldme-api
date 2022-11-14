package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"oldme-api/internal/model/entity"
)

type DemoCreateReq struct {
	g.Meta   `path:"/demo/create" method:"post" sm:"示例增加数据" tags:"示例"`
	Username string
	Password string
}

type DemoCreateRes struct {
}

type DemoUpdateReq struct {
	g.Meta   `path:"/demo/update" method:"post" sm:"示例更新数据" tags:"示例"`
	Username string
	Password string
}

type DemoUpdateRes struct {
}

type DemoReadReq struct {
	g.Meta   `path:"/demo/read" method:"get" sm:"示例查询数据" tags:"示例"`
	Username string `v:"required|length:2,20"`
}

type DemoReadRes struct {
	*entity.Admin
}

type DemoDelReq struct {
	g.Meta `path:"/demo/del" method:"get" sm:"示例删除数据" tags:"示例"`
	Id     int
}

type DemoDelRes struct {
}
