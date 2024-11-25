package Helper

import (
	"2024_akutansi_project/Models/Common"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Payload    interface{} `json:"payload,omitempty"`
	Error      interface{} `json:"errors,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
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
		Message:    "invalid request body",
		StatusCode: 404,
		Error:      errors,
	}

	ctx.JSON(404, responseBody)
}

func SetPaginationResponse(ctx *gin.Context, message string, payload interface{}, totalData int64, limit int, page int, statusCode int) {
	meta := Common.PaginateMetadata(ctx, totalData, limit, page)

	responseBody := Response{
		Success:    true,
		Message:    message,
		StatusCode: statusCode,
		Payload:    payload,
		Meta:       meta,
	}

	ctx.JSON(statusCode, responseBody)
}
