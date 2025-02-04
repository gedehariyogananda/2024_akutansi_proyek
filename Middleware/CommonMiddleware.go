package Middleware

import (
	"fmt"
	"net/http"
	"strings"

	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type (
	ICommonMiddleware interface {
		IsAuthenticate(ctx *gin.Context)
		RolesAll(ctx *gin.Context)
		RoleEmployee(ctx *gin.Context)
		RoleOwner(ctx *gin.Context)
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

type Claims struct {
	Key         string
	CompanyID   string
	Name        string
	IsEmployee  bool
	CompanyCode string
}

func (m *CommondMiddleware) IsAuthenticate(ctx *gin.Context) {
	claims, err := m.extractClaims(ctx)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusUnauthorized)
		ctx.Abort()
		return
	}

	ctx.Set("id", claims.Key)
	ctx.Set("company_id", claims.CompanyID)
	ctx.Set("is_employee", claims.IsEmployee)
	ctx.Set("company_code", claims.CompanyCode)
	ctx.Set("name", claims.Name)

	ctx.Next()
}

func (m *CommondMiddleware) RolesAll(ctx *gin.Context) {
	_, err := m.roleMiddleware(ctx)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusUnauthorized)
		ctx.Abort()
		return
	}

	ctx.Next()
}

func (m *CommondMiddleware) RoleEmployee(ctx *gin.Context) {
	claims, err := m.roleMiddleware(ctx)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusUnauthorized)
		ctx.Abort()
		return
	}

	if !claims.IsEmployee {
		Helper.SetErrorResponse(ctx, "E_FORBIDDEN_ACCESS", http.StatusForbidden)
		ctx.Abort()
		return
	}

	ctx.Next()
}

func (m *CommondMiddleware) RoleOwner(ctx *gin.Context) {
	claims, err := m.roleMiddleware(ctx)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusUnauthorized)
		ctx.Abort()
		return
	}

	if claims.IsEmployee {
		Helper.SetErrorResponse(ctx, "E_FORBIDDEN_ACCESS", http.StatusForbidden)
		ctx.Abort()
		return
	}

	ctx.Next()
}

func (m *CommondMiddleware) extractClaims(ctx *gin.Context) (*Claims, error) {
	tokenClient := ctx.GetHeader("Authorization")
	if tokenClient == "" {
		return nil, fmt.Errorf("E_UNAUTHORIZE_ACCESS")
	}

	if len(tokenClient) > 7 && strings.ToLower(tokenClient[:7]) == "bearer " {
		tokenClient = tokenClient[7:]
	}

	claims, err := m.jwtService.ParseToken(tokenClient)
	if err != nil {
		return nil, fmt.Errorf("E_UNAUTHORIZE_ACCESS")
	}

	key, ok := claims["id"].(string)
	if !ok {
		return nil, fmt.Errorf("INVALID_KEY")
	}

	token, err := m.redisClient.Get(ctx, key).Result()
	if err != nil || token != tokenClient {
		return nil, fmt.Errorf("E_UNAUTHORIZE_ACCESS")
	}

	return &Claims{
		Key:         key,
		CompanyID:   claims["company_id"].(string),
		Name:        claims["name"].(string),
		IsEmployee:  claims["is_employee"].(bool),
		CompanyCode: claims["company_code"].(string),
	}, nil
}

func (m *CommondMiddleware) roleMiddleware(ctx *gin.Context) (*Claims, error) {
	claims, err := m.extractClaims(ctx)
	if err != nil {
		Helper.SetErrorResponse(ctx, err.Error(), http.StatusUnauthorized)
		ctx.Abort()
		return nil, nil
	}

	ctx.Set("id", claims.Key)
	ctx.Set("company_id", claims.CompanyID)
	ctx.Set("is_employee", claims.IsEmployee)
	ctx.Set("company_code", claims.CompanyCode)
	ctx.Set("name", claims.Name)

	return claims, nil
}

func (m *CommondMiddleware) getTokenRedis(ctx *gin.Context, key string, tokenClient string) (string, error) {
	token, err := m.redisClient.Get(ctx, key).Result()
	if err != nil || token != tokenClient {
		return "", fmt.Errorf("E_UNAUTHORIZE_ACCESS")
	}
	return token, nil
}
