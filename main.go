package main

import (
	"os"

	"2024_akutansi_project/Config"
	"2024_akutansi_project/Routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// config db
	Config.Connect()
	db := Config.DB

	if db == nil {
		panic("Failed to connect to database!")
	}

	// setup gin config 
	setup := gin.Default()

	// max memory
	setup.MaxMultipartMemory = 8 << 20

	// akses folder uploads
	setup.Static("/uploads", "./public/uploads")

	// setup cors origin config
	setup.Use(cors.New(cors.Config{
		AllowHeaders: []string{"Origin,Content-Type,Accept,User-Agent,Content-Length,Authorization"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"},
	}))

	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env file!")
	}

	portServer := os.Getenv("SERVER_PORT")
	if portServer == "" {
		portServer = "8888"
	}

	app := os.Getenv("APP_ENV")

	var server string
	if app == "local" {
		server = "127.0.0.1:" + portServer
	} else {
		server = portServer
	}

	// init route
	Routes.Init(setup, db)

	if err := setup.Run(server); err != nil {
		panic("Failed to run server!")
	}

}
