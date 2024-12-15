package Middleware

import (
	"net/http"
	"strings"

	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type (
	ICommonMiddleware interface {
		IsAuthenticate(ctx *gin.Context)
	}

	CommondMiddleware struct {
		jwtService  Services.IJwtService
		redisClient *redis.Client
	}
)

func CommonMiddlewareProvider(jwtService Services.IJwtService, redisClient *redis.Client) *CommondMiddleware {
	return &CommondMiddleware{
		jwtService:  jwtService,
		redisClient: redisClient,
	}
}

func (m *CommondMiddleware) IsAuthenticate(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "E_UNAUTHORIZE_ACCESS"})
		ctx.Abort()
		return
	}

	if len(token) > 7 && strings.ToLower(token[:7]) == "bearer " {
		token = token[7:]
	}

	claims, err := m.jwtService.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "E_UNAUTHORIZE_ACCESS"})
		ctx.Abort()
		return
	}

	// all claims
	key, ok := claims["id"].(string)
	companyId, _ := claims["company_id"].(string)
	name, _ := claims["name"].(string)
	isEmployee, _ := claims["is_employee"].(bool)
	companyCode, _ := claims["company_code"].(string)

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "INVALID_KEY"})
		ctx.Abort()
		return
	}

	checkTokenRedis, err := m.redisClient.Get(ctx, key).Result()

	if err != nil || checkTokenRedis != token {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "E_UNAUTHORIZE_ACCESS"})
		ctx.Abort()
		return
	}

	// set to context
	ctx.Set("id", key)
	ctx.Set("company_id", companyId)
	ctx.Set("is_employee", isEmployee)
	ctx.Set("company_code", companyCode)
	ctx.Set("name", name)

	ctx.Next()
}
