package article

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/packed"
	"oldme-api/internal/service"
)

type sArticle struct {
}

func init() {
	service.RegisterArticle(&sArticle{})
}

// Cre 创建文章
func (s *sArticle) Cre(ctx context.Context, in *model.ArticleInput) (lastId uint, err error) {
	// 判断该分类是否存在
	if ok := service.ArticleGrp().IsExist(ctx, in.GrpId); !ok {
		err = packed.Code.SetErr(10101)
		return
	}
	// 保存thumb到正式目录
	if len(in.Thumb) != 0 {
		info, err := service.File().SaveImg(ctx, in.Thumb)
		if err == nil {
			in.Thumb = info.Url
		}
	}
	// 保存富文本
	if len(in.Content) != 0 {
		service.Rich().Save(ctx, &in.Content)
	}
	res, err := dao.Article.Ctx(ctx).Data(do.Article{
		GrpId:       in.GrpId,
		Title:       in.Title,
		Author:      in.Author,
		Thumb:       in.Thumb,
		Tags:        in.Tags,
		Description: in.Description,
		Content:     in.Content,
		Order:       in.Order,
		Ontop:       in.Ontop,
		Onshow:      in.Onshow,
		Hist:        in.Hist,
		Post:        in.Post,
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
		LastedAt:    gtime.Now(),
	}).Insert()
	id, _ := res.LastInsertId()
	return uint(id), err
}

// Upt 更新文章
func (s *sArticle) Upt(ctx context.Context, id uint, in *model.ArticleInput) (err error) {
	info, err := service.Article().Show(ctx, id)
	if err != nil {
		return
	}
	// 判断该分类是否存在
	if ok := service.ArticleGrp().IsExist(ctx, in.GrpId); !ok {
		err = packed.Code.SetErr(10101)
		return
	}
	// 保存thumb到正式目录
	if in.Thumb != info.Thumb {
		// 不等于要删除掉原图片
		_ = service.File().Del(ctx, info.Thumb)
		// 重新保存新图片
		info, _ := service.File().SaveImg(ctx, in.Thumb)
		if err == nil {
			in.Thumb = info.Url
		}
	}
	// 保存富文本
	if len(in.Content) != 0 {
		service.Rich().Edit(ctx, &info.Content, &in.Content)
	}
	_, err = dao.Article.Ctx(ctx).Data(do.Article{
		GrpId:       in.GrpId,
		Title:       in.Title,
		Author:      in.Author,
		Thumb:       in.Thumb,
		Tags:        in.Tags,
		Description: in.Description,
		Content:     in.Content,
		Order:       in.Order,
		Ontop:       in.Ontop,
		Onshow:      in.Onshow,
		Hist:        in.Hist,
		Post:        in.Post,
		UpdatedAt:   gtime.Now(),
	}).Where("id", id).Update()
	return
}

// List 读取文章列表
func (s *sArticle) List(ctx context.Context, grpId uint) (list *model.ArticleList, err error) {
	return
}

// Show 读取文章详情
func (s *sArticle) Show(ctx context.Context, id uint) (info *entity.Article, err error) {
	err = dao.Article.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = packed.Code.SetErr(10100)
	}
	return
}
