package main

import (
	CSVFile "awscostapi/CSVHandle"
	"awscostapi/mail"
)

func main() {
	CSVFile.Createfile()
	mail.SendMail()
}

// Sending Email Using Smtp in Golang
