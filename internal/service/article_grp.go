// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"oldme-api/internal/model"
	"oldme-api/internal/model/entity"
)

type (
	IArticleGrp interface {
		// Cre 创建文章分类
		Cre(ctx context.Context, in *model.ArticleGrpInput) (err error)
		// Upd 更新文章分类
		Upd(ctx context.Context, id model.Id, in *model.ArticleGrpInput) (err error)
		// Del 删除文章分类
		Del(ctx context.Context, id model.Id) (err error)
		// List 读取文章分类列表
		List(ctx context.Context) (list []entity.ArticleGrp, err error)
		// ListArticleCount 获取已经发布文章的文章分类列表统计
		ListArticleCount(ctx context.Context) (map[uint]uint, error)
		// Show 读取文章分类详情
		Show(ctx context.Context, id model.Id) (info *entity.ArticleGrp, err error)
		// IsExist 根据id判断一个文章分类是否存在
		IsExist(ctx context.Context, id model.Id) bool
	}
)

var (
	localArticleGrp IArticleGrp
)

func ArticleGrp() IArticleGrp {
	if localArticleGrp == nil {
		panic("implement not found for interface IArticleGrp, forgot register?")
	}
	return localArticleGrp
}

func RegisterArticleGrp(i IArticleGrp) {
	localArticleGrp = i
}
