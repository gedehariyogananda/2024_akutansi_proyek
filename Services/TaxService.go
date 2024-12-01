package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Repositories"
	"net/http"
)

type (
	ITaxService interface {
		GetAll() (taxs []*Models.Tax, statusCode int, err error)
	}

	TaxService struct {
		TaxRepository Repositories.ITaxRepository
	}
)

func TaxServiceProvider(taxRepository Repositories.ITaxRepository) *TaxService {
	return &TaxService{TaxRepository: taxRepository}
}

func (s *TaxService) GetAll() (taxs []*Models.Tax, statusCode int, err error) {
	taxs, err = s.TaxRepository.GetAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return taxs, http.StatusOK, nil
}
