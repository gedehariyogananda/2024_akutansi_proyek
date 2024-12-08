package Services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	IJwtService interface {
		GenerateToken(id string, companyID string, name string, isEmployee bool, me bool) (token string, duration time.Duration, err error)
		ParseToken(token string) (claims jwt.MapClaims, err error)
	}

	JwtService struct {
	}
)

func JwtServiceProvider() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GenerateToken(id string, companyID string, name string, isEmployee bool, me bool) (token string, duration time.Duration, err error) {

	duration = 7 * 24 * time.Hour
	expiredTime := time.Now().Add(duration) // 1 minggu

	if me {
		duration = 100 * 365 * 24 * time.Hour // 100 tahun wkwk
		expiredTime = time.Now().Add(duration)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          id,
		"company_id":  companyID,
		"is_employee": isEmployee,
		"name":        name,
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
