package Services

import (
	"2024_akutansi_project/Consts"
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Repositories"
	"context"
	"encoding/json"
	"firebase.google.com/go/messaging"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

type (
	IWorkerService interface {
		SendPushNotification(ctx context.Context) (err error)
	}

	WorkerService struct {
		deviceTokenRepo     *Repositories.DeviceTokenRepository
		messagingClient     *messaging.Client
		notificationRepo    *Repositories.NotificationRepository
		materialStockRepo   *Repositories.MaterialStockRepository
		sellableProductRepo *Repositories.SellableProductRepository
		sellableStockRepo   *Repositories.SellableStockRepository
	}
)

func WorkerServiceProvider(deviceTokenRepo *Repositories.DeviceTokenRepository,
	messagingClient *messaging.Client,
	notificationRepo *Repositories.NotificationRepository,
	materialStockRepo *Repositories.MaterialStockRepository,
	sellableProductRepo *Repositories.SellableProductRepository,
	sellableStockRepo *Repositories.SellableStockRepository,
) *WorkerService {
	return &WorkerService{
		deviceTokenRepo:     deviceTokenRepo,
		messagingClient:     messagingClient,
		notificationRepo:    notificationRepo,
		materialStockRepo:   materialStockRepo,
		sellableProductRepo: sellableProductRepo,
		sellableStockRepo:   sellableStockRepo,
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

	wg, ctxg := errgroup.WithContext(ctx)
	wg.SetLimit(3)

	// Material Stock
	func(ctxWG context.Context, eg *errgroup.Group) {
		wg.Go(func() error {
			materialStock := s.materialStockRepo.FetchMaterialStockToPushNotification(ctx)
			for _, ms := range materialStock {
				// todo :: adjust business logic
				if err := s.notificationRepo.StoreNotification(ctx, Models.Notification{
					Scheme: Consts.Pending,
					UserID: "",
					Status: "",
					AdditionalData: func() json.RawMessage {
						byteData, _ := json.Marshal(ms)
						return byteData
					}(),
					QueuedAt:    time.Time{},
					ScheduledAt: time.Time{},
					SentAt:      nil,
				}); err != nil {
					return err
				}
			}
			return nil
		})
	}(ctxg, wg)

	// Sellable Product

	// Sellable Stock

	// Transaction

	if err := wg.Wait(); err != nil {
		return err
	}

	return nil
}
