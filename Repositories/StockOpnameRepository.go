package Repositories

import (
	"2024_akutansi_project/Models"
	"2024_akutansi_project/Models/Common"
	"2024_akutansi_project/Utils"
	"gorm.io/gorm"
)

type (
	IStockOpnameRepository interface {
		GetAll(query *Common.Query) (data []*Models.StockOpname, totalData int64, err error)
		Create(data *Models.StockOpname, trx *gorm.DB) (err error)
		CreateItem(data []*Models.StockOpnameItem, trx *gorm.DB) (err error)
		GetById(stockOpnameID string) (data *Models.StockOpname, err error)
	}

	StockOpnameRepository struct {
		DB *gorm.DB
	}
)

func StockOpnameRepositoryProvider(db *gorm.DB) *StockOpnameRepository {
	return &StockOpnameRepository{DB: db}
}

func (r *StockOpnameRepository) GetAll(query *Common.Query) (data []*Models.StockOpname, totalData int64, err error) {
	if err := r.DB.
		Where("company_id = ?", query.CompanyID).
		Scopes(
			Utils.Paginate(query.Page, query.Limit)).
		Order("created_at desc").
		Find(&data).Error; err != nil {
		return nil, 0, err
	}

	if data == nil || len(data) == 0 {
		err = &Utils.NotFoundError{Message: "Tidak ditemukan data"}
		return nil, 0, err
	}

	totalData, err = Utils.CountModelRecords(r.DB, &data)
	if err != nil {
		return nil, 0, err
	}

	return data, totalData, nil
}

func (r *StockOpnameRepository) Create(data *Models.StockOpname, trx *gorm.DB) (err error) {
	db := trx
	if db == nil {
		db = r.DB
	}

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (r *StockOpnameRepository) CreateItem(data []*Models.StockOpnameItem, trx *gorm.DB) (err error) {
	db := trx
	if db == nil {
		db = r.DB
	}

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (r *StockOpnameRepository) GetById(stockOpnameID string) (data *Models.StockOpname, err error) {
	var stockOpname Models.StockOpname

	if err := r.DB.
		Preload("Items").
		Where("id = ?", stockOpnameID).
		First(&stockOpname).Error; err != nil {
		return nil, err
	}

	return &stockOpname, nil
}
