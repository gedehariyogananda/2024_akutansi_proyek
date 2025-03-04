package Services

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Models/Dto"
	"2024_akutansi_project/Models/Dto/Response"
	"2024_akutansi_project/Repositories"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sync"
	"time"
)

type (
	IStockOpnameService interface {
		GetAll(query *Common.Query) (data []Response.StockOpnameAllResponse, meta Common.Meta, err error)
		Create(data *Dto.CreateStockOpnameDto) (err error)
		GetAvailableStock(companyID string) (data []*Response.AvailableStockResponse, err error)
		GetById(stockOpnameID string) (data *Response.StockOpnameResponse, err error)
	}

	StockOpnameService struct {
		StockOpnameRepository   Repositories.IStockOpnameRepository
		DB                      *gorm.DB
		SellableStockRepository Repositories.ISellableStockRepository
		MaterialStockRepository Repositories.IMaterialStockRepository
	}
)

func StockOpnameServiceProvider(
	stockOpnameRepository Repositories.IStockOpnameRepository,
	sellableStockRepository Repositories.ISellableStockRepository,
	MaterialstockRepository Repositories.IMaterialStockRepository,
	DB *gorm.DB) *StockOpnameService {
	return &StockOpnameService{
		StockOpnameRepository:   stockOpnameRepository,
		SellableStockRepository: sellableStockRepository,
		MaterialStockRepository: MaterialstockRepository,
		DB:                      DB}
}

func (s *StockOpnameService) GetAll(query *Common.Query) (data []Response.StockOpnameAllResponse, meta Common.Meta, err error) {
	result, totalData, err := s.StockOpnameRepository.GetAll(query)
	if err != nil {
		return nil, Common.Meta{}, err
	}

	meta = Common.PaginateMetadata(nil, totalData, query.Limit, query.Page)

	data = Response.MapFromStockOpnameAll(result)

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
		var systemQuantity int
		var expiredDate time.Time
		var initialQuantity int

		if item.Type == "material" {
			materialStock, err := s.MaterialStockRepository.Get(item.StockId)
			if err != nil {
				return err
			}

			systemQuantity = materialStock.CurrentQuantity
			expiredDate = materialStock.ExpiredDate
			initialQuantity = materialStock.Quantity
		} else {
			sellableStock, err := s.SellableStockRepository.Get(item.StockId)
			if err != nil {
				return err
			}

			systemQuantity = sellableStock.CurrentQuantity
			expiredDate = sellableStock.ExpiredDate
			initialQuantity = sellableStock.Quantity
		}

		stockOpnameItem := &Models.StockOpnameItem{
			StockOpnameID:      stockOpnameId.String(),
			Quantity:           item.Quantity,
			StockID:            item.StockId,
			DifferenceQuantity: item.Quantity - systemQuantity,
			ExpiredDate:        expiredDate,
			ProductType:        item.Type,
			Name:               item.Name,
			InitialQuantity:    initialQuantity,
		}

		items = append(items, stockOpnameItem)

		if item.Type == "material" {
			err = s.MaterialStockRepository.UpdateCurrent(trx, item.StockId, item.SystemQuantity-item.Quantity)
		} else {
			err = s.SellableStockRepository.UpdateCurrent(trx, item.StockId, item.SystemQuantity-item.Quantity)
		}

		if err != nil {
			return err
		}
	}

	err = s.StockOpnameRepository.CreateItem(items, trx)
	if err != nil {
		return err
	}

	return nil
}

func (s *StockOpnameService) GetAvailableStock(companyID string) (data []*Response.AvailableStockResponse, err error) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		stock, err := s.SellableStockRepository.GetAvailableStock(companyID)
		if err != nil {
			return
		}
		data = append(data, Response.ToSellStockSlice(stock)...)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		materialStock, err := s.MaterialStockRepository.GetAvailableStock(companyID)
		if err != nil {
			return
		}
		data = append(data, Response.ToMaterialStockSlice(materialStock)...)
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *StockOpnameService) GetById(stockOpnameID string) (data *Response.StockOpnameResponse, err error) {
	stockOpname, e := s.StockOpnameRepository.GetById(stockOpnameID)
	if e != nil {
		return nil, e
	}

	data = Response.MapFromStockOpname(*stockOpname)

	return data, nil
}
