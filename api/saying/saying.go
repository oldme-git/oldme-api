// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package saying

import (
	"context"

	"github.com/oldme-git/oldme-api/api/saying/v1"
)

type ISayingV1 interface {
	SayingCre(ctx context.Context, req *v1.SayingCreReq) (res *v1.SayingCreRes, err error)
	SayingUpd(ctx context.Context, req *v1.SayingUpdReq) (res *v1.SayingUpdRes, err error)
	SayingDel(ctx context.Context, req *v1.SayingDelReq) (res *v1.SayingDelRes, err error)
	SayingList(ctx context.Context, req *v1.SayingListReq) (res *v1.SayingListRes, err error)
}
