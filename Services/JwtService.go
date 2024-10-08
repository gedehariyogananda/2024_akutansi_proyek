package Services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	IJwtService interface {
		GenerateToken(userId string, me bool) (token string, err error)
		ParseToken(token string) (claims jwt.MapClaims, err error)
		GenerateTokenWithCompany(userId string, company_id string) (token string, err error)
	}

	JwtService struct {
	}
)

func JwtServiceProvider() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GenerateToken(userId string, me bool) (token string, err error) {
	expiredTime := time.Now().Add(7 * 24 * time.Hour)

	if me {
		expiredTime = time.Now().Add(1 * 30 * 24 * time.Hour) // 1 bulan
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    expiredTime.Unix(),
	})

	token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
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
