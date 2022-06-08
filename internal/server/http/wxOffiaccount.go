package http

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils/box"

	"github.com/gin-gonic/gin"
)

type WxOffiaccount struct {
	*Server
}

//FindAll 查询所有公众号
func (t *WxOffiaccount) FindAll(c *box.Context) {
	r, err := t.service.WxOffiaccount.FindAll()
	c.JSON(r, err)
}

//Find 根据username查询公众号
func (t *WxOffiaccount) Find(c *box.Context) {
	var param model.WxOffiaccount
	c.ShouldBindJSON(&param)
	fmt.Printf("wxOffiaccount->Find:%+v\n", param)
	r, err := t.service.WxOffiaccount.Find(param.Username)
	c.JSON(r, err)
}

func (t *WxOffiaccount) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("wxOffiaccount")
	{
		r.POST("findAll", box.Handler(t.FindAll))
		r.POST("find", box.Handler(t.Find))
	}
}
