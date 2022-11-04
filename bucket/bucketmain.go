package bucket

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
)

func InitBucket() {
sess, err := session.NewSessionWithOptions(session.Options{
	Profile: "default",
	Config: aws.Config{
		Region: aws.String("us-east-2"),
	},
})

if err != nil {
	fmt.Printf("Failed to initialize new session: %v", err)
	return
}

s3Client := s3.New(sess)

bucketName := "AWSBILL"
err = CreateBucket(s3Client, bucketName)
if err != nil {
	fmt.Printf("Couldn't create bucket: %v", err)
	return
}

fmt.Println("Successfully created bucket")

uploader := s3manager.NewUploader(sess)
filename := "AWSBilling.csv"

err = UploadFile(uploader, "AWSBilling.csv", bucketName, filename)
if err != nil {
	fmt.Printf("Failed to upload file: %v", err)
}

fmt.Println("Successfully uploaded file!")

}
