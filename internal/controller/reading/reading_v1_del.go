package reading

import (
	"context"

	"github.com/oldme-git/oldme-api/api/reading/v1"
	"github.com/oldme-git/oldme-api/internal/logic/reading"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = reading.Del(ctx, req.Id)
	return nil, err
}
