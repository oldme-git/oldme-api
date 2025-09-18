package other

import (
	"context"

	"github.com/oldme-git/oldme-api/api/other/app"
	"github.com/oldme-git/oldme-api/internal/logic/sentence"
	"github.com/oldme-git/oldme-api/utility/uinit"
)

func (c *ControllerApp) Poem(ctx context.Context, req *app.PoemReq) (res *app.PoemRes, err error) {
	text, err := sentence.Text(ctx, uinit.PoemTagId)
	if err != nil {
		return nil, err
	}
	return &app.PoemRes{
		Poem: text.Sentence,
	}, nil
}
