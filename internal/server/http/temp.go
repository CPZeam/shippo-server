package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type TempServer struct {
	*Server
}

func NewTempServer(s *Server) *TempServer {
	return &TempServer{s}
}

func (t *TempServer) InitRouter(Router *gin.RouterGroup) {
	var h = box.NewBoxHandler(&t)

	r := Router.Group("temp")
	{
		r.POST("temp_trade_20220108/find", h.H(t.Temp_trade_20220108_find, box.AccessAll))
		r.POST("temp_trade_20220108/add", h.H(t.Temp_trade_20220108_add, box.AccessAll))
		r.POST("temp_trade_20220108/findNoExist", h.H(t.Temp_trade_20220108_findNoExist, box.AccessAll))
	}
}

func (t *TempServer) Temp_trade_20220108_find(c *box.Context) {

	var param = new(struct {
		Qq string `json:"qq"`
		Id string `json:"id"`
	})
	c.ShouldBindJSON(&param)
	fmt.Printf("temp_trade_20220108_find: %+v\n", param)

	// 如果参数中含有QQ，那么就按照QQ查找，否则按照订单号。
	if param.Qq != "" {
		data, err := t.service.Temp.Temp_trade_20220108_findByUserQQ(c, param.Qq)
		c.JSON(data, err)
	} else {
		data, err := t.service.Temp.Temp_trade_20220108_findByTradeId(c, param.Id)
		c.JSON(data, err)
	}
}

func (t *TempServer) Temp_trade_20220108_add(c *box.Context) {
	var param model.Temp_trade_20220108_TradeAddParam
	c.ShouldBindJSON(&param)
	fmt.Printf("temp_trade_20220108_add: %+v\n", param)

	data, err := t.service.Temp.Temp_trade_20220108_add(c, param)
	c.JSON(data, err)
}

func (t *TempServer) Temp_trade_20220108_findNoExist(c *box.Context) {
	var param = new(struct {
		List []string `json:"list"`
	})
	c.ShouldBindJSON(&param)

	data, err := t.service.Temp.Temp_trade_20220108_findNoExist(c, param.List)
	c.JSON(data, err)
}
