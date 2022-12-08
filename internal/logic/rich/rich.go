package rich

import (
	"context"
	"oldme-api/internal/service"
	"oldme-api/utility"
	"oldme-api/utility/uregex"
	"strings"
)

type sRich struct {
}

func init() {
	service.RegisterRich(&sRich{})
}

const (
	imgTmpPattern = `(?i)(https?)://(((?!http).)*)tmp(((?!http).)*)?(.png|.jpg|.jpeg)`
	imgPattern    = `(?i)(https?)://((?!http).)*?(.png|.jpg|.jpeg)`
)

// Save 提取富文本中的tmpimg，移入到正式库
func (s *sRich) Save(ctx context.Context, text *string) {
	// 正则所有的img：png、jpg、jpeg
	match, _ := uregex.MatchAllString(imgTmpPattern, *text)
	// 获取所有图片的链接
	for _, v := range match {
		info, err := service.File().Save(ctx, v, "rich")
		if err == nil {
			// 替换富文本中的内存
			*text = strings.Replace(*text, v, info.Url, -1)
		}
	}
	return
}

func (s *sRich) Edit(ctx context.Context, oldText *string, newText *string) {
	// 新的富文本内存需要先保存
	service.Rich().Save(ctx, newText)

	// 正则出新旧的图片切片
	oldSlice, _ := uregex.MatchAllString(imgPattern, *oldText)
	newSlice, _ := uregex.MatchAllString(imgPattern, *newText)

	for _, v := range oldSlice {
		// 在newSlice的去除
		if ok, _ := utility.InArray(v, newSlice); !ok {
			_ = service.File().Del(ctx, v)
		}
	}
}

// Del 删除
func (s *sRich) Del(ctx context.Context, text *string) (err error) {
	// 正则所有的img：png、jpg、jpeg
	match, err := uregex.MatchAllString(imgPattern, *text)
	if err != nil {
		panic(err)
	}
	// 获取所有图片的链接
	for _, v := range match {
		info, err := service.File().Save(ctx, v, "rich")
		if err == nil {
			// 替换富文本中的内存
			*text = strings.Replace(*text, v, info.Url, -1)
		}
	}
	return
}
