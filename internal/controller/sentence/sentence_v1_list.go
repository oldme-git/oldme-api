package sentence

import (
	"context"

	"github.com/oldme-git/oldme-api/api/sentence/v1"
	"github.com/oldme-git/oldme-api/internal/logic/sentence"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	if req.Paging == nil {
		req.Paging = &model.Paging{
			Page: 1,
			Size: 15,
		}
	}

	var ids []model.Id
	// 如果有tagId，根据tagId查询ids
	if len(req.TagIds) > 0 {
		ids, err = sentence.GetIdsByTagIds(ctx, req.TagIds, uint(req.Paging.Size))
		if err != nil {
			return nil, err
		}
	}

	query := &model.SentenceQuery{
		Paging: *req.Paging,
		BookId: req.BookId,
		Ids:    ids,
	}

	data, total, err := sentence.List(ctx, query)
	if err != nil {
		return nil, err
	}
	var (
		list []v1.List
		s    string
	)
	for _, v := range data {
		if len(v.Sentence) > 100 {
			s = v.Sentence[:100]
		} else {
			s = v.Sentence
		}
		list = append(list, v1.List{
			Id:       model.Id(v.Id),
			BookId:   model.Id(v.BookId),
			Sentence: s,
		})
	}
	return &v1.ListRes{
		List:  list,
		Total: total,
	}, nil
}
