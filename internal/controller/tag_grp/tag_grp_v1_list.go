package tag_grp

import (
	"context"

	"github.com/oldme-git/oldme-api/api/tag_grp/v1"
	"github.com/oldme-git/oldme-api/internal/logic/tag_grp"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	data, err := tag_grp.List(ctx)
	if err != nil {
		return nil, err
	}

	var list []v1.List
	for _, v := range data {
		list = append(list, v1.List{
			Id:   model.Id(v.Id),
			Name: v.Name,
		})
	}

	return &v1.ListRes{
		List: list,
	}, err
}
