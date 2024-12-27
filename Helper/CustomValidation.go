package Helper

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

func DateFormat(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	if dateStr == "" {
		fmt.Println("date str kosong")
		return true
	}
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
