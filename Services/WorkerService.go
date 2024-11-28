package Services

import (
	"2024_akutansi_project/Repositories"
	"context"
	"firebase.google.com/go/messaging"
	"fmt"
	"time"
)

type (
	IWorkerService interface {
		SendPushNotification(ctx context.Context) (err error)
	}

	WorkerService struct {
		deviceTokenRepo *Repositories.DeviceTokenRepository
		messagingClient *messaging.Client
	}
)

func WorkerServiceProvider(deviceTokenRepo *Repositories.DeviceTokenRepository, messagingClient *messaging.Client) *WorkerService {
	return &WorkerService{
		deviceTokenRepo: deviceTokenRepo,
		messagingClient: messagingClient,
	}
}

func (s *WorkerService) SendPushNotification(ctx context.Context) (err error) {
	fmt.Println(fmt.Sprintf("Running worker push notification on = %v", time.Now()))

	//deviceToken, err := s.deviceTokenRepo.Get(ctx, userID)
	//if err != nil {
	//	return err
	//}
	//
	//message := &messaging.Message{
	//	Token: deviceToken.DeviceToken,
	//	Notification: &messaging.Notification{
	//		Title: "Push Notification",
	//		Body:  fmt.Sprintf(""),
	//	},
	//}
	//
	//response, err := s.messagingClient.Send(ctx, message)
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Printf("Successfully sent message: %s\n", response)

	return nil
}
