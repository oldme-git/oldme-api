package other

import (
	"context"

	"github.com/oldme-git/oldme-api/api/other/app"
	"github.com/oldme-git/oldme-api/internal/logic/sentence"
	"github.com/oldme-git/oldme-api/utility/uinit"
)

func (c *ControllerApp) Saying(ctx context.Context, req *app.SayingReq) (res *app.SayingRes, err error) {
	text, err := sentence.Text(ctx, uinit.SayingTagId)
	if err != nil {
		return nil, err
	}
	return &app.SayingRes{
		Saying: text.Sentence,
	}, nil
}
