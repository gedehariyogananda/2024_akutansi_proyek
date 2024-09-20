package Utils

import (
	"2024_akutansi_project/Models"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"gorm.io/gorm"
)

func GenerateUniqueSuffix() string {
	now := time.Now()
	timePart := fmt.Sprintf("%02d%02d", now.Minute(), now.Second())

	rand.Seed(time.Now().UnixNano())
	randomDigit := rand.Intn(10) // 0-9

	uniqueSuffix := fmt.Sprintf("%s%1d", timePart, randomDigit)
	return uniqueSuffix
}

func GenerateCodeCompany(companyName string) string {
	code := strings.ToUpper(companyName[:3])
	suffix := GenerateUniqueSuffix()
	uniqueCode := fmt.Sprintf("%s-%s", code, suffix)
	return uniqueCode
}

// companyCode - sufixUniq - yearmonth - countNumberInvoice
// PAT-4032-202406-00001
func GenerateInvoiceNumber(db *gorm.DB, companyCode string, company_id string) (string, error) {
	yearMonth := time.Now().Format("200601")

	var countNumberInvoice int64
	if err := db.Model(&Models.Invoice{}).Where("company_id = ?", company_id).Count(&countNumberInvoice).Error; err != nil {
		return "", err
	}

	if countNumberInvoice == 0 {
		countNumberInvoice = 1
	} else {
		countNumberInvoice++
	}

	invoiceNumber := fmt.Sprintf("%s-%s-%05d", companyCode, yearMonth, countNumberInvoice)

	return invoiceNumber, nil

}
