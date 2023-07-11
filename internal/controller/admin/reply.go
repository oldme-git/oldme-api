package admin

import (
	"context"
	v1 "oldme-api/api/admin/v1"
	"oldme-api/internal/service"
)

var Replay = &cReplay{}

type cReplay struct {
}

func (c *cReplay) Cre(ctx context.Context, req *v1.ReplyCreReq) (res *v1.ReplyCreRes, err error) {
	err = service.Reply().Cre(ctx, req.ReplyInput)
	return
}

func (c *cReplay) Upd(ctx context.Context, req *v1.ReplyUpdReq) (res *v1.ReplyUpdRes, err error) {
	err = service.Reply().Upd(ctx, req.Id, req.ReplyInput)
	return
}

func (c *cReplay) Show(ctx context.Context, req *v1.ReplyShowReq) (res *v1.ReplyShowRes, err error) {
	info, err := service.Reply().Show(ctx, req.Id)
	if err == nil {
		res = &v1.ReplyShowRes{
			ReplyShow: info,
		}
	}
	return
}

func (c *cReplay) Del(ctx context.Context, req *v1.ReplyDelReq) (res *v1.ReplyDelRes, err error) {
	err = service.Reply().Del(ctx, req.Id)
	return
}

func (c *cReplay) List(ctx context.Context, req *v1.ReplyListReq) (res *v1.ReplyListRes, err error) {
	list, total, err := service.Reply().List(ctx, req.ReplyQuery)
	if err == nil {
		// 查询数据表里总共的数据条数
		res = &v1.ReplyListRes{
			List:  list,
			Total: total,
		}
	}
	return
}

func (c *cReplay) Check(ctx context.Context, req *v1.ReplyCheckReq) (res *v1.ReplyCheckRes, err error) {
	err = service.Reply().Check(ctx, req.Id, req.Result, req.Reason)
	return
}
