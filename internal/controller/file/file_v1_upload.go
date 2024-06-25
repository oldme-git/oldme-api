package file

import (
	"context"

	"github.com/oldme-git/oldme-api/api/file/v1"
	"github.com/oldme-git/oldme-api/internal/logic/file"
	"github.com/oldme-git/oldme-api/internal/model"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	var info *model.FileInfo
	info, err = file.Upload(ctx, req.File)
	if err == nil {
		res = &v1.UploadRes{
			FileInfo: *info,
		}
	}
	return
}
