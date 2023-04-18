package account

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v4"
	"oldme-api/internal/consts"
	"oldme-api/internal/dao"
	"oldme-api/internal/model"
	"oldme-api/internal/model/entity"
	"oldme-api/internal/packed"
	"oldme-api/internal/service"
	"time"
)

type sAccount struct {
}

func init() {
	service.RegisterAccount(&sAccount{})
}

var jwtKey = consts.JwtKey

type AdminClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

// Login 登录
func (s *sAccount) Login(ctx context.Context, in *model.Login) (tokenString string, err error) {
	admin := dao.Admin.GetAdmin(in.Username)
	// 校验账号和密码
	if admin.Id != 0 && service.Admin().ValidPass(in.Password, admin) {
		adminClaims := &AdminClaims{
			Id:       admin.Id,
			Username: admin.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, adminClaims)
		tokenString, _ = token.SignedString(jwtKey)
		err = nil
	} else {
		err = packed.Err.Skip(20100)
	}
	return
}

// Logout 登出
func (s *sAccount) Logout(ctx context.Context) (err error) {
	// TODO 暂时依靠客户端，后续可以考虑改为Redis黑名单处理
	return nil
}

// Info 获取当前已经登录的管理员信息
func (s *sAccount) Info(ctx context.Context) (admin *entity.Admin, err error) {
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := tokenClaims.Claims.(*AdminClaims); ok && tokenClaims.Valid {
		admin = dao.Admin.GetAdmin(claims.Username)
	} else {
		err = packed.Err.Skip(10100)
	}
	return
}
