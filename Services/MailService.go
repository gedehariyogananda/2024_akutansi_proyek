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

	EmailService struct{}
)

func EmailServiceProvider() *EmailService {
	return &EmailService{}
}

func (e *EmailService) Send(to, subject, otp string) (err error) {
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	fmt.Println("email:", email)
	fmt.Println("to:", to)
	fmt.Println("subject:", subject)

	auth := smtp.PlainAuth("", email, password, smtpHost)

	msg := "From: " + email + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n\r\n" +
		"Ini OTP Anda: " + otp + "\r\n"

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println("Failed to send email:", err)
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
