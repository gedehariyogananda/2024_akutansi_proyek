package main

import (
	"2024_akutansi_project/Config"
	"2024_akutansi_project/Dependencies"
	"2024_akutansi_project/Middleware"
	"2024_akutansi_project/Routes"
	"2024_akutansi_project/Utils"
	_ "2024_akutansi_project/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 2024akutansi
// @version 1.0
// @description This is a documentation API SPEC for the 2024akutansi project.
// @BasePath /api/v1
// @schemes http https
// @host localhost:8899

func main() {
	Utils.LoadEnv()

	deps := Dependencies.InitDependencies(
		Dependencies.WithDB(),
		Dependencies.WithRedis(),
		// Dependencies.WithMongo(),
	)

	setup := gin.Default()
	setup.RemoveExtraSlash = true
	setup.Use(Middleware.ExecutionTimeMiddleware())

	setup.MaxMultipartMemory = int64(Config.MaxMultipartMemory)

	setup.Static(Config.StaticFileRoute, Config.StaticFileDir)

	setup.Use(Middleware.SetupCORS())

	Routes.Init(setup, deps)

	setup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setup.GET("/checker", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "checked health",
		})
	})

	server := Config.GetServerAddress()
	if err := setup.Run(server); err != nil {
		panic("Failed to run server!")
	}
}
