package sentence

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/sentence"
	"github.com/oldme-git/oldme-api/internal/model"

	"github.com/oldme-git/oldme-api/api/sentence/v1"
)

func (c *ControllerV1) Cre(ctx context.Context, req *v1.CreReq) (res *v1.CreRes, err error) {
	err = sentence.Cre(ctx, &model.SentenceInput{
		BookId:   req.BookId,
		TagIds:   req.TagIds,
		Sentence: req.Sentence,
	})
	return
}
