package Services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	IJwtService interface {
		GenerateToken(id string, company_id string, is_employee bool, me bool) (token string, duration time.Duration, err error)
		ParseToken(token string) (claims jwt.MapClaims, err error)
		GenerateTokenWithCompany(userId string, company_id string) (token string, err error)
	}

	JwtService struct {
	}
)

func JwtServiceProvider() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GenerateToken(id string, company_id string, is_employee bool, me bool) (token string, duration time.Duration, err error) {

	duration = 7 * 24 * time.Hour
	expiredTime := time.Now().Add(duration) // 1 minggu

	if me {
		duration = 30 * 24 * time.Hour
		expiredTime = time.Now().Add(duration) // 1 bulan
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          id,
		"companyId":   company_id,
		"is_employee": is_employee,
		"exp":         expiredTime.Unix(),
	})

	token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", 0, fmt.Errorf("failed to sign JWT token: %w", err)
	}

	return token, duration, nil
}

func (s *JwtService) ParseToken(token string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println("error parse token : ", err)
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, err
	}

	return claims, nil
}

func (s *JwtService) GenerateTokenWithCompany(userId string, company_id string) (token string, err error) {
	expiredTime := time.Now().Add(1 * 30 * 24 * time.Hour) // 1 bulan

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"companyId": company_id,
		"exp":       expiredTime.Unix(),
	})

	token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
