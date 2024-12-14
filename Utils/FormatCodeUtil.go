package Utils

import (
	"2024_akutansi_project/Consts"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateUniqueSuffix() string {
	now := time.Now()
	timePart := fmt.Sprintf("%02d%02d", now.Minute(), now.Second())

	rand.Seed(time.Now().UnixNano())
	randomDigit := rand.Intn(10) // 0-9

	uniqueSuffix := fmt.Sprintf("%s%1d", timePart, randomDigit)
	return uniqueSuffix
}

func SuffixDigitsToUpper(value string, digits int) string {
	return strings.ToUpper(value[:digits])
}

func GenerateCodeCompany(companyName string) string {
	code := SuffixDigitsToUpper(companyName, Consts.DigitCompanyCode)
	suffix := GenerateUniqueSuffix()
	uniqueCode := fmt.Sprintf("%s-%s", code, suffix)
	return uniqueCode
}

func GenerateTransactionRecord(latestCount int64, companyCode string) string {
	if latestCount == 0 {
		latestCount = 1
	} else {
		latestCount += 1
	}

	format := fmt.Sprintf("%d-%s-%s", latestCount, companyCode, time.Now().Format("02012006"))
	return format
}
