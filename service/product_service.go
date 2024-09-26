package service

import (
	"github.com/gin-gonic/gin"
	"jamal/api/models/web"
)

type ProductService interface {
	Create(ctx *gin.Context, createReq web.ProductCreate) web.ProductResponse
	Delete(ctx *gin.Context, productId int) web.ProductResponse
	Update(ctx *gin.Context, updateReq web.ProductUpdate, productId int) web.ProductResponse
	FindById(ctx *gin.Context, productId int) web.ProductResponse
	FindAll(ctx *gin.Context) web.ProductResponse
}
