package Helper

import (
	"fmt"

	"gorm.io/gorm"
)

func FilterSearch(query string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" {
			return db
		}
		fmt.Println("query", query)
		return db.Where("name ILIKE ?", "%"+query+"%")
	}
}

func FilterCompanyID(companyID string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if companyID == "" {
			return db
		}
		return db.Where("company_id = ?", companyID)
	}
}

func FilterStatus(status bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}

func FilterSearchRiwayatTransaction(query string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" {
			return db
		}

		return db.Where("customer_name ILIKE ? OR invoice_number ILIKE ?", "%"+query+"%", "%"+query+"%")
	}
}
