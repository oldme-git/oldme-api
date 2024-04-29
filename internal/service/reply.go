// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/oldme-git/oldme-api/internal/model"
	"github.com/oldme-git/oldme-api/internal/model/entity"
)

type (
	IReply interface {
		// Cre 新增回复
		Cre(ctx context.Context, in *entity.Reply) (lastId model.Id, err error)
		// Upd 更新文章回复
		Upd(ctx context.Context, id model.Id, in *model.ReplyBody) (err error)
		// Del 删除文章回复
		Del(ctx context.Context, id model.Id) (err error)
		// List 读取文章回复列表
		List(ctx context.Context, query *model.ReplyQuery) (list []entity.Reply, total uint, err error)
		// Show 读取文章回复详情，携带父级信息
		Show(ctx context.Context, id model.Id) (info *model.ReplyShow, err error)
		// Details 读取文章回复详情
		Details(ctx context.Context, id model.Id) (info *entity.Reply, err error)
		// Check 审核
		Check(ctx context.Context, id model.Id, result bool, reasonSlice ...string) error
		// GetRid 根据父id获取根id，如果父回复已经是根回复了，则rid=0，否则rid=pid
		GetRid(ctx context.Context, pid model.Id) (model.Id, error)
		// GetAid 根据父id获取aid，如果父回复已经是根回复了，则aid返回0
		GetAid(ctx context.Context, pid model.Id) (model.Id, error)
		// ListForAid 根据Aid读取回复列表
		ListForAid(ctx context.Context, aid model.Id) (uint, []model.ReplyFloorApp, error)
		// GetReplyFloor 根据rid读取本层楼的所有回复信息，id
		GetReplyFloor(list []entity.Reply, rid model.Id) (reply []model.ReplyFloorApp)
	}
)

var (
	localReply IReply
)

func Reply() IReply {
	if localReply == nil {
		panic("implement not found for interface IReply, forgot register?")
	}
	return localReply
}

func RegisterReply(i IReply) {
	localReply = i
}
