package Helper

import "gorm.io/gorm"

func FilterSearch(query string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" {
			return db
		}
		return db.Where("category_name LIKE ?", "%"+query+"%")
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
