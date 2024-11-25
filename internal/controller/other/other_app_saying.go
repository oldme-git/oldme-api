package other

import (
	"context"

	"github.com/oldme-git/oldme-api/api/other/app"
)

func (c *ControllerApp) Saying(ctx context.Context, req *app.SayingReq) (res *app.SayingRes, err error) {
	// sayingOne, err := saying.Show(ctx)
	// if err == nil {
	// 	res = &app.SayingRes{Saying: sayingOne}
	// }
	return
}
