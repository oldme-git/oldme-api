package file

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"oldme-api/internal/model"
	"oldme-api/internal/service"
)

type sFile struct {
}

func init() {
	service.RegisterFile(&sFile{})
}

// Upload 上传文件到临时库
func (s *sFile) Upload(ctx context.Context, file *ghttp.UploadFile) (info *model.FileInfo, err error) {
	path := service.File().GetPath(ctx, "tmp")
	name, err := file.Save(path, true)
	if err != nil {
		return
	}
	url := service.File().GetUrl(ctx, "tmp", name)
	info = &model.FileInfo{
		Name: name,
		Url:  url,
	}
	return
}

// MoveTmp 移动tmp中的文件到正式目录中
func (s *sFile) MoveTmp(ctx context.Context, dir string, file string) (info *model.FileInfo, err error) {
	var (
		srcPath = service.File().GetPath(ctx, "tmp") + "/" + file
		dstPath = service.File().GetPath(ctx, dir) + "/" + file
	)
	err = gfile.Copy(srcPath, dstPath)
	if err != nil {
		return
	}
	// 删除原文件
	_ = gfile.Remove(srcPath)
	info = &model.FileInfo{
		Name: file,
		Url:  service.File().GetUrl(ctx, dir, file),
	}
	return
}

// GetConf 获取upload配置
func (s *sFile) GetConf(ctx context.Context) (upload map[string]string, err error) {
	cfg, _ := gcfg.New()
	uploadVar, err := cfg.Get(ctx, "upload")
	if err != nil {
		return
	}
	upload = uploadVar.MapStrStr()
	return
}

// GetPath 获取文件保存的具体目录，例如：“path/tmp”
func (s *sFile) GetPath(ctx context.Context, srcDir string) string {
	conf, err := service.File().GetConf(ctx)
	if err != nil {
		panic(err)
	}
	return conf["path"] + "/" + srcDir
}

// GetUrl 获取静态资源路径
func (s *sFile) GetUrl(ctx context.Context, dir string, file string) string {
	conf, err := service.File().GetConf(ctx)
	if err != nil {
		panic(err)
	}
	return conf["prefix"] + "/" + dir + "/" + file
}
