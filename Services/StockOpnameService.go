package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
)

type (
	IStockOpnameService interface {
		GetAll(query *Common.Query) (data []*Models.StockOpname, meta Common.Meta, err error)
		Create(data *Dto.CreateStockOpnameDto) (err error)
		Update(data *Dto.UpdateStockOpnameDto) (err error)
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

	meta = Common.PaginateMetadata(nil, totalData, query.Limit, query.Page)

	return data, meta, nil
}

func (s *StockOpnameService) Create(data *Dto.CreateStockOpnameDto) (err error) {
	stockOpname := &Models.StockOpname{
		Title:     data.Title,
		CompanyID: data.CompanyID,
	}

	err = s.StockOpnameRepository.Create(stockOpname)
	if err != nil {
		return err
	}

	return nil
}

func (s *StockOpnameService) Update(data *Dto.UpdateStockOpnameDto) (err error) {
	stockOpname := &Models.StockOpname{
		ID:          data.ID,
		ChangerName: &data.User,
	}

	err = s.StockOpnameRepository.Update(stockOpname)
	if err != nil {
		return err
	}

	return nil
}
