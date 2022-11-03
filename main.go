package main

import (
	CSVFile "awscostapi/CSVHandle"
	"log"
	"net/smtp"
)

func main() {
	CSVFile.Createfile()
	SendMail()
}

// Sending Email Using Smtp in Golang

func SendMail() {

	from := "noti.awscost@gmail.com"
	password := "noti.awscost@Password"
	to := []string{"malkhandi.anibrata@tftus.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "897"

	message := []byte("My super secret message.")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
}

// func newFunction() string {
// 	smtpPort := "587"
// 	return smtpPort
// }
