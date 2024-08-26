package Middleware

import (
	"net/http"
	"strings"

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
	return &CommondMiddleware{jwtService: jwtService}
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

	ctx.Set("user_id", claims["userId"])
	ctx.Next()
}
