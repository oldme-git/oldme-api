package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"oldme-api/internal/model"
)

type UploadReq struct {
	g.Meta `path:"file/upload/img" mime:"multipart/form-data" method:"post" sm:"上传文件到临时库" tags:"文件"`
	File   *ghttp.UploadFile `json:"file" v:"required" type:"file" dc:"选择文件上传上传，Max:5M"`
}

type UploadRes struct {
	model.FileInfo
}
