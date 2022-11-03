package bucket

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/evalphobia/aws-sdk-go-wrapper/s3"
)

func InitBucket() {
	// creating bucket for string csv file.........
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("ap-south-1"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	// initialize s3 bucket
	s3Client := s3.New()

	bucketName := "AwsCostAPI"
	err = bucket.CreateBucket(s3Client, bucketName)
	if err != nil {
		fmt.Printf("Couldn't create bucket: %v", err)
		return
	}

	fmt.Println("Successfully created bucket")

	//  uploading download csv file into s3 bucket.....
	uploader := s3manager.NewUploader(sess)
	filename := "BillingApp.csv"

	err = bucket.UploadFile(uploader, "BillingApp.csv", bucketName, filename)
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}

	fmt.Println("Successfully uploaded file!")

}
