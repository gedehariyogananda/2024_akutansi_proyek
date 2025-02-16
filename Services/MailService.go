package Services

import (
	"2024_akutansi_project/Models"
	"fmt"
	"net/smtp"
	"os"
)

type (
	IEmailService interface {
		Send(emailMessage Models.EmailMessage) (err error)
	}

	EmailService struct{}
)

func EmailServiceProvider() *EmailService {
	return &EmailService{}
}

func (e *EmailService) Send(emailMessage Models.EmailMessage) (err error) {
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	fmt.Println("email:", email)
	fmt.Println("to:", emailMessage.To)
	fmt.Println("subject:", emailMessage.Subject)

	auth := smtp.PlainAuth("", email, password, smtpHost)

	msg := "From: " + email + "\r\n" +
		"To: " + emailMessage.To + "\r\n" +
		"Subject: " + emailMessage.Subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
		emailMessage.Body + "\r\n"

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{emailMessage.To}, []byte(msg))

	if err != nil {
		fmt.Println("Failed to send email:", err)
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
