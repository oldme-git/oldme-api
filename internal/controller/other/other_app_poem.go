package other

import (
	"context"

	"github.com/oldme-git/oldme-api/api/other/app"
	"github.com/oldme-git/oldme-api/internal/logic/sentence"
)

func (c *ControllerApp) Poem(ctx context.Context, req *app.PoemReq) (res *app.PoemRes, err error) {
	poem, err := sentence.Poem(ctx)
	if err != nil {
		return nil, err
	}
	return &app.PoemRes{
		Text: poem.Sentence,
	}, nil
}
