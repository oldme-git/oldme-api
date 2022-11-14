package admin

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"oldme-api/internal/service"
)

type sAdmin struct {
}

func init() {
	service.RegisterAdmin(&sAdmin{})
}

func New() *sAdmin {
	return &sAdmin{}
}

func (s sAdmin) Create(r *ghttp.Request) {

}

func (s sAdmin) Update(r *ghttp.Request) {

}

func (s sAdmin) Read(r *ghttp.Request) {

}

func (s sAdmin) Del(r *ghttp.Request) {

}
