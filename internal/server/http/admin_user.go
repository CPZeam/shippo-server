package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
)

type AdminUserServer struct {
	*Server
}

func NewAdminUserServer(s *Server) *AdminUserServer {
	return &AdminUserServer{s}
}

func (t *AdminUserServer) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("admin/user")
	{
		r.POST("create", box.Handler(t.UserCreateEmail))
	}
}

func (t *AdminUserServer) UserCreateEmail(c *box.Context) {
	var param = new(struct {
		Email string `json:"email"`
	})
	c.ShouldBindJSON(&param)
	fmt.Printf("userCreateEmail: %+v\n", param)

	_, err := t.service.AdminUser.AdminUserCreateEmail(param.Email)
	c.JSON(nil, err)
}
