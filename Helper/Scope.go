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

func FilterStatus(status *bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status == nil {
			return db
		}
		return db.Where("status = ?", status)
	}
}

func FilterTypeAccount(typeAccount string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if typeAccount == "" {
			return db
		}
		return db.Where("type = ?", typeAccount)
	}
}

func FilterIslock(isLock *bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if isLock == nil {
			return db
		}

		return db.Where("is_locked = ?", isLock)
	}
}

func FilterSearchRiwayatTransaction(query *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == nil {
			return db
		}

		return db.Where("customer_name ILIKE ? OR invoice_number ILIKE ?", "%"+*query+"%", "%"+*query+"%")
	}
}

func FilterCategoryID(categoryID *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if categoryID == nil {
			return db
		}

		return db.Where("category_id = ?", categoryID)
	}
}

func FilterUnitID(unitID string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if unitID == "" {
			return db
		}

		return db.Where("smallest_unit_id = ?", unitID)
	}
}

func FilterSearchProduct(query *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == nil {
			return db
		}

		return db.Where("name ILIKE ?", "%"+*query+"%")
	}
}

func FilterTransaksiLain(query *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == nil {
			return db
		}

		return db.Where("title ILIKE ? OR payment_type ILIKE ? OR payment_method ILIKE ?", "%"+*query+"%", "%"+*query+"%", "%"+*query+"%")
	}
}

func FilterManagementStock(query string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" {
			return db
		}

		if query == "active" {
			return db.Where("status = ?", true)
		}
		if query == "inactive" {
			return db.Where("status = ?", false)
		}
		if query == "empty" {
			return db.Where("current_quantity = ?", 0)
		}

		return db
	}
}

func FilterDateInvoice(startDate string, endDate string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fmt.Println("scope", startDate)
		if startDate == "" || endDate == "" {
			return db
		}

		return db.Where("DATE(created_at) BETWEEN ? AND ?", startDate, endDate)
	}
}

func FilterHistoryTransaction(query string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" {
			return db
		}

		if query == "paid" {
			return db.Where("refund_at IS NULL AND status = ?", true)
		}
		if query == "refund" {
			return db.Where("refund_at IS NOT NULL AND status = ?", true)
		}
		if query == "unpaid" {
			return db.Where("status = ?", false)
		}

		return db
	}
}

func FilterType(types string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if types == "" {
			return db
		}

		return db.Where("type = ?", types)
	}
}

func FilterDateInvoiceDashboard(startDate, endDate string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startDate == "" && endDate == "" {
			return db
		}

		query := db
		if startDate != "" && endDate != "" {
			query = query.Where("DATE(invoices.created_at) BETWEEN ? AND ?", startDate, endDate)
		} else if startDate != "" {
			query = query.Where("DATE(invoices.created_at) >= ?", startDate)
		} else if endDate != "" {
			query = query.Where("DATE(invoices.created_at) <= ?", endDate)
		}

		return query
	}
}
