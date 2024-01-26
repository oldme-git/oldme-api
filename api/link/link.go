// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package link

import (
	"context"

	"github.com/oldme-git/oldme-api/api/link/v1"
)

type ILinkV1 interface {
	LinkCre(ctx context.Context, req *v1.LinkCreReq) (res *v1.LinkCreRes, err error)
	LinkUpd(ctx context.Context, req *v1.LinkUpdReq) (res *v1.LinkUpdRes, err error)
	LinkDel(ctx context.Context, req *v1.LinkDelReq) (res *v1.LinkDelRes, err error)
	LinkList(ctx context.Context, req *v1.LinkListReq) (res *v1.LinkListRes, err error)
}
