package sentence

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/logic/sentence"

	"github.com/oldme-git/oldme-api/api/sentence/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = sentence.Del(ctx, req.Id)
	return
}
