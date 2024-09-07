package Helper

import (
	"github.com/gin-gonic/gin"
)

func SetResponse(ctx *gin.Context, responseBody gin.H, statusCode int) {
	ctx.Set("response_body", responseBody)
	ctx.Set("status_code", statusCode)
}
