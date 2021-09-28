package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/seigaalghi/majoo-pos/controllers"
	"github.com/seigaalghi/majoo-pos/middlewares"
	"github.com/seigaalghi/majoo-pos/models"
)

func main() {
	// Server setup
	server := gin.Default()

	// DB Setup
	db := models.Database()
	db.AutoMigrate(&models.Merchant{})
	db.AutoMigrate(&models.Product{})
	server.Use(func(context *gin.Context) {
		context.Set("db", db)
	})

	// V1 Route
	serverV1 := server.Group("/api/v1")
	serverV1.POST("/login", controllers.Login)

	// V1 Merchant Route
	merchantV1 := serverV1.Group("/merchant")
	merchantV1.Use(middlewares.Protect())
	merchantV1.POST("/", controllers.CreateMerchant)
	merchantV1.PUT("/:id", controllers.UpdateMerchant)
	merchantV1.DELETE("/:id", controllers.DeleteMerchant)
	merchantV1.GET("/", controllers.GetMerchants)
	merchantV1.GET("/:id", controllers.GetMerchant)

	// V1 Product Route
	productV1 := serverV1.Group("/product")
	productV1.Use(middlewares.AuthorizeJWT())
	productV1.POST("/", controllers.CreateProduct)
	productV1.GET("/:id", controllers.GetProduct)
	productV1.GET("/", controllers.GetProducts)
	productV1.PUT("/:id", controllers.UpdateProduct)
	productV1.DELETE("/:id", controllers.DeleteProduct)
	server.Run(":5000")
}

// gin --appPort 5000 --all -i run main.go
