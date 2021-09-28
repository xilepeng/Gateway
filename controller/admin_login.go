package controller

import (
	"github.com/e421083458/Gateway/dto"
	"github.com/e421083458/Gateway/middleware"
	"github.com/gin-gonic/gin"
)

type AdminLoginController struct {

}

func AdminLoginRegister(group *gin.RouterGroup)  {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}

func (adminlogin *AdminLoginController)AdminLogin(c *gin.Context)  {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c,1001,err)
		return
	}
	middleware.ResponseSuccess(c,"")
}