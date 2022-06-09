package http

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils/box"

	"github.com/gin-gonic/gin"
)

type WxOffiaccountServer struct {
	*Server
}

func NewWxOffiaccountServer(s *Server) *WxOffiaccountServer {
	return &WxOffiaccountServer{s}
}

func (t *WxOffiaccountServer) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("wxOffiaccount")
	{
		r.POST("findAll", box.Handler(t.FindAll))
		r.POST("find", box.Handler(t.Find))
	}
}

//FindAll 查询所有公众号
func (t *WxOffiaccountServer) FindAll(c *box.Context) {
	var param model.WxOffiaccount
	c.ShouldBindJSON(&param)
	fmt.Printf("FindAll: %+v\n", param)

	r, err := t.service.WxOffiaccount.FindAll()
	c.JSON(r, err)
}

//Find 根据username查询公众号
func (t *WxOffiaccountServer) Find(c *box.Context) {
	var param model.WxOffiaccount
	c.ShouldBindJSON(&param)
	fmt.Printf("wxOffiaccount->Find:%+v\n", param)
	r, err := t.service.WxOffiaccount.Find(param.Username)
	c.JSON(r, err)
}

