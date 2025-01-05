package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
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
		Delete(id string) (statusCode int, err error)
		Update(dto *Dto.UpdatePromoDto, id string) (res Response.PromoResponse, statusCode int, err error)
		FindAll(companyID string, query Common.Query) (res []Response.PromoResponse, meta Common.Meta, err error)
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

func (s *PromoService) Delete(id string) (statusCode int, err error) {
	_, err = s.PromoRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, err
	}

	err = s.PromoRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = s.PromoItemRepository.DeleteByPromoID(id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *PromoService) Update(dto *Dto.UpdatePromoDto, id string) (res Response.PromoResponse, statusCode int, err error) {
	promo, err := s.PromoRepository.FindById(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if err != nil {
		return
	}

	promo.Name = dto.Name
	promo.StartDate = dto.StartDate
	promo.EndDate = dto.EndDate
	promo.Amount = dto.Amount
	promo.IsAll = dto.IsAll
	promo.Type = dto.Type
	promo.CompanyID = dto.CompanyID

	promo, err = s.PromoRepository.Update(promo, id)
	if err != nil {
		return
	}

	err = s.PromoItemRepository.DeleteByPromoID(id)
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
	res.ID = id

	return res, http.StatusOK, nil

}

func (s *PromoService) FindAll(companyID string, query Common.Query) (res []Response.PromoResponse, meta Common.Meta, err error) {
	promos, total, err := s.PromoRepository.FindAll(companyID, query)

	if err != nil {
		return
	}

	res = Response.ToPromoResponseSlice(promos)
	meta.TotalData = total

	return res, meta, nil
}
