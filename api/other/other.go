// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package other

import (
	"context"

	"github.com/oldme-git/oldme-api/api/other/app"
)

type IOtherApp interface {
	Saying(ctx context.Context, req *app.SayingReq) (res *app.SayingRes, err error)
	Link(ctx context.Context, req *app.LinkReq) (res *app.LinkRes, err error)
	Reading(ctx context.Context, req *app.ReadingReq) (res *app.ReadingRes, err error)
	Poem(ctx context.Context, req *app.PoemReq) (res *app.PoemRes, err error)
}
