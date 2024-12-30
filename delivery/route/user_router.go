package routers

import (
	"onlineshopbackend/config"
	"onlineshopbackend/delivery/controller"

	"onlineshopbackend/infrastructure"
	"onlineshopbackend/repo"
	"onlineshopbackend/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	timeout := time.Duration(config.ContextTimeout) * time.Second

	repository := repo.NewUserRepository(DB, config.UserCollection)

	tokenGen := infrastructure.NewTokenGenerator()
	passwordSvc := infrastructure.NewPasswordService()

	usecase := usecase.NewUserUseCase(repository, timeout, tokenGen, passwordSvc)

	userController := controller.NewUserController(usecase)

	user := route.Group("/user")
	{

		//user/register
		user.POST("/register", userController.CreateAccount)
		user.POST("/login", userController.Login)
		user.Use(infrastructure.AuthMiddleware())
		user.GET("/:ID", userController.GetByID)
		user.PUT("/updateProfile", userController.UpdateProfile)

		user.GET("/get-all", userController.GetAllUser)

	}

}
