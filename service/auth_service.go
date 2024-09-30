package service

import (
	"github.com/gin-gonic/gin"
	"jamal/api/models/web"
)

type AuthService interface {
	Register(ctx *gin.Context, register web.AuthRequestRegister) web.WebResponse
	Login(ctx *gin.Context, login web.AuthRequestLogin) web.WebResponse
	Logout(ctx *gin.Context) web.WebResponse
}
