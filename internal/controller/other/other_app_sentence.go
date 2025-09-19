package other

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/sentence"
	"github.com/oldme-git/oldme-api/utility/uinit"

	"github.com/oldme-git/oldme-api/api/other/app"
)

func (c *ControllerApp) Sentence(ctx context.Context, req *app.SentenceReq) (res *app.SentenceRes, err error) {
	poem, err := sentence.Text(ctx, uinit.PoemTagId)
	if err != nil {
		return nil, err
	}

	slang, err := sentence.Text(ctx, uinit.SlangTagId)
	if err != nil {
		return nil, err
	}

	return &app.SentenceRes{
		Poem:  poem.Sentence,
		Slang: slang.Sentence,
	}, nil

}
