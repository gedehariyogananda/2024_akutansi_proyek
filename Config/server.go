package Config

import "os"

func GetServerAddress() string {
	portServer := os.Getenv("SERVER_PORT")
	if portServer == "" {
		portServer = "8888"
	}

	app := os.Getenv("APP_ENV")

	if app == "local" {
		return "127.0.0.1:" + portServer
	}
	return ":" + portServer
}
