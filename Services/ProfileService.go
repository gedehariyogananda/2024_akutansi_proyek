package Services

import (
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type (
	IProfileService interface {
		IntegrateShopeeProfile(ctx context.Context, request *Dto.ShopeeIntegrateRequest) (statusCode int, err error)
		GetIntegratedProfile(ctx context.Context) (resp *Dto.IntegratedProfile, statusCode int, err error)
	}

	ProfileService struct {
		ProfileRepository Repositories.IProfileRepository
	}
)

func ProfileServiceProvider(paymentMethodRepository Repositories.IProfileRepository) *ProfileService {
	return &ProfileService{ProfileRepository: paymentMethodRepository}
}

func (service *ProfileService) IntegrateShopeeProfile(ctx context.Context, request *Dto.ShopeeIntegrateRequest) (statusCode int, err error) {
	userID := ctx.Value("user_id").(string)
	integratedProfile, err := service.ProfileRepository.GetIntegratedProfile(ctx, userID)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return http.StatusBadRequest, err
		}

		integratedProfile, err = service.ProfileRepository.IntegrateProfile(ctx, Dto.IntegratedProfile{
			UserID:      userID,
			ShopeeToken: nil,
			TiktokToken: nil,
		})
		if err != nil {
			return http.StatusBadRequest, err
		}
	}

	var shopeeToken string
	if integratedProfile.ShopeeToken == nil {
		// todo :: integrate to get shopee token and adjust mock token
		shopeeToken = "1234qwerasdf5678"
	} else {
		shopeeToken = *integratedProfile.ShopeeToken
	}

	_, err = service.ProfileRepository.IntegrateProfile(ctx, Dto.IntegratedProfile{
		UserID:      userID,
		ShopeeToken: &shopeeToken,
		TiktokToken: integratedProfile.TiktokToken,
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
