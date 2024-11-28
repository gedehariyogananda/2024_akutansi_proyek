package Services

import (
	"2024_akutansi_project/Repositories"
	"context"
	"fmt"
	"time"
)

type (
	IWorkerService interface {
		SendPushNotification(ctx context.Context) (err error)
	}

	WorkerService struct {
		deviceTokenRepo *Repositories.DeviceTokenRepository
	}
)

func WorkerServiceProvider(deviceTokenRepo *Repositories.DeviceTokenRepository) *WorkerService {
	return &WorkerService{
		deviceTokenRepo: deviceTokenRepo,
	}
}

func (s *WorkerService) SendPushNotification(ctx context.Context) (err error) {
	fmt.Println(fmt.Sprintf("Running worker push notification on = %v", time.Now()))

	return nil
}
