package main

import (
	"fmt"
	"onlineshopbackend/config"
	routers "onlineshopbackend/delivery/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	env, err := config.LoadEnv()

	if err != nil {
		fmt.Print("Error in env.load", err)
	}

	DB, client, err := config.ConnectDB(env.DatabaseUrl, env.DbName)

	if err != nil {
		fmt.Print("Error in connectDB")
	}

	defer config.CloseDB(client)

	routers.Router(server.Group("api/"), env, DB)
	server.Run(fmt.Sprintf(":%d", env.Port))

}
