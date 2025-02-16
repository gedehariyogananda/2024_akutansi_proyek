package Services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	IJwtService interface {
		GenerateToken(id string, companyID string, name string, companyCode string, isEmployee bool, me bool) (token string, duration time.Duration, err error)
		ParseToken(token string) (claims jwt.MapClaims, err error)
		GenerateTokenForOtp(userID string, otp string) (token string, err error)
		ParseTokenOtp(token string) (claims jwt.MapClaims, err error)
		GenerateTokenForVerificationAccount(email string) (token string, err error)
		ParseTokenVerificationAccount(token string) (claims jwt.MapClaims, err error)
	}

	JwtService struct {
	}
)

func JwtServiceProvider() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GenerateToken(id string, companyID string, name string, companyCode string, isEmployee bool, me bool) (token string, duration time.Duration, err error) {

	duration = 7 * 24 * time.Hour
	expiredTime := time.Now().Add(duration) // 1 minggu

	if me {
		duration = 100 * 365 * 24 * time.Hour // 100 tahun wkwk
		expiredTime = time.Now().Add(duration)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":           id,
		"company_id":   companyID,
		"is_employee":  isEmployee,
		"name":         name,
		"company_code": companyCode,
		"exp":          expiredTime.Unix(),
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

func (s *JwtService) GenerateTokenForOtp(userId string, otp string) (token string, err error) {
	duration := 1 * time.Minute
	expiredTime := time.Now().Add(duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"otp": otp,
		"exp": expiredTime.Unix(),
	})

	token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET_OTP")))
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
	}

	return token, nil
}

func (s *JwtService) ParseTokenOtp(token string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_OTP")), nil
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

func (s *JwtService) GenerateTokenForVerificationAccount(email string) (token string, err error) {
	duration := 1 * time.Minute
	expiredTime := time.Now().Add(duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   expiredTime.Unix(),
	})

	token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET_VERIFICATION_ACCOUNT")))
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
	}

	return token, nil
}

func (s *JwtService) ParseTokenVerificationAccount(token string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_VERIFICATION_ACCOUNT")), nil
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
