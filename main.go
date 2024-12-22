package main

import (
	"2024_akutansi_project/Config"
	"2024_akutansi_project/Consts"
	"2024_akutansi_project/Dependencies"
	"2024_akutansi_project/Middleware"
	"2024_akutansi_project/Routes"
	"2024_akutansi_project/Utils"
	"2024_akutansi_project/Worker"
	_ "2024_akutansi_project/docs"
	"os"
	"strconv"

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
	Utils.InitValidator()

	deps := Dependencies.InitDependencies(
		Dependencies.WithDB(),
		Dependencies.WithRedis(),
		// Dependencies.WithMongo(),
		// Dependencies.WithFirebase(),
		// Dependencies.WithMessagingClient(),
	)

	setup := gin.Default()
	setup.RedirectTrailingSlash = true
	setup.Use(Middleware.SetupCORS())
	setup.Use(Middleware.ExecutionTimeMiddleware())

	setup.MaxMultipartMemory = int64(Consts.MaxMultipartMemory)

	setup.Static(Consts.StaticFileRoute, Consts.StaticFileDir)

	Routes.Init(setup, deps)

	if func() bool {
		scheduler := os.Getenv("USE_SCHEDULER")
		b, _ := strconv.ParseBool(scheduler)
		return b
	}() {
		Worker.InitScheduler(deps)
	}

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
