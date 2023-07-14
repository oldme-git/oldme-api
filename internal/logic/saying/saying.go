package article_grp

import (
	"context"
	"math/rand"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/packed"
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
	if err != nil {
		err = packed.Err.SysDb("insert", "saying")
	}
	return
}

// Upd 更新句子
func (s *sSaying) Upd(ctx context.Context, id model.Id, saying string) (err error) {
	_, err = dao.Saying.Ctx(ctx).Data(do.Saying{
		Saying: saying,
	}).Where("id", id).Update()
	if err != nil {
		err = packed.Err.SysDb("update", "saying")
	}
	return
}

// Del 删除句子
func (s *sSaying) Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Saying.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		err = packed.Err.SysDb("delete", "saying")
	}
	return
}

// List 读取句子列表
func (s *sSaying) List(ctx context.Context) (list []entity.Saying, err error) {
	res, err := dao.Saying.Ctx(ctx).All()
	if err != nil {
		return nil, packed.Err.SysDb("select", "saying")
	}
	_ = res.Structs(&list)
	return
}

// Show 随机读取一条句子
func (s *sSaying) Show(ctx context.Context) (saying string, err error) {
	// 不直接使用一条sql查询，先聚合查询在通过go取随机数读取随机数据
	count, _ := dao.Saying.Ctx(ctx).Count()
	v, err := dao.Saying.Ctx(ctx).Fields("saying").Limit(rand.Intn(int(count))-1, 1).Value()
	if err != nil {
		err = packed.Err.SysDb("select", "saying")
		return
	}
	saying = v.String()
	return
}
