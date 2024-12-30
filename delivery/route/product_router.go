package routers

import (
	"onlineshopbackend/config"
	"onlineshopbackend/delivery/controller"
	"onlineshopbackend/infrastructure"
	"onlineshopbackend/repo"
	"onlineshopbackend/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProductRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	productRepository := repo.NewProductRepository(DB, config.ProductCollection)
	productUsecase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUsecase)

	productRouter := route.Group("/product")
	productRouter.Use(infrastructure.AuthMiddleware())
	{
		productRouter.POST("/", productController.CreateProduct)
		productRouter.GET("/all", productController.GetAllProduct)
		productRouter.GET("/", productController.GetProductByID)
		productRouter.PUT("/", productController.UpdateProduct)
		productRouter.DELETE("/", productController.DeleteProduct)

	}

}
