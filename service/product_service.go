package service

import (
	"github.com/gin-gonic/gin"
	"jamal/api/models/web"
)

type ProductService interface {
	Create(ctx *gin.Context, createReq web.ProductCreate) web.WebResponse
	Delete(ctx *gin.Context, productId int) web.WebResponse
	Update(ctx *gin.Context, updateReq web.ProductUpdate, productId int) web.WebResponse
	FindById(ctx *gin.Context, productId int) web.WebResponse
	FindAll(ctx *gin.Context) web.WebResponse
}
