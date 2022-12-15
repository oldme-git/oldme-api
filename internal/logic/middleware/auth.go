package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"oldme-api/internal/consts"
)

func (s *sMiddleware) Auth(r *ghttp.Request) {
	var (
		jwtKey      = consts.JwtKey
		tokenString = r.Header.Get("Authorization")
	)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
	}

	r.Middleware.Next()
}
