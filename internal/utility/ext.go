package utility

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
)

type pExt struct {
}

var Ext pExt

// Img 限制图片类文件类型
func (e pExt) Img(imgPath string) error {
	ext := "png, jpg, jpeg, gif"
	var (
		pathExt  = gfile.ExtName(imgPath)
		sliceExt = strings.Split(ext, ",")
	)

	for _, item := range sliceExt {
		if strings.ToLower(pathExt) == strings.ToLower(item) {
			return nil
		}
	}
	return errors.New("图片类型不正确，应为" + ext)
}
