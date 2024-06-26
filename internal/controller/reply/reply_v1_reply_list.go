package reply

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reply/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reply"
)

func (c *ControllerV1) ReplyList(ctx context.Context, req *v1.ReplyListReq) (res *v1.ReplyListRes, err error) {
	list, total, err := reply.List(ctx, req.ReplyQuery)
	if err == nil {
		// 查询数据表里总共的数据条数
		res = &v1.ReplyListRes{
			List:  list,
			Total: total,
		}
	}
	return
}
