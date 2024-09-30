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
	var registerProduct web.AuthRequestRegister
	if err := c.ShouldBindJSON(&registerProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": "Invalid input data",
		})
		return
	}

	registerResponse := controller.AuthService.Register(c, registerProduct)

	c.JSON(registerResponse.Code, registerResponse)
}

func (controller AuthControllerImpl) Login(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller AuthControllerImpl) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
