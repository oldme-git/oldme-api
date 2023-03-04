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
		LastedAt:    gtime.Now(),
	}).Insert()
	id, _ := res.LastInsertId()
	return uint(id), err
}

// Upd 更新文章
func (s *sArticle) Upd(ctx context.Context, id uint, in *model.ArticleInput) (err error) {
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
	}).Where("id", id).Update()
	return
}

// List 读取文章列表
func (s *sArticle) List(ctx context.Context, query *model.ArticleQuery) (list *[]model.ArticleList, total uint, err error) {
	// 对于查询初始值的处理
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Size == 0 {
		query.Size = 15
	}
	// 组成查询链
	db := dao.Article.Ctx(ctx)
	// 是否查询指定的grpId
	if query.GrpId != 0 {
		db = db.Where("grp_id", query.GrpId)
	}
	// 0 查询所有文章 1 查询只发布的文章
	if query.Onshow == 1 {
		db = db.Where("onshow", 1)
	}
	// 搜索文本
	if len(query.Search) != 0 {
		db = db.Where("title like ?", "%"+query.Search+"%").
			WhereOr("tags like ?", "%"+query.Search+"%").
			WhereOr("description like ?", "%"+query.Search+"%")
	}
	db = db.Order("created_at desc, id desc")
	// 是否查询删除掉的文章
	if query.IsDel {
		db = db.Unscoped().Where("deleted_at is not null")
	}

	data, err := db.Page(query.Page, query.Size).All()
	totalInt, _ := db.Ctx(ctx).Count()
	total = uint(totalInt)
	if err != nil {
		return
	}
	list = &[]model.ArticleList{}
	_ = data.Structs(list)
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

// Del 删除文章
func (s *sArticle) Del(ctx context.Context, id uint, isReal bool) (err error) {
	if isReal {
		info, err := service.Article().Show(ctx, id)
		if err != nil {
			return err
		}
		if info != nil {
			// 删除文件资源
			if len(info.Thumb) > 0 {
				_ = service.File().Del(ctx, info.Thumb)
			}
			if len(info.Content) > 0 {
				_ = service.Rich().Del(ctx, &info.Content)
			}
			_, err = dao.Article.Ctx(ctx).Where("id", id).Unscoped().Delete()
		}
	} else {
		_, err = dao.Article.Ctx(ctx).Where("id", id).Delete()
	}
	return
}

// ReCre 重新创建已经删除的文章
func (s *sArticle) ReCre(ctx context.Context, id uint) (err error) {
	_, err = dao.Article.Ctx(ctx).Where("id", id).Unscoped().Update("deleted_at=null")
	return
}
