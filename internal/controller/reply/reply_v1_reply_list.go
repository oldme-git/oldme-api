package reply

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reply/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reply"
)

func (c *ControllerV1) ReplyList(ctx context.Context, req *v1.ReplyListReq) (*v1.ReplyListRes, error) {
	list, total, err := reply.List(ctx, req.ReplyQuery)
	if err != nil {
		return nil, err
	}
	// 查询数据表里总共的数据条数
	return &v1.ReplyListRes{
		List:  list,
		Total: total,
	}, nil
}
