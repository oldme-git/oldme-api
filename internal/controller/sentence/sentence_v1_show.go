package sentence

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/sentence"

	"github.com/oldme-git/oldme-api/api/sentence/v1"
)

func (c *ControllerV1) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	data, err := sentence.Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	tagList, err := sentence.ShowTag(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.ShowRes{
		Sentence: data,
		TagList:  tagList,
	}, nil
}
