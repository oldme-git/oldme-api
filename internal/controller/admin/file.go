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

func (c *cFile) UploadFile(ctx context.Context, req *v1.UploadImgReq) (res *v1.UploadImgRes, err error) {
	var info *model.FileInfo
	info, err = service.File().UploadImg(ctx, req.Img)
	if err == nil {
		res = &v1.UploadImgRes{
			FileInfo: *info,
		}
	}
	return
}
