package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jamal/api/controller"
	"jamal/api/models/domain"
	"jamal/api/repository"
	"jamal/api/service"
	"log"
)

func InitDB() *gorm.DB {
	dsn := "root:korie123@tcp(127.0.0.1:3306)/api_product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(errors.New("Failed Connected into data base"))
	}

	err = db.AutoMigrate(&domain.Product{}, &domain.User{})
	if err != nil {
		panic(errors.New("Failed Migrated"))
	}

	return db
}

func main() {
	db := InitDB()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, db)
	productController := controller.NewProductController(productService)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(db, authRepository)
	authController := controller.NewAuthController(authService)

	router := gin.Default()

	router.POST("api/auth/register", authController.Register)

	router.POST("api/products", productController.Create)
	router.GET("api/products", productController.FindAll)
	router.GET("api/products/:id", productController.FindById)
	router.PUT("api/products/:id", productController.Update)
	router.DELETE("api/products/:id", productController.Delete)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
