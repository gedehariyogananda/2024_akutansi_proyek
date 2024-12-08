package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Repositories"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	IStockOpnameService interface {
		GetAll(query *Common.Query) (data []*Models.StockOpname, meta Common.Meta, err error)
		Create(data *Dto.CreateStockOpnameDto) (err error)
	}

	StockOpnameService struct {
		StockOpnameRepository   Repositories.IStockOpnameRepository
		DB                      *gorm.DB
		SellableStockRepository Repositories.ISellableStockRepository
	}
)

func StockOpnameServiceProvider(
	stockOpnameRepository Repositories.IStockOpnameRepository,
	sellableStockRepository Repositories.ISellableStockRepository,
	DB *gorm.DB) *StockOpnameService {
	return &StockOpnameService{
		StockOpnameRepository:   stockOpnameRepository,
		SellableStockRepository: sellableStockRepository,
		DB:                      DB}
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
	trx := s.DB.Begin()
	if trx.Error != nil {
		return trx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
			err = fmt.Errorf("panic occurred: %v", r)
		} else if err != nil {
			trx.Rollback()
		} else {
			trx.Commit()
		}
	}()

	stockOpnameId, err := uuid.NewV7()

	stockOpname := &Models.StockOpname{
		Title:       data.Title,
		CompanyID:   data.CompanyID,
		ID:          stockOpnameId.String(),
		ChangerName: &data.ChangerName,
	}

	err = s.StockOpnameRepository.Create(stockOpname, trx)
	if err != nil {
		return err
	}

	items := make([]*Models.StockOpnameItem, 0)
	for _, item := range data.Items {
		stockOpnameItem := &Models.StockOpnameItem{
			StockOpnameID:      stockOpnameId.String(),
			Quantity:           item.Quantity,
			StockID:            item.StockId,
			DifferenceQuantity: item.Quantity - item.SystemQuantity,
			ProductType:        "-",
		}

		items = append(items, stockOpnameItem)

		err = s.SellableStockRepository.UpdateCurrent(trx, item.StockId, item.SystemQuantity-item.Quantity)
	}

	err = s.StockOpnameRepository.CreateItem(items, trx)
	if err != nil {
		return err
	}

	return nil
}
