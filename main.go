package main

import (
	"awscostapi/CreateFile"
	"log"
	"net/smtp"
)

func main() {
	CreateFile.Createfile()
	SendMail()
}

// Sending Email Using Smtp in Golang

func SendMail() {

	

	from := "example@gmail.com"
	password := "example@Password"
	to := []string{"malkhandi.anibrata@example.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "465"

	message := []byte("My super secret message.")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
}
