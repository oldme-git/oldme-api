package saying

import (
	"context"

	"github.com/oldme-git/oldme-api/api/saying/v1"
	"github.com/oldme-git/oldme-api/internal/logic/saying"
)

func (c *ControllerV1) SayingList(ctx context.Context, req *v1.SayingListReq) (res *v1.SayingListRes, err error) {
	list, err := saying.List(ctx)
	if err == nil {
		res = &v1.SayingListRes{
			List:  list,
			Total: uint(len(list)),
		}
	}
	return
}
