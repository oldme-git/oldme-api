// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/app"
	"github.com/oldme-git/oldme-api/api/article/v1"
)

type IArticleApp interface {
	List(ctx context.Context, req *app.ListReq) (res *app.ListRes, err error)
	Rank(ctx context.Context, req *app.RankReq) (res *app.RankRes, err error)
	Show(ctx context.Context, req *app.ShowReq) (res *app.ShowRes, err error)
	About(ctx context.Context, req *app.AboutReq) (res *app.AboutRes, err error)
	Hist(ctx context.Context, req *app.HistReq) (res *app.HistRes, err error)
	Reply(ctx context.Context, req *app.ReplyReq) (res *app.ReplyRes, err error)
}

type IArticleV1 interface {
	Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error)
	Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error)
	ReCre(ctx context.Context, req *v1.ReCreReq) (res *v1.ReCreRes, err error)
}
