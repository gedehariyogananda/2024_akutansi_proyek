package Middleware

import (
	"net/http"
	"strings"

	"2024_akutansi_project/Config"
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	ICommonMiddleware interface {
		IsAuthenticate(ctx *gin.Context)
	}

	CommondMiddleware struct {
		jwtService Services.IJwtService
	}
)

func CommonMiddlewareProvider(jwtService Services.IJwtService) *CommondMiddleware {
	return &CommondMiddleware{
		jwtService: jwtService,
	}
}

func (m *CommondMiddleware) IsAuthenticate(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Token Not Found"})
		ctx.Abort()
		return
	}

	if len(token) > 7 && strings.ToLower(token[:7]) == "bearer " {
		token = token[7:]
	}

	claims, err := m.jwtService.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		ctx.Abort()
		return
	}

	// all claims
	id, ok := claims["id"].(string)
	companyId, _ := claims["companyId"].(string)
	isEmployee, _ := claims["is_employee"].(bool)

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid ID"})
		ctx.Abort()
		return
	}

	// check safety token in redis
	checkTokenRedis, err := Config.GetFromRedis(id)

	if err != nil || checkTokenRedis != token {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "UNAUTHORIZE: Token Mismatch"})
		ctx.Abort()
		return
	}

	// set to context
	ctx.Set("id", id)
	ctx.Set("companyId", companyId)
	ctx.Set("isEmployee", isEmployee)

	ctx.Next()
}
