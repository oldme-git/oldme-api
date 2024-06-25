package article

import (
	"context"

	"github.com/oldme-git/oldme-api/api/article/app"
	"github.com/oldme-git/oldme-api/internal/logic/article"
)

func (c *ControllerApp) Hist(ctx context.Context, req *app.HistReq) (res *app.HistRes, err error) {
	err = article.Hist(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return
}
