package article

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/article/app"
)

func (c *ControllerApp) Hist(ctx context.Context, req *app.HistReq) (res *app.HistRes, err error) {
	err = service.Article().Hist(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return
}
