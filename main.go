package main

import (
	CSVFile "awscostapi/CSVHandle"
	"awscostapi/bucket"
	"awscostapi/mail"
)

func main() {
	CSVFile.Createfile()
	
	mail.SendMail()
	bucket.InitBucket()

}

// Sending Email Using Smtp in Golang
