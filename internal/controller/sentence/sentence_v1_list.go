package sentence

import (
	"context"

	"github.com/oldme-git/oldme-api/api/sentence/v1"
	"github.com/oldme-git/oldme-api/internal/logic/sentence"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	if req.Paging == nil {
		req.Paging = &model.Paging{
			Page: 1,
			Size: 15,
		}
	}

	// bookId 和 tagId 不能同时存在，tagId 优先级更高
	var (
		data  []entity.Sentence
		total uint
	)
	if len(req.TagIds) > 0 {
		data, total, err = listByTagIds(ctx, req)
	} else {
		data, total, err = listByBookId(ctx, req)
	}

	if err != nil {
		return nil, err
	}

	var list []v1.List
	for _, v := range data {
		list = append(list, v1.List{
			Id:       model.Id(v.Id),
			BookId:   model.Id(v.BookId),
			Sentence: v.Sentence,
		})
	}
	return &v1.ListRes{
		List:  list,
		Total: total,
	}, nil
}

func listByTagIds(ctx context.Context, req *v1.ListReq) (data []entity.Sentence, total uint, err error) {
	var ids []model.Id
	ids, total, err = sentence.GetIdsByTagIds(ctx, req.TagIds, *req.Paging)
	if err != nil {
		return
	}

	// 重置页码到1，使用新的ids控制分页
	req.Paging.Page = 1

	query := &model.SentenceQuery{
		Paging: *req.Paging,
		Ids:    ids,
	}

	data, _, err = sentence.List(ctx, query)
	return
}

func listByBookId(ctx context.Context, req *v1.ListReq) (data []entity.Sentence, total uint, err error) {
	query := &model.SentenceQuery{
		Paging: *req.Paging,
		BookId: req.BookId,
	}

	data, total, err = sentence.List(ctx, query)
	return
}
