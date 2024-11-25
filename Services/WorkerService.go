package Services

import (
	"context"
	"fmt"
	"time"
)

type (
	IWorkerService interface {
		SendPushNotification(ctx context.Context) (err error)
	}

	WorkerService struct {
	}
)

func WorkerServiceProvider() *WorkerService {
	return &WorkerService{}
}

func (s *WorkerService) SendPushNotification(ctx context.Context) (err error) {
	fmt.Println(fmt.Sprintf("Running worker push notification on = %v", time.Now()))

	return nil
}
