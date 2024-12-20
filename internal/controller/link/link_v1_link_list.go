package link

import (
	"context"

	"github.com/oldme-git/oldme-api/api/link/v1"
	"github.com/oldme-git/oldme-api/internal/logic/link"
)

func (c *ControllerV1) LinkList(ctx context.Context, req *v1.LinkListReq) (res *v1.LinkListRes, err error) {
	list, err := link.List(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.LinkListRes{
		List:  list,
		Total: uint(len(list)),
	}, nil
}
