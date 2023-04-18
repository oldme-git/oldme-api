package article_grp

import (
	"context"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/packed"
	"oldme-api/internal/service"
)

type sArticleGrp struct {
}

func init() {
	service.RegisterArticleGrp(&sArticleGrp{})
}

// Cre 创建文章分类
func (s *sArticleGrp) Cre(ctx context.Context, in *model.ArticleGrpInput) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Data(do.ArticleGrp{
		Name:        in.Name,
		Tags:        in.Tags,
		Description: in.Description,
		Onshow:      in.Onshow,
	}).Insert()
	return
}

// Upd 更新文章分类
func (s *sArticleGrp) Upd(ctx context.Context, id uint, in *model.ArticleGrpInput) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Data(do.ArticleGrp{
		Name:        in.Name,
		Tags:        in.Tags,
		Description: in.Description,
		Onshow:      in.Onshow,
	}).Where("id", id).Update()
	return
}

// Del 删除文章分类
func (s *sArticleGrp) Del(ctx context.Context, id uint) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Delete()
	// 软删除掉该分类下的文章
	var list []model.ArticleList
	res, err := dao.Article.Ctx(ctx).Where("grp_id", id).All()
	_ = res.Structs(&list)
	for _, v := range list {
		_ = service.Article().Del(ctx, v.Id, false)
	}
	return
}

// List 读取文章分类列表
func (s *sArticleGrp) List(ctx context.Context) (list *[]entity.ArticleGrp, err error) {
	list = &[]entity.ArticleGrp{}
	res, err := dao.ArticleGrp.Ctx(ctx).All()
	_ = res.Structs(list)
	return
}

// Show 读取文章分类详情
func (s *sArticleGrp) Show(ctx context.Context, id uint) (info *entity.ArticleGrp, err error) {
	err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = packed.Err.Skip(10100)
	}
	return
}

// IsExist 根据id判断一个文章分类是否存在
func (s *sArticleGrp) IsExist(ctx context.Context, id uint) bool {
	num, _ := dao.ArticleGrp.Ctx(ctx).Where("id", id).Count()
	return num == 1
}
