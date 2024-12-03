package Config

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
)

func InitFirebase() *firebase.App {
	opt := option.WithCredentialsFile(os.Getenv("SERVICE_ACCOUNT_PATH"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v", err)
	}

	fmt.Println("Connected to Firebase App!")

	return app
}

func InitMessagingClient(app *firebase.App) *messaging.Client {
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error initializing Messaging Client: %v", err)
	}

	fmt.Println("Messaging Client Initialized!")

	return client
}
