package Middleware

import (
	"fmt"
	"net/http"
	"strings"

	"2024_akutansi_project/Repositories"
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	ICommonMiddleware interface {
		IsAuthenticate(ctx *gin.Context)
	}

	CommondMiddleware struct {
		jwtService     Services.IJwtService
		authRepository Repositories.IAuthRepository
	}
)

func CommonMiddlewareProvider(jwtService Services.IJwtService, authRespository Repositories.IAuthRepository) *CommondMiddleware {
	return &CommondMiddleware{
		jwtService:     jwtService,
		authRepository: authRespository,
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

	userID, ok := claims["userId"].(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid User ID"})
		ctx.Abort()
		return
	}

	ctx.Set("user_id", userID)

	if companyID, ok := claims["companyId"].(string); ok {
		ctx.Set("company_id", companyID)
	} else {
		fmt.Println("companyId not found in claims")
	}

	if err = m.authRepository.CheckToken(token, userID); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: Token Mismatch"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
