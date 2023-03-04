package article_grp

import (
	"context"
	"math/rand"
	"oldme-api/internal/dao"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/service"
)

type sSaying struct {
}

func init() {
	service.RegisterSaying(&sSaying{})
}

// Cre 创建句子
func (s *sSaying) Cre(ctx context.Context, saying string) (err error) {
	_, err = dao.Saying.Ctx(ctx).Data(do.Saying{
		Saying: saying,
	}).Insert()
	return
}

// Upd 更新句子
func (s *sSaying) Upd(ctx context.Context, id uint, saying string) (err error) {
	_, err = dao.Saying.Ctx(ctx).Data(do.Saying{
		Saying: saying,
	}).Where("id", id).Update()
	return
}

// Del 删除句子
func (s *sSaying) Del(ctx context.Context, id uint) (err error) {
	_, err = dao.Saying.Ctx(ctx).Where("id", id).Delete()
	return
}

// List 读取句子列表
func (s *sSaying) List(ctx context.Context) (list *[]entity.Saying, err error) {
	list = &[]entity.Saying{}
	res, err := dao.Saying.Ctx(ctx).All()
	_ = res.Structs(list)
	return
}

// Show 随机读取一条句子
func (s *sSaying) Show(ctx context.Context) (saying string, err error) {
	// 不直接使用一条sql查询，先聚合查询在通过go取随机数读取随机数据
	count, _ := dao.Saying.Ctx(ctx).Count()
	v, err := dao.Saying.Ctx(ctx).Fields("saying").Limit(rand.Intn(int(count))-1, 1).Value()
	if err != nil {
		return
	}
	saying = v.String()
	return
}
