package mail

import (
	"net/smtp"
	"strings"
)

// func SendMail() {

// 	from := "noti.awscost@gmail.com"
// 	password := "noti.awscost@Password"
// 	to := []string{"bittuanibrata@gmail.com"}
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "25"
// 	// find the port ????????????????????????????
// 	message := []byte("My super secret message.")

// 	// Create authentication
// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	// Send actual message
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func newFunction() string {
// 	smtpPort := "587"
// 	return smtpPort
// }

// func SendMail() {

// 	m := gomail.NewMessage()

// 	// Set E-Mail sender
// 	m.SetHeader("From", "noti.awscost@gmail.com")

// 	// Set E-Mail receivers
// 	m.SetHeader("To", "bittuanibrata@gmail.com")

// 	// Set E-Mail subject
// 	m.SetHeader("Subject", "Gomail test subject")

// 	// Set E-Mail body. You can set plain text or html with text/html
// 	m.SetBody("text/plain", "This is Gomail test body")

// 	// Settings for SMTP server
// 	d := gomail.NewDialer("smtp.gmail.com", 25, "noti.awscost@gmail.com", "noti.awscost@Password")

// 	// This is only needed when SSL/TLS certificate is not valid on server.
// 	// In production this should be set to false.
// 	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

// 	// Now send E-Mail
// 	if err := d.DialAndSend(m); err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}

// 	return
// }

func SendMail() {

	from := "noti.awscost@gmail.com"
	password := "noti.awscost@Password"

	toEmailAddress := "bittuanibrata@gmail.com"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "25"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", strings.TrimSpace(from), strings.TrimSpace(password), host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}

}
