package rich

import (
	"context"
	"github.com/gogf/gf/v2/text/gregex"
	"oldme-api/internal/service"
	"strings"
)

type sRich struct {
}

func init() {
	service.RegisterRich(&sRich{})
}

const imgPattern = `(?i)(https?)://.*?(.png|.jpg|.jpeg)`

// Save 提取富文本中的tmpimg，移入到正式库
func (s *sRich) Save(ctx context.Context, text *string) (err error) {
	// 正则所有的img：png、jpg、jpeg
	match, err := gregex.MatchAllString(imgPattern, *text)
	if err != nil {
		panic(err)
	}
	// 获取所有图片的链接
	for _, v := range match {
		info, err := service.File().Save(ctx, v[0], "rich")
		if err == nil {
			// 替换富文本中的内存
			*text = strings.Replace(*text, v[0], info.Url, -1)
		}
	}
	return
}

// Del 删除
func (s *sRich) Del(ctx context.Context, text *string) (err error) {
	// 正则所有的img：png、jpg、jpeg
	match, err := gregex.MatchAllString(imgPattern, *text)
	if err != nil {
		panic(err)
	}
	// 获取所有图片的链接
	for _, v := range match {
		info, err := service.File().Save(ctx, v[0], "rich")
		if err == nil {
			// 替换富文本中的内存
			*text = strings.Replace(*text, v[0], info.Url, -1)
		}
	}
	return
}
