package saying

import (
	"context"
	"math/rand"

	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/do"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

// Cre 创建句子
func Cre(ctx context.Context, saying string) (err error) {
	_, err = dao.Saying.Ctx(ctx).Data(do.Saying{
		Saying: saying,
	}).Insert()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// Upd 更新句子
func Upd(ctx context.Context, id model.Id, saying string) (err error) {
	_, err = dao.Saying.Ctx(ctx).Data(do.Saying{
		Saying: saying,
	}).Where("id", id).Update()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// Del 删除句子
func Del(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Saying.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		err = utility.Err.Sys(err)
	}
	return
}

// List 读取句子列表
func List(ctx context.Context) (list []entity.Saying, err error) {
	res, err := dao.Saying.Ctx(ctx).All()
	if err != nil {
		return nil, utility.Err.Sys(err)
	}
	_ = res.Structs(&list)
	return
}

// Show 随机读取一条句子
func Show(ctx context.Context) (saying string, err error) {
	// 不直接使用一条sql查询，先聚合查询在通过go取随机数读取随机数据
	count, _ := dao.Saying.Ctx(ctx).Count()
	if count == 0 {
		return
	}
	v, err := dao.Saying.Ctx(ctx).Fields("saying").Limit(rand.Intn(int(count))-1, 1).Value()
	if err != nil {
		err = utility.Err.Sys(err)
		return
	}
	saying = v.String()
	return
}
