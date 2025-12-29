package main

import (
	"github.com/Mboukhal/FactoryBase/internal/adapter/email"
	"github.com/joho/godotenv"
)

func main() {

	// load .env and other setup code here
	errenv := godotenv.Load()
	if errenv != nil {
		println("Error loading .env file")
	}
	// send email test
	err := email.SendEmailSys("lios80466@gmail.com", "Test Subject", "This is a test email body.")
	if err != nil {
		panic("Failed to send email: " + err.Error())
	}
	println("Email sent successfully!")
}
