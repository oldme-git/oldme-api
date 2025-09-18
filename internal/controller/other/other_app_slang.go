package other

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/sentence"
	"github.com/oldme-git/oldme-api/utility/uinit"

	"github.com/oldme-git/oldme-api/api/other/app"
)

func (c *ControllerApp) Slang(ctx context.Context, req *app.SlangReq) (res *app.SlangRes, err error) {
	text, err := sentence.Text(ctx, uinit.SlangTagId)
	if err != nil {
		return nil, err
	}
	return &app.SlangRes{
		Slang: text.Sentence,
	}, nil
}
