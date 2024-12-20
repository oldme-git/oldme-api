package file

import (
	"context"

	"github.com/oldme-git/oldme-api/api/file/v1"
	"github.com/oldme-git/oldme-api/internal/logic/file"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (*v1.UploadRes, error) {
	info, err := file.Upload(ctx, req.File)
	if err != nil {
		return nil, err
	}
	return &v1.UploadRes{
		FileInfo: *info,
	}, nil
}
