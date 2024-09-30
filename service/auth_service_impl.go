package service

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"jamal/api/helper"
	"jamal/api/models/domain"
	"jamal/api/models/web"
	"jamal/api/repository"
	"net/http"
)

type AuthServiceImpl struct {
	DB             *gorm.DB
	AuthRepository repository.AuthRepository
}

func NewAuthService(db *gorm.DB, authRepository repository.AuthRepository) AuthService {
	return &AuthServiceImpl{
		DB:             db,
		AuthRepository: authRepository,
	}
}

func (service AuthServiceImpl) Register(ctx *gin.Context, register web.AuthRequestRegister) web.WebResponse {
	var response web.WebResponse
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		//<-- hash password
		hashPsw, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
		if err != nil { // <-- response ketika gagal hash
			response = web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "InternalServerError",
				Data: map[string]interface{}{
					"Error": "Error Hash Password",
				},
			}
			return nil
		}
		register.Password = string(hashPsw)

		//<-- add format request register
		registerData := domain.User{
			UserName: register.UserName,
			Password: register.Password,
			Name:     register.Name,
		}

		err = service.AuthRepository.Register(tx, registerData)
		if err != nil { //<-- response ketika gagal create
			response = web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   err.Error(),
			}
			return nil
		}

		//<-- response sukses
		response = web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "Register Success",
		}
		return nil
	})
	helper.HandleErrorResponse(&response, err) // <-- ketika error transaction

	return response
}

func (service AuthServiceImpl) Login(ctx *gin.Context, login web.AuthRequestLogin) web.WebResponse {
	var response web.WebResponse
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		//<-- hash password
		hashPsw, err := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.DefaultCost)
		if err != nil { // <-- response ketika gagal hash
			response = web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "InternalServerError",
				Data: map[string]interface{}{
					"Error": "Error Hash Password",
				},
			}
			return nil
		}
		login.Password = string(hashPsw)

		// <-- add adta format
		requestLogin := domain.User{
			Model:    gorm.Model{},
			UserName: login.UserName,
			Password: login.Password,
			Name:     "",
		}

		err = service.AuthRepository.Login(tx, requestLogin)
		if err != nil { //<-- response ketika gagal create
			response = web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   err.Error(),
			}
			return nil
		}

		//<-- response sukses
		response = web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "Register Success",
		}

		//<-- Generate JWT TOKEN

		return nil

	})

	helper.HandleErrorResponse(&response, err) // <-- ketika error transaction

	return response
}

func (service AuthServiceImpl) Logout(ctx *gin.Context) web.WebResponse {
	//TODO implement me
	panic("implement me")
}
