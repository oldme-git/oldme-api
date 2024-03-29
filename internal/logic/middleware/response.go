// 统一响应处理

package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/oldme-git/oldme-api/internal/packed"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"    dc:"业务码"`
	Message string      `json:"message" dc:"业务码说明"`
	Data    interface{} `json:"data"    dc:"返回的数据"`
}

func (s *sMiddleware) Response(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	// 先过滤掉服务器内部错误
	if r.Response.Status >= http.StatusInternalServerError {
		// 清除掉缓存区，防止服务器信息泄露到客户端
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器打盹了，请稍后再来找他！")
	}

	var (
		res  = r.GetHandlerResponse()
		err  = r.GetError()
		code = gerror.Code(err)
		msg  string
	)

	if err != nil {
		msg = err.Error()
	} else {
		code = gcode.CodeOK
		msg = packed.Err.GetMsg(code.Code())
	}

	r.Response.WriteJson(Response{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
