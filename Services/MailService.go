package Services

import (
	"fmt"
	"net/smtp"
	"os"
)

type (
	IEmailService interface {
		Send(to, subject, otp string) (err error)
	}

	EmailService struct {
	}
)

func EmailServiceProvider() *EmailService {
	return &EmailService{}
}

func (e *EmailService) Send(to, subject, otp string) (err error) {
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	fmt.Println("email : ", email)
	fmt.Println("password : ", password)
	fmt.Println("smtpHost : ", smtpHost)
	fmt.Println("smtpPort : ", smtpPort)

	auth := smtp.PlainAuth("", email, password, smtpHost)

	msg := []byte("From : Duit <" + email + ">\n" +
		"To : " + to + "\n" +
		"Subject : " + subject +
		"\nBody : Ini OTP anda" + otp,
	)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{to}, msg)

	if err != nil {
		return err
	}

	return
}
