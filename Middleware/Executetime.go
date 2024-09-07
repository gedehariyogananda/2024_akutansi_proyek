package Middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func ExecutionTimeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		// Calculate execution time
		executionTime := time.Since(startTime).Milliseconds()
		fmt.Printf("Middleware finished, execution time: %d ms\n", executionTime)

		responseBody, exists := ctx.Get("response_body")
		statusCode, _ := ctx.Get("status_code")

		if exists {
			if responseMap, ok := responseBody.(gin.H); ok {
				responseMap["execution_time"] = fmt.Sprintf("%d ms", executionTime)
				responseMap["app_version"] = os.Getenv("APP_VERSION")
				responseMap["app_environment"] = os.Getenv("APP_ENV")
				ctx.JSON(statusCode.(int), responseMap)
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "internal server error",
				})
			}
		}
	}
}
