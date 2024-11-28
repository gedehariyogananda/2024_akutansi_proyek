package Services

import (
	"2024_akutansi_project/Connector"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IProfileService interface {
		IntegrateShopeeProfile(ctx context.Context, request *Dto.ShopeeIntegrateRequest) (statusCode int, err error)
		IntegrateTiktokProfile(ctx context.Context, request *Dto.TiktokIntegrateRequest) (statusCode int, err error)
		GetIntegratedProfile(ctx context.Context) (resp *Dto.IntegratedProfile, statusCode int, err error)
	}

	ProfileService struct {
		ProfileRepository Repositories.IProfileRepository
		ShopeeConnector   Connector.IShopeeConnector
	}
)

func ProfileServiceProvider(profileRepository Repositories.IProfileRepository, shopeeConnector Connector.IShopeeConnector) *ProfileService {
	return &ProfileService{
		ProfileRepository: profileRepository,
		ShopeeConnector:   shopeeConnector,
	}
}

func (service *ProfileService) IntegrateShopeeProfile(ctx context.Context, request *Dto.ShopeeIntegrateRequest) (statusCode int, err error) {
	userID := ctx.Value("user_id").(string)
	integratedProfile, err := service.ProfileRepository.GetIntegratedProfile(ctx, userID)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return http.StatusBadRequest, err
		}

		integratedProfile, err = service.ProfileRepository.IntegrateProfile(ctx, Dto.IntegratedProfile{
			UserID: userID,
			Shopee: nil,
			Tiktok: nil,
		})
		if err != nil {
			return http.StatusBadRequest, err
		}
	}

	var shopeeToken Dto.Integration
	// force reintegrate
	if err = service.ShopeeConnector.SetPushNotification(ctx, &Dto.SetPushNotificationShopeeRequest{
		PartnerID:         request.PartnerID,
		PartnerKey:        request.PartnerKey,
		BlockedShopIdList: nil,
		CallbackUrl:       "",
		SetPushConfigOff:  nil,
		SetPushConfigOn:   nil,
	}); err != nil {
		return http.StatusBadRequest, err
	}

	// build integrated credential to store
	shopeeToken = Dto.Integration{
		Integrated: true,
		Credential: map[string]string{
			"partner_id":  request.PartnerID,
			"partner_key": request.PartnerKey,
		},
	}

	// store integrate status and credential on user
	_, err = service.ProfileRepository.IntegrateProfile(ctx, Dto.IntegratedProfile{
		UserID: userID,
		Shopee: &shopeeToken,
		Tiktok: integratedProfile.Tiktok,
	})

	return http.StatusOK, nil
}

func (service *ProfileService) IntegrateTiktokProfile(ctx context.Context, request *Dto.TiktokIntegrateRequest) (statusCode int, err error) {
	userID := ctx.Value("user_id").(string)
	integratedProfile, err := service.ProfileRepository.GetIntegratedProfile(ctx, userID)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return http.StatusBadRequest, err
		}

		integratedProfile, err = service.ProfileRepository.IntegrateProfile(ctx, Dto.IntegratedProfile{
			UserID: userID,
			Shopee: nil,
			Tiktok: nil,
		})
		if err != nil {
			return http.StatusBadRequest, err
		}
	}

	var tiktokToken Dto.Integration
	// force reintegrate
	// todo :: integrate tiktok

	// build integrated credential to store
	tiktokToken = Dto.Integration{
		Integrated: true,
		Credential: map[string]string{
			"client_id":     request.ClientID,
			"client_secret": request.ClientSecret,
		},
	}

	// store integrate status and credential on user
	_, err = service.ProfileRepository.IntegrateProfile(ctx, Dto.IntegratedProfile{
		UserID: userID,
		Shopee: integratedProfile.Tiktok,
		Tiktok: &tiktokToken,
	})

	return http.StatusOK, nil
}

func (service *ProfileService) GetIntegratedProfile(ctx context.Context) (resp *Dto.IntegratedProfile, statusCode int, err error) {
	userID := ctx.Value("user_id").(string)
	resp, err = service.ProfileRepository.GetIntegratedProfile(ctx, userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, http.StatusNotFound, fmt.Errorf("no document found with user_id: %s", userID)
		}
		return nil, http.StatusBadRequest, err
	}

	return resp, http.StatusOK, nil
}
