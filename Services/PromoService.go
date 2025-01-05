package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type (
	IPromoService interface {
		Create(dto *Dto.CreatePromoDto) (res Response.PromoResponse, err error)
		FindByID(id string) (res Response.PromoResponse, statusCode int, err error)
	}

	PromoService struct {
		PromoRepository     Repositories.IPromoRepository
		PromoItemRepository Repositories.IPromoItemRepository
	}
)

func PromoServiceProvider(promoRepository Repositories.IPromoRepository, promoItemRepository Repositories.IPromoItemRepository) *PromoService {
	return &PromoService{PromoRepository: promoRepository, PromoItemRepository: promoItemRepository}
}

func (s *PromoService) Create(dto *Dto.CreatePromoDto) (res Response.PromoResponse, err error) {
	promo := &Models.Promo{
		Name:      dto.Name,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
		Amount:    dto.Amount,
		CompanyID: dto.CompanyID,
		IsAll:     dto.IsAll,
		Type:      dto.Type,
	}

	promo, err = s.PromoRepository.Create(promo)
	if err != nil {
		return
	}

	if !dto.IsAll {
		for _, item := range dto.SellableProductIDS {
			promoItem := &Models.PromoItem{
				PromoID:           promo.ID,
				SellableProductID: item,
			}

			_, err = s.PromoItemRepository.Create(promoItem)
			if err != nil {
				return
			}
		}
	}

	res = Response.ToPromoResponse(*promo)

	return res, nil
}

func (s *PromoService) FindByID(id string) (res Response.PromoResponse, statusCode int, err error) {
	promo, err := s.PromoRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if err != nil {
		return
	}

	res = Response.ToPromoResponse(*promo)

	return res, http.StatusOK, nil
}
