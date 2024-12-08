package Helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success    bool        `json:"success,omitempty"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
	Error      interface{} `json:"errors,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
}

type ErrorFieldsResponse struct {
	FieldName string `json:"field_name"`
	Message   string `json:"message"`
}

func SetSuccessResponse(ctx *gin.Context, message string, payload interface{}, statusCode int) {
	responseBody := Response{
		Message: message,
		Payload: payload,
	}

	ctx.JSON(statusCode, responseBody)
}

func SetErrorResponse(ctx *gin.Context, message string, statusCode int) {
	responseBody := Response{
		Message: message,
	}

	ctx.JSON(statusCode, responseBody)
}

func SetValidationErrorResponse(ctx *gin.Context, errors interface{}) {
	responseBody := Response{
		Message: "E_VALIDATION_EXCEPTION",
		Error:   errors,
	}

	ctx.JSON(422, responseBody)
}

func SetPaginationResponse(ctx *gin.Context, message string, payload interface{}, meta interface{}, statusCode int) {
	fmt.Println("meta", meta)
	responseBody := Response{
		Message: message,
		Payload: payload,
		Meta:    meta,
	}

	ctx.JSON(statusCode, responseBody)
}
