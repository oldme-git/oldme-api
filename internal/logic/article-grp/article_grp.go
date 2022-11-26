package article_grp

import (
	"context"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/service"
)

type sArticleGrp struct {
}

func init() {
	service.RegisterArticleGrp(&sArticleGrp{})
}

// Cre 创建文章分类
func (s sArticleGrp) Cre(ctx context.Context, in model.ArticleGrpInput) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Data(do.ArticleGrp{
		Name:        in.Name,
		Tags:        in.Tags,
		Description: in.Description,
		Onshow:      in.Onshow,
	}).Insert()
	return
}

// Upd 更新文章分类
func (s sArticleGrp) Upd(ctx context.Context, id uint32, in model.ArticleGrpInput) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Data(do.ArticleGrp{
		Name:        in.Name,
		Tags:        in.Tags,
		Description: in.Description,
		Onshow:      in.Onshow,
	}).Where("id", id).Update()
	return
}

// Del 删除文章分类
func (s sArticleGrp) Del(ctx context.Context, id uint32) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Delete()
	return
}

// Read 读取文章分类
func (s sArticleGrp) Read(ctx context.Context, id uint32) (data entity.ArticleGrp, err error) {
	err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Scan(&data)
	return
}
