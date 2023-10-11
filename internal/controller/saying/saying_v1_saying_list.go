package saying

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/saying/v1"
)

func (c *ControllerV1) SayingList(ctx context.Context, req *v1.SayingListReq) (res *v1.SayingListRes, err error) {
	list, err := service.Saying().List(ctx)
	if err == nil {
		res = &v1.SayingListRes{
			List:  list,
			Total: uint(len(list)),
		}
	}
	return
}
