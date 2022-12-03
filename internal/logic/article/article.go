package article

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/packed"
	"oldme-api/internal/service"
)

type sArticle struct {
}

func init() {
	service.RegisterArticle(&sArticle{})
}

// Cre 创建文章
func (s *sArticle) Cre(ctx context.Context, in *model.ArticleInput) (err error) {
	// 判断该分类是否存在
	if b := service.ArticleGrp().IsExist(ctx, in.GrpId); !b {
		return packed.Code.SetErr(10101)
	}
	// 处理文章thumb
	if err = packed.Ext.Img(in.Thumb); err != nil {
		return
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
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
		LastedAt:    gtime.Now(),
	}).Insert()
	return err
}
