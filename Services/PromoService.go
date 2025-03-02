package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	IPromoService interface {
		Create(dto *Dto.CreatePromoDto) (res Response.PromoResponse, err error)
		CreatePromoOnly(dto *Dto.CreatePromoOnly) (res Response.PromoResponse, err error)
		FindByID(id string) (res Response.PromoResponse, statusCode int, err error)
		Delete(id string) (statusCode int, err error)
		Update(dto *Dto.UpdatePromoDto, id string) (res Response.PromoResponse, statusCode int, err error)
		FindAll(companyID string, query Common.Query) (res []Response.PromoResponse, meta Common.Meta, err error)
		AsginPromo(dto *Dto.AsignPromoDto) (res Response.PromoResponse, satusCode int, err error)
		checkAvailableProduct(promoItems []Models.PromoItem) bool
		asignPromoToAllProduct(companyID string, promoID string) error
		UnasignPromo(dto *Dto.AsignPromoDto) (statusCode int, err error)
	}

	PromoService struct {
		PromoRepository           Repositories.IPromoRepository
		PromoItemRepository       Repositories.IPromoItemRepository
		SellableProductRepository Repositories.ISellableProductRepository
	}
)

func PromoServiceProvider(promoRepository Repositories.IPromoRepository, promoItemRepository Repositories.IPromoItemRepository, sellableProductRepository Repositories.ISellableProductRepository) *PromoService {
	return &PromoService{
		PromoRepository:           promoRepository,
		PromoItemRepository:       promoItemRepository,
		SellableProductRepository: sellableProductRepository,
	}
}

func (s *PromoService) Create(dto *Dto.CreatePromoDto) (res Response.PromoResponse, err error) {
	startDate, err := time.Parse("2006-01-02", dto.StartDate)

	if err != nil {
		return
	}

	endDate, err := time.Parse("2006-01-02", dto.EndDate)

	if err != nil {
		return
	}

	promo := &Models.Promo{
		Name:      dto.Name,
		StartDate: startDate,
		EndDate:   endDate,
		Amount:    dto.Amount,
		CompanyID: dto.CompanyID,
		IsAll:     dto.IsAll,
		Type:      dto.Type,
	}

	promo, err = s.PromoRepository.Create(promo)
	if err != nil {
		return
	}

	promoItems, err := s.PromoItemRepository.FindByPromoID(promo.ID)

	if err != nil {
		return res, err
	}
	if !dto.IsAll {

		if !s.checkAvailableProduct(promoItems) {
			return res, errors.New("product is not available")
		}

		for _, item := range dto.SellableProductIDS {
			promoItem := &Models.PromoItem{
				PromoID:           promo.ID,
				SellableProductID: item,
			}

			_, err = s.PromoItemRepository.Create(promoItem)
			if err != nil {
				return res, err
			}
		}
	} else {
		err = s.asignPromoToAllProduct(dto.CompanyID, promo.ID)
		if err != nil {
			return res, err
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

	startDate, err := time.Parse("2006-01-02", dto.StartDate)

	if err != nil {
		return
	}

	endDate, err := time.Parse("2006-01-02", dto.EndDate)

	if err != nil {
		return
	}

	promo.Name = dto.Name
	promo.StartDate = startDate
	promo.EndDate = endDate
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

func (s *PromoService) CreatePromoOnly(dto *Dto.CreatePromoOnly) (res Response.PromoResponse, err error) {
	startDate, err := time.Parse("2006-01-02", dto.StartDate)

	if err != nil {
		return
	}

	endDate, err := time.Parse("2006-01-02", dto.EndDate)

	if err != nil {
		return
	}

	promo := &Models.Promo{
		Name:      dto.Name,
		StartDate: startDate,
		EndDate:   endDate,
		Amount:    dto.Amount,
		CompanyID: dto.CompanyID,
		IsAll:     false,
		Type:      dto.Type,
	}

	promo, err = s.PromoRepository.Create(promo)
	if err != nil {
		return
	}

	res = Response.ToPromoResponse(*promo)

	return res, nil
}

func (s *PromoService) AsginPromo(dto *Dto.AsignPromoDto) (res Response.PromoResponse, statusCode int, err error) {
	promo, err := s.PromoRepository.FindById(dto.PromoID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, http.StatusNotFound, err
	}

	if err != nil {
		return
	}

	promoItems, err := s.PromoItemRepository.FindByPromoID(dto.PromoID)

	if err != nil {
		return
	}
	if !s.checkAvailableProduct(promoItems) {
		return res, http.StatusBadRequest, errors.New("product is not available")
	}

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

	res = Response.ToPromoResponse(*promo)

	return res, http.StatusOK, nil
}

func (s *PromoService) checkAvailableProduct(promoItems []Models.PromoItem) bool {
	for _, item := range promoItems {
		endDate := item.Promo.EndDate

		if endDate.Before(time.Now()) {
			return false
		}
	}

	return true

}

func (s *PromoService) asignPromoToAllProduct(companyID string, promoID string) error {
	products, err := s.SellableProductRepository.GetByCompanyID(companyID)

	if err != nil {
		return err
	}

	var promoItems []Models.PromoItem

	for _, product := range products {
		// Variabel flag untuk menentukan apakah promo item harus dibuat
		createPromo := false

		if len(product.PromoItems) > 0 {
			// Jika ada promo item, cek apakah setidaknya salah satu promo belum berakhir
			for _, promoItem := range product.PromoItems {
				if promoItem.Promo.EndDate.After(time.Now()) {
					createPromo = true
					break
				}
			}
		} else {
			// Jika tidak ada promo item, maka langsung buat promo item baru
			createPromo = true
		}

		// Jika flag createPromo bernilai true, buat promo item baru
		if createPromo {
			uuid, err := uuid.NewV7()

			if err != nil {
				return err
			}

			newPromoItem := &Models.PromoItem{
				PromoID:           promoID,
				SellableProductID: product.ID,
				ID:                uuid.String(),
			}

			promoItems = append(promoItems, *newPromoItem)
		}

	}
	if err := s.PromoItemRepository.BulkCreate(promoItems); err != nil {
		return err
	}

	return nil
}

func (s *PromoService) UnasignPromo(dto *Dto.AsignPromoDto) (statusCode int, err error) {
	_, err = s.PromoRepository.FindById(dto.PromoID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, err
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	for _, item := range dto.SellableProductIDS {
		err = s.PromoItemRepository.DeleteBySellableProductID(item)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}
