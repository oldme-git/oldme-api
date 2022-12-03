package admin

import (
	"context"
	v1 "oldme-api/api/v1"
	"oldme-api/internal/model"
	"oldme-api/internal/service"
)

var File = cFile{}

type cFile struct {
}

func (c *cFile) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	var info *model.FileInfo
	info, err = service.File().Upload(ctx, req.Name)
	if err == nil {
		res = &v1.UploadRes{
			FileInfo: *info,
		}
	}
	return
}
