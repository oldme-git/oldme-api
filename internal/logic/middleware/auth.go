package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"oldme-api/internal/consts"
	"oldme-api/internal/utils"
)

func (s sMiddleware) Auth(r *ghttp.Request) {
	url := r.GetRouterMap()
	utils.Log(url)
	var (
		jwtKey      = []byte(consts.JwtKey)
		tokenString = r.Header.Get("Authorization")
	)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
	}

	r.Middleware.Next()
}
