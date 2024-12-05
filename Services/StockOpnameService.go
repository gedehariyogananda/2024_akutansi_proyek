package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Repositories"
)

type (
	IStockOpnameService interface {
		GetAll(query *Common.Query) (data []*Models.StockOpname, meta Common.Meta, err error)
	}

	StockOpnameService struct {
		StockOpnameRepository Repositories.IStockOpnameRepository
	}
)

func StockOpnameServiceProvider(stockOpnameRepository Repositories.IStockOpnameRepository) *StockOpnameService {
	return &StockOpnameService{StockOpnameRepository: stockOpnameRepository}
}

func (s *StockOpnameService) GetAll(query *Common.Query) (data []*Models.StockOpname, meta Common.Meta, err error) {
	data, totalData, err := s.StockOpnameRepository.GetAll(query)
	if err != nil {
		return nil, Common.Meta{}, err
	}

	meta = Common.Meta{
		TotalData: totalData,
		Limit:     query.Limit,
		Page:      query.Page,
	}

	return data, meta, nil
}
