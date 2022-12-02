package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"oldme-api/internal/model"
)

type UploadImgReq struct {
	g.Meta `path:"file/upload/img" mime:"multipart/form-data" method:"post" sm:"上传图片" tags:"文件"`
	Img    *ghttp.UploadFile `json:"img" type:"file" dc:"选择图片上传"`
}

type UploadImgRes struct {
	model.FileInfo
}
