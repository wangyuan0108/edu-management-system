package controller

import (
	"edu-management-system/schema"
	api "edu-management-system/service/api/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user = new(schema.UserBasic)
	account := c.PostForm("account")
	password := c.PostForm("password")

	api.Login(account, password, user)

	c.JSON(http.StatusOK, schema.Status{
		Code:    http.StatusOK,
		Message: "登陆成功",
		Body:    user,
	})
}
