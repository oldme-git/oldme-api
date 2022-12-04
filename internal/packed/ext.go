package packed

import (
	"errors"
	"path"
	"strings"
)

type pExt struct {
}

var Ext pExt

// Img 限制图片类文件类型
func (e pExt) Img(imgPath string) error {
	ext := ".png, .jpg, .jpeg, .gif"
	var (
		pathExt  = path.Ext(imgPath)
		sliceExt = strings.Split(ext, ",")
	)

	for _, item := range sliceExt {
		if strings.ToLower(pathExt) == strings.ToLower(item) {
			return nil
		}
	}
	return errors.New("图片类型不正确，应为" + ext)
}
