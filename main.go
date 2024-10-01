package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jamal/api/config"
	"jamal/api/controller"
	"jamal/api/models/domain"
	"jamal/api/repository"
	"jamal/api/service"
	"log"
	"net/http"
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

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Ambil token dari cookie
		cookie, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, please login"})
			ctx.Abort() // Menghentikan request jika tidak ada token
			return
		}

		// Log untuk debugging
		log.Printf("Token received: %s", cookie)

		// Verifikasi token
		token, err := jwt.ParseWithClaims(cookie, &config.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_KEY), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, invalid token"})
			// Log kesalahan untuk debugging
			log.Printf("Token error: %v", err)
			ctx.Abort() // Hentikan request jika token tidak valid
			return
		}

		// Lanjutkan ke handler berikutnya
		ctx.Next()
	}
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

	router.POST("api/auth/login", authController.Login)
	router.POST("api/auth/logout", authController.Logout)
	router.POST("api/auth/register", authController.Register)

	// <-- only can access if login
	authorized := router.Group("/")
	authorized.Use(AuthMiddleware())
	{
		authorized.POST("api/products", productController.Create)
		authorized.GET("api/products", productController.FindAll)
		authorized.GET("api/products/:id", productController.FindById)
		authorized.PUT("api/products/:id", productController.Update)
		authorized.DELETE("api/products/:id", productController.Delete)

	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
