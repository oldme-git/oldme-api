package admin

import (
	"context"
	"github.com/oldme-git/oldme-api/api/admin/v1"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/service"
)

var File = cFile{}

type cFile struct {
}

func (c *cFile) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	var info *model.FileInfo
	info, err = service.File().Upload(ctx, req.File)
	if err == nil {
		res = &v1.UploadRes{
			FileInfo: *info,
		}
	}
	return
}
