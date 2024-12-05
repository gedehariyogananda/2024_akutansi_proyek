package Services

import (
	"2024_akutansi_project/Consts"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Repositories"
	"context"
	"errors"
	"firebase.google.com/go/messaging"
	"fmt"
	"time"
)

type (
	IWorkerService interface {
		SendPushNotification(ctx context.Context) (err error)
	}

	WorkerService struct {
		deviceTokenRepo  *Repositories.DeviceTokenRepository
		messagingClient  *messaging.Client
		notificationRepo *Repositories.NotificationRepository
	}
)

func WorkerServiceProvider(deviceTokenRepo *Repositories.DeviceTokenRepository, messagingClient *messaging.Client, notificationRepo *Repositories.NotificationRepository) *WorkerService {
	return &WorkerService{
		deviceTokenRepo:  deviceTokenRepo,
		messagingClient:  messagingClient,
		notificationRepo: notificationRepo,
	}
}

func (s *WorkerService) SendPushNotification(ctx context.Context) error {
	fmt.Println(fmt.Sprintf("Running worker push notification on = %v", time.Now()))

	pendingNotification := s.notificationRepo.FindNotification(ctx, Models.Notification{
		Status: Consts.Pending,
	})

	if pendingNotification != nil {
		for _, f := range *pendingNotification {
			if err := s.sendPushNotification(context.WithoutCancel(ctx), f); err != nil {
				fmt.Println(fmt.Sprintf("error sending push notification to %s, got err = %v", f.UserID, err))
				continue
			}
		}
	}

	return nil
}

func (s *WorkerService) sendPushNotification(ctx context.Context, data Models.Notification) error {
	deviceToken, err := s.deviceTokenRepo.Get(ctx, data.UserID)
	if err != nil {
		return err
	}

	message := &messaging.Message{
		Token:        deviceToken.DeviceToken,
		Notification: Consts.PushNotificationScheme(data.Scheme).DefinePushNotificationMessages(),
	}

	var timeNow = time.Now()
	response, err := s.messagingClient.Send(ctx, message)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully sent message: %s\n", response)

	data.Status = Consts.Sent
	data.SentAt = &timeNow
	_ = s.notificationRepo.UpdateNotification(ctx, data)

	return nil
}

func (s *WorkerService) QueuePushNotification(ctx context.Context) error {
	return fmt.Errorf("error : %v", errors.New("no data to be queued"))
}
