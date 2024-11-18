package Utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(pwd string, plainPwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(plainPwd))
	if err != nil {
		return err
	}
	return nil
}
