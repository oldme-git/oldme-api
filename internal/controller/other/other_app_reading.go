package other

import (
	"context"
	"time"

	"github.com/oldme-git/oldme-api/internal/logic/reading"

	"github.com/oldme-git/oldme-api/api/other/app"
)

func (c *ControllerApp) Reading(ctx context.Context, req *app.ReadingReq) (res *app.ReadingRes, err error) {
	list, err := reading.List(ctx)
	if err != nil {
		return nil, err
	}
	res = new(app.ReadingRes)
	for _, v := range list {
		res.List = append(res.List, app.ReadingList{
			Id:         v.Id,
			Name:       v.Name,
			Author:     v.Author,
			Status:     v.Status,
			FinishedAt: v.FinishedAt.Layout(time.DateOnly),
		})
	}
	return
}
