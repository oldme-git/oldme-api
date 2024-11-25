package other

import (
	"context"

	"github.com/oldme-git/oldme-api/api/other/app"
	"github.com/oldme-git/oldme-api/internal/logic/sentence"
)

func (c *ControllerApp) Saying(ctx context.Context, req *app.SayingReq) (res *app.SayingRes, err error) {
	saying, err := sentence.Saying(ctx)
	if err != nil {
		return nil, err
	}
	return &app.SayingRes{
		Saying: saying.Sentence,
	}, nil
}
