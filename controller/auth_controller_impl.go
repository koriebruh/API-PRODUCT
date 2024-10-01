package controller

import (
	"github.com/gin-gonic/gin"
	"jamal/api/models/web"
	"jamal/api/service"
	"net/http"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{AuthService: authService}
}

func (controller AuthControllerImpl) Register(c *gin.Context) {
	// DECODE DARI JSON KE STRUCT
	var registerUser web.AuthRequestRegister
	if err := c.ShouldBindJSON(&registerUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": "Invalid input data",
		})
		return
	}

	registerResponse := controller.AuthService.Register(c, registerUser)

	c.JSON(registerResponse.Code, registerResponse)
}

func (controller AuthControllerImpl) Login(c *gin.Context) {
	// DECODE DARI JSON KE STRUCT
	var LoginUser web.AuthRequestLogin
	if err := c.ShouldBindJSON(&LoginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": "Invalid input data",
		})
		return
	}

	logins := controller.AuthService.Login(c, LoginUser)

	c.JSON(logins.Code, logins)
}

func (controller AuthControllerImpl) Logout(c *gin.Context) {

	logins := controller.AuthService.Logout(c)

	c.JSON(logins.Code, logins)
}
