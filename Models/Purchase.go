package Models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var monthNames = map[time.Month]string{
	1:  "Januari",
	2:  "Februari",
	3:  "Maret",
	4:  "April",
	5:  "Mei",
	6:  "Juni",
	7:  "Juli",
	8:  "Agustus",
	9:  "September",
	10: "Oktober",
	11: "November",
	12: "Desember",
}

type Purchase struct {
	ID                  string     `json:"id"`
	TotalPurchaseAmount float32    `json:"total_purchase_amount"`
	PurchaseNumber      string     `json:"purchase_number"`
	CompanyID           string     `json:"company_id"`
	Tax                 float32    `json:"tax"`
	Discount            float32    `json:"discount"`
	IsDiscountPercent   bool       `json:"is_discount_percent"`
	Payment             string     `json:"payment_type"`
	DueDate             *time.Time `json:"due_date"`
	Note                *string    `json:"note"`
	CreatedAt           time.Time  `json:"created_at"`
}

func (p *Purchase) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}

	if p.PurchaseNumber == "" {
		purchaseNumber, err := p.generateTransactionNumber(tx)
		if err != nil {
			return err
		}

		p.PurchaseNumber = purchaseNumber
	}

	return
}

func (p *Purchase) generateTransactionNumber(db *gorm.DB) (string, error) {
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	monthName := monthNames[currentMonth]

	var lastPurchase Purchase
	fmt.Println(p.CompanyID)
	likePattern := fmt.Sprintf("%d/%s/%%", currentYear, monthName)
	err := db.Where("purchase_number LIKE ?", likePattern).
		Where("company_id = ?", p.CompanyID).
		Order("purchase_number DESC").
		First(&lastPurchase).Error
	if err != gorm.ErrRecordNotFound && err != nil {

		return "", err
	}

	var newNumber int

	if err == gorm.ErrRecordNotFound {
		newNumber = 1
	} else {
		lastNumber := lastPurchase.PurchaseNumber[len(lastPurchase.PurchaseNumber)-4:]
		fmt.Sscanf(lastNumber, "%d", &newNumber)
		newNumber++
	}
	fmt.Println(newNumber)
	fromattedNumber := fmt.Sprintf("%d/%s/%04d", currentYear, monthName, newNumber)
	return fromattedNumber, nil
}
