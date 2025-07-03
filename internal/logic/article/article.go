package article

import (
	"context"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/oldme-git/oldme-api/internal/dao"
	"github.com/oldme-git/oldme-api/internal/logic/file"
	"github.com/oldme-git/oldme-api/internal/logic/rich"
	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/do"
	"github.com/oldme-git/oldme-api/internal/model/entity"
	"github.com/oldme-git/oldme-api/internal/utility"
)

// Cre 创建文章
func Cre(ctx context.Context, in *model.ArticleInput) (lastId model.Id, err error) {
	// 保存thumb到正式目录
	if in.Thumb != "" {
		info, err := file.SaveImg(ctx, in.Thumb)
		if err != nil {
			return 0, err
		} else {
			in.Thumb = info.Url
		}
	}
	// 保存富文本
	if in.Content != "" {
		rich.Save(ctx, &in.Content)
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
	if err != nil {
		return 0, utility.Err.Sys(err)
	}
	id, _ := res.LastInsertId()
	return model.Id(id), err
}

// Upd 更新文章
func Upd(ctx context.Context, id model.Id, in *model.ArticleInput) (err error) {
	info, err := Show(ctx, id)
	if err != nil {
		return
	}

	// 保存thumb到正式目录
	if in.Thumb == "" {
		// 删除原图片
		_ = file.Del(ctx, info.Thumb)
	} else {
		if in.Thumb != info.Thumb {
			// 重新保存新图片
			imgInfo, err := file.SaveImg(ctx, in.Thumb)
			if err != nil {
				return err
			} else {
				in.Thumb = imgInfo.Url
				// 删除原图片
				_ = file.Del(ctx, info.Thumb)
			}
		}
	}

	// 保存富文本
	if len(in.Content) != 0 {
		rich.Edit(ctx, &info.Content, &in.Content)
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
	if err != nil {
		return utility.Err.Sys(err)
	}

	return
}

// List 读取文章列表
func List(ctx context.Context, query *model.ArticleQuery) (list []entity.Article, total uint, err error) {
	if query == nil {
		query = &model.ArticleQuery{}
	}
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
		db = db.Where("grp_id", query.GrpId).Where("id")
	}
	// 是否查询只发布的文章
	if query.Onshow {
		db = db.Where("onshow", true)
	}
	// 搜索文本
	if len(query.Search) != 0 {
		db = db.Where(db.Builder().Where("title like ?", "%"+query.Search+"%").
			WhereOr("tags like ?", "%"+query.Search+"%").
			WhereOr("description like ?", "%"+query.Search+"%"))
	}
	// 是否查询删除掉的文章
	if query.IsDel {
		db = db.Unscoped().Where("deleted_at is not null")
	}
	db = db.Order("created_at desc, id desc").Page(query.Page, query.Size)

	data, totalInt, err := db.AllAndCount(true)
	if err != nil {
		err = utility.Err.Sys(err)
		return
	}
	list = []entity.Article{}
	_ = data.Structs(&list)

	total = uint(totalInt)
	return
}

// Show 读取文章详情
func Show(ctx context.Context, id model.Id) (info *entity.Article, err error) {
	info = &entity.Article{}
	err = dao.Article.Ctx(ctx).Where("id", id).Scan(&info)
	if err != nil {
		err = utility.Err.TableNotData(err)
	}
	return
}

// Del 删除文章
func Del(ctx context.Context, id model.Id, isReal bool) (err error) {
	if isReal {
		info, err := Show(ctx, id)
		if err != nil {
			return err
		}
		if info != nil {
			// 删除文件资源
			if len(info.Thumb) > 0 {
				_ = file.Del(ctx, info.Thumb)
			}
			if len(info.Content) > 0 {
				_ = rich.Del(ctx, &info.Content)
			}
			_, err = dao.Article.Ctx(ctx).Where("id", id).Unscoped().Delete()
		}
	} else {
		_, err = dao.Article.Ctx(ctx).Where("id", id).Delete()
	}
	return
}

// ReCre 重新创建已经删除的文章
func ReCre(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Article.Ctx(ctx).Where("id", id).Unscoped().Update("deleted_at=null")
	return
}

// Hist 为文章增加一个点击数
func Hist(ctx context.Context, id model.Id) (err error) {
	var (
		redis = g.Redis()
		ip    = g.RequestFromCtx(ctx).GetClientIp()
		key   = "hist" + ip + gconv.String(id)
	)
	// 判断缓存中是否有ip
	ok, _ := redis.Get(ctx, key)
	if !ok.Bool() {
		// +1个hist，设置缓存，当前未接入全局缓存，所以时间控制是没有用的
		_, _ = dao.Article.Ctx(ctx).Where("id", id).Increment("hist", 1)
		_, _ = redis.Set(ctx, key, true)
		_, _ = redis.Expire(ctx, key, 60)
	}
	return
}

// UpdLastedAt 更新最后阅读时间
func UpdLastedAt(ctx context.Context, id model.Id) (err error) {
	_, err = dao.Article.Ctx(ctx).Where("id", id).Data(do.Article{
		LastedAt: gtime.Now(),
	}).Update()
	return
}

// IsExist 根据id判断一个文章是否存在
func IsExist(ctx context.Context, id model.Id) bool {
	num, _ := dao.Article.Ctx(ctx).Where("id", id).Count()
	return num == 1
}
