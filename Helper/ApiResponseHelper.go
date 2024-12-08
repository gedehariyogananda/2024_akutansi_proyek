package Helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Payload    interface{} `json:"payload,omitempty"`
	Error      interface{} `json:"errors,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
}

type ErrorFieldsResponse struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}

func SetSuccessResponse(ctx *gin.Context, message string, payload interface{}, statusCode int) {
	responseBody := Response{
		Success:    true,
		Message:    message,
		StatusCode: statusCode,
		Payload:    payload,
	}

	ctx.JSON(statusCode, responseBody)
}

func SetErrorResponse(ctx *gin.Context, message string, statusCode int) {
	responseBody := Response{
		Success:    false,
		Message:    message,
		StatusCode: statusCode,
	}

	ctx.JSON(statusCode, responseBody)
}

func SetValidationErrorResponse(ctx *gin.Context, errors interface{}) {
	responseBody := Response{
		Success:    false,
		Message:    "E_VALIDATION_EXCEPTION",
		StatusCode: 422,
		Error:      errors,
	}

	ctx.JSON(422, responseBody)
}

func SetPaginationResponse(ctx *gin.Context, message string, payload interface{}, meta interface{}, statusCode int) {
	fmt.Println("meta", meta)
	responseBody := Response{
		Success:    true,
		Message:    message,
		StatusCode: statusCode,
		Payload:    payload,
		Meta:       meta,
	}

	ctx.JSON(statusCode, responseBody)
}
