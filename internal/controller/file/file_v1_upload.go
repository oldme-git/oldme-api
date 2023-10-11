package file

import (
	"context"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/service"

	"github.com/oldme-git/oldme-api/api/file/v1"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	var info *model.FileInfo
	info, err = service.File().Upload(ctx, req.File)
	if err == nil {
		res = &v1.UploadRes{
			FileInfo: *info,
		}
	}
	return
}
