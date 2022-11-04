package main

import (
	CSVFile "awscostapi/CSVHandle"
	"awscostapi/bucket"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	CSVFile.Createfile()
	time.Sleep(time.Second)
	// mail.SendMail()
	InitBucket()
	time.Sleep(time.Second)

}

// Sending Email Using Smtp in Golang
func InitBucket() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	bucketName := "awsbillaws"

	uploader := s3manager.NewUploader(sess)
	filename := "AWSBilling.csv"

	err = bucket.UploadFile(uploader, "AWSBilling.csv", bucketName, filename)
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}

	fmt.Println("Successfully uploaded file!")

}
