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
		Order:       in.Order,
	}).Insert()
	if err != nil {
		return packed.Err.SysDb("insert", "article_grp")
	}
	return
}

// Upd 更新文章分类
func (s *sArticleGrp) Upd(ctx context.Context, id model.Id, in *model.ArticleGrpInput) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Data(do.ArticleGrp{
		Name:        in.Name,
		Tags:        in.Tags,
		Description: in.Description,
		Onshow:      in.Onshow,
		Order:       in.Order,
	}).Where("id", id).Update()
	if err != nil {
		return packed.Err.SysDb("update", "article_grp")
	}
	return
}

// Del 删除文章分类
func (s *sArticleGrp) Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Delete()
	// 软删除掉该分类下的文章
	var list []model.ArticleList
	res, err := dao.Article.Ctx(ctx).Where("grp_id", id).All()
	_ = res.Structs(&list)
	for _, v := range list {
		_ = service.Article().Del(ctx, model.Id(v.Id), false)
	}
	return
}

// List 读取文章分类列表
func (s *sArticleGrp) List(ctx context.Context) (list []entity.ArticleGrp, err error) {
	res, err := dao.ArticleGrp.Ctx(ctx).Order("order desc").All()
	_ = res.Structs(&list)
	return
}

// ListArticleCount 获取已经发布文章的文章分类列表统计
func (s *sArticleGrp) ListArticleCount(ctx context.Context) (map[uint]uint, error) {
	listCount, err := dao.Article.Ctx(ctx).
		Fields("count(*) count,grp_id").
		Where("onshow", 1).
		Group("grp_id").All()
	if err != nil {
		return nil, packed.Err.SysDb("select", "article")
	}
	idCountMap := make(map[uint]uint, len(listCount))
	for _, v := range listCount {
		idCountMap[v["grp_id"].Uint()] = v["count"].Uint()
	}

	return idCountMap, nil
}

// Show 读取文章分类详情
func (s *sArticleGrp) Show(ctx context.Context, id model.Id) (info *entity.ArticleGrp, err error) {
	err = dao.ArticleGrp.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = packed.Err.Skip(10100)
	}
	return
}

// IsExist 根据id判断一个文章分类是否存在
func (s *sArticleGrp) IsExist(ctx context.Context, id model.Id) bool {
	num, _ := dao.ArticleGrp.Ctx(ctx).Where("id", id).Count()
	return num == 1
}
