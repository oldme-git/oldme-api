package file

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"oldme-api/internal/model"
	"oldme-api/internal/packed"
	"oldme-api/internal/service"
	"path"
	"strings"
	"time"
)

type sFile struct {
}

func init() {
	service.RegisterFile(&sFile{})
}

// Upload 上传文件到临时库
func (s *sFile) Upload(ctx context.Context, file *ghttp.UploadFile) (info *model.FileInfo, err error) {
	var (
		name    string
		lib     = "tmp"
		dirPath = getDir(ctx, lib)
		urlPath string
	)
	name, err = file.Save(dirPath, true)
	if err != nil {
		return
	}
	urlPath = getUrlFile(ctx, "tmp", name)
	info = &model.FileInfo{
		Name: name,
		Url:  urlPath,
	}
	// 启动定时器，让临时图片只能存在30分钟
	gtimer.AddOnce(ctx, 30*time.Minute, func(ctx context.Context) {
		_ = service.File().Del(ctx, urlPath)
	})
	return
}

// Save 保存文件到正式库，如果src不是tmp库中的文件，则不做操作，原路返回
func (s *sFile) Save(ctx context.Context, src string, lib string) (info *model.FileInfo, err error) {
	// 先从src获取文件名称
	name, err := getTmpName(ctx, src)
	if err != nil {
		return &model.FileInfo{
			Name: path.Base(src),
			Url:  src,
		}, err
	}
	// 移动
	info, err = moveTmp(ctx, lib+"/"+gtime.Now().Format("Ym"), name)
	if err != nil {
		return &model.FileInfo{
			Name: path.Base(src),
			Url:  src,
		}, packed.Err.Skip(10503, err.Error())
	}
	return
}

// SaveImg 保存图片文件到img库
func (s *sFile) SaveImg(ctx context.Context, src string) (info *model.FileInfo, err error) {
	if err = packed.Ext.Img(src); err != nil {
		err = packed.Err.Skip(10501, err.Error())
		return
	}
	return service.File().Save(ctx, src, "img")
}

// Del 删除
func (s *sFile) Del(ctx context.Context, src string) (err error) {
	conf, _ := getConf(ctx)
	// url换成dir形式
	src = strings.Replace(src, conf["url"], conf["dir"], 1)
	if !gfile.IsFile(src) {
		err = packed.Err.Skip(10503)
		return
	}
	err = gfile.Remove(src)
	if err != nil {
		return
	}
	return
}

// getLib 从url或者dir获取库的名称
func getLib(ctx context.Context, src string) (lib string, err error) {
	conf, _ := getConf(ctx)
	var (
		dir = conf["dir"]
		url = conf["url"]
		str string
	)
	if strings.HasPrefix(src, dir) {
		dirLen := len(dir + "/")
		str = gstr.SubStrRune(src, dirLen)
	} else if strings.HasPrefix(src, url) {
		urlLen := len(url + "/")
		str = gstr.SubStrRune(src, urlLen)
	} else {
		err = errors.New("没有找到库")
		return
	}
	libSlice, err := gregex.MatchString(`(?i)\w+`, str)
	if err != nil {
		err = errors.New("没有找到库")
		return
	}
	lib = libSlice[0]
	return
}

// getName 从url或者dir读取库中文件名称
func getName(ctx context.Context, lib string, src string) (name string, err error) {
	lib = lib + "/"
	var (
		dirPath = getDir(ctx, lib)
		urlPath = getUrl(ctx, lib)
	)

	if strings.HasPrefix(src, dirPath) || strings.HasPrefix(src, urlPath) {
		name = path.Base(src)
	} else {
		err = errors.New(src + "不存在")
	}
	return
}

// getTmpName 从url或者dir读取临时库中文件名称
func getTmpName(ctx context.Context, src string) (name string, err error) {
	name, err = getName(ctx, "tmp", src)
	return
}

// moveTmp tmp库的文件移入到正式库
func moveTmp(ctx context.Context, lib string, name string) (info *model.FileInfo, err error) {
	var (
		srcPath = getDirFile(ctx, "tmp", name)
		dstPath = getDirFile(ctx, lib, name)
	)
	if ok := gfile.IsFile(srcPath); !ok {
		err = errors.New(name + "不存在")
		return
	}
	err = gfile.Copy(srcPath, dstPath)
	if err != nil {
		return
	}
	// 删除原文件
	_ = gfile.Remove(srcPath)
	info = &model.FileInfo{
		Name: name,
		Url:  getUrlFile(ctx, lib, name),
		Dir:  dstPath,
	}
	return
}

// getConf 获取upload配置
func getConf(ctx context.Context) (upload map[string]string, err error) {
	cfg, _ := gcfg.New()
	uploadVar, err := cfg.Get(ctx, "upload")
	if err != nil {
		return
	}
	upload = uploadVar.MapStrStr()
	return
}

// getDir 获取disk dir路径
func getDir(ctx context.Context, lib string) string {
	conf, err := getConf(ctx)
	if err != nil {
		panic(err)
	}
	return conf["dir"] + "/" + lib
}

// getUrl 获取http url路径
func getUrl(ctx context.Context, lib string) string {
	conf, err := getConf(ctx)
	if err != nil {
		panic(err)
	}
	return conf["url"] + "/" + lib
}

// getDirFile 获取文件的disk dir路径
func getDirFile(ctx context.Context, lib string, name string) string {
	dir := getDir(ctx, lib)
	return dir + "/" + name
}

// getUrlFile 获取文件的http url路径
func getUrlFile(ctx context.Context, lib string, name string) string {
	url := getUrl(ctx, lib)
	return url + "/" + name
}
