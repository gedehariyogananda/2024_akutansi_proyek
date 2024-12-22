package Helper

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func DateFormat(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	return err == nil
}
