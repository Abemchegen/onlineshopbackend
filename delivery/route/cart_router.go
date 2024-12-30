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

func NewCartRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	CartRepository := repo.NewCartRepository(DB, config.CartCollection)
	CartUsecase := usecase.NewCartUseCase(CartRepository)
	CartController := controller.NewCartController(CartUsecase)

	CartRouter := route.Group("/Cart")
	CartRouter.Use(infrastructure.AuthMiddleware())
	{
		CartRouter.POST("/", CartController.CreateCart)
		CartRouter.GET("/", CartController.GetCartByUserID)
		CartRouter.PUT("/", CartController.UpdateCart)
		CartRouter.DELETE("/", CartController.DeleteCart)

	}

}
