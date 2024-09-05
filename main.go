package main

import (
	"2024_akutansi_project/Config"
	"2024_akutansi_project/Middleware"
	"2024_akutansi_project/Routes"
	"2024_akutansi_project/Utils"

	_ "2024_akutansi_project/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	Utils.LoadEnv()

	Config.Connect()
	db := Config.DB
	if db == nil {
		panic("Failed to connect to database!")
	}

	setup := gin.Default()
	setup.MaxMultipartMemory = 8 << 20 // 8 MB

	setup.Static("/uploads", "./public/uploads")

	// Setup CORS
	setup.Use(Middleware.SetupCORS())

	// Setup routes
	Routes.Init(setup, db)

	// Setup Swagger
	setup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run server
	server := Config.GetServerAddress()
	if err := setup.Run(server); err != nil {
		panic("Failed to run server!")
	}
}
