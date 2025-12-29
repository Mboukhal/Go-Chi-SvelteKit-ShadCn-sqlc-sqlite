package email

import (
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)


func init() {
	// load .env
	godotenv.Load()
}

func SendEmailSys(to string, subject string, body string) error {
	// Authentication information
	smtpHost := os.Getenv("SMTP_HOST") // e.g., smtp.gmail.com
	smtpPort := os.Getenv("SMTP_PORT") // e.g., "587"
	senderEmail := os.Getenv("SMTP_EMAIL")
	senderPassword := os.Getenv("SMTP_PASSWORD") // Use an app password if using Gmail

	// println("SMTP Host:", smtpHost)
	// println("SMTP Port:", smtpPort)
	// println("Sender Email:", senderEmail)
	

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	// Email content (RFC 822-style format)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

			// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}