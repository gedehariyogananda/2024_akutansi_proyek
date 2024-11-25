package Utils

import (
	"2024_akutansi_project/Helper"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func InitValidator() {
	Validator = validator.New()
}

func ValidateRequest(ctx *gin.Context, data interface{}) []Helper.ErrorFieldsResponse {
	var errFields []Helper.ErrorFieldsResponse

	err := Validator.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errField Helper.ErrorFieldsResponse
			switch err.Tag() {
			case "email":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = "Email tidak boleh kosong"
			case "min":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = err.Field() + " harus memiliki minimal " + err.Param() + " karakter"
			case "required":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = err.Field() + " tidak boleh kosong"
			case "max" :
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = err.Field() + " harus memiliki maksimal " + err.Param() + " karakter"
			}
			errFields = append(errFields, errField)
		}

		ctx.Set("validationErrors", errFields)
		return errFields
	}
	return nil
}
