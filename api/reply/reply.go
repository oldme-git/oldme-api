// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package reply

import (
	"context"
	
	"github.com/oldme-git/oldme-api/api/reply/app"
	"github.com/oldme-git/oldme-api/api/reply/v1"
)

type IReplyApp interface {
	Reply(ctx context.Context, req *app.ReplyReq) (res *app.ReplyRes, err error)
}

type IReplyV1 interface {
	ReplyCre(ctx context.Context, req *v1.ReplyCreReq) (res *v1.ReplyCreRes, err error)
	ReplyUpd(ctx context.Context, req *v1.ReplyUpdReq) (res *v1.ReplyUpdRes, err error)
	ReplyDel(ctx context.Context, req *v1.ReplyDelReq) (res *v1.ReplyDelRes, err error)
	ReplyList(ctx context.Context, req *v1.ReplyListReq) (res *v1.ReplyListRes, err error)
	ReplyShow(ctx context.Context, req *v1.ReplyShowReq) (res *v1.ReplyShowRes, err error)
	ReplyCheck(ctx context.Context, req *v1.ReplyCheckReq) (res *v1.ReplyCheckRes, err error)
}


