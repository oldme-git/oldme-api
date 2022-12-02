package admin

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/do"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/service"
)

type sAdmin struct {
}

func init() {
	service.RegisterAdmin(&sAdmin{})
}

// Create 创建管理员
func (s *sAdmin) Create(ctx context.Context, in *model.AdminInput) (err error) {
	var (
		salt     = genSalt()
		password = encryptPass(in.Password, salt)
	)
	_, err = dao.Admin.Ctx(ctx).Data(do.Admin{
		Username:  in.Username,
		Password:  password,
		Nickname:  in.Nickname,
		Avatar:    in.Avatar,
		Salt:      salt,
		LastLogin: gtime.Now(),
		Register:  gtime.Now(),
	}).Insert()
	return
}

// ValidPass 校验密码
func (s *sAdmin) ValidPass(pass string, admin *entity.Admin) bool {
	return admin.Password == encryptPass(pass, admin.Salt)
}

// GenSalt 生成32位盐值盐值
func genSalt() string {
	return grand.S(32, false)
}

// EncryptPass 加密密码
func encryptPass(pass string, salt string) (encrypt string) {
	encrypt, _ = gmd5.EncryptString(pass + salt)
	return
}
