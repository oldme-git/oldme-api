package sentence

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/sentence"
	"github.com/oldme-git/oldme-api/internal/model"

	"github.com/oldme-git/oldme-api/api/sentence/v1"
)

func (c *ControllerV1) Upd(ctx context.Context, req *v1.UpdReq) (res *v1.UpdRes, err error) {
	err = sentence.Upd(ctx, &model.SentenceInput{
		Id:       req.Id,
		BookId:   req.BookId,
		Sentence: req.Sentence,
		TagIds:   req.TagIds,
	})
	return
}
