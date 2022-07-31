package upload

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadFile(file *os.File, options *UploadOptions) {
	sess := session.Must(session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(options.PATH_STYLE),
		Region:           aws.String(endpoints.ApSoutheast1RegionID),
		Endpoint:         aws.String(options.ENDPOINT),
		Credentials:      credentials.NewStaticCredentials(options.AK, options.SK, options.TOKEN),
	}))
	svc := s3.New(sess)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	_, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(options.BUCKET),
		Key:    aws.String(file.Name()),
		Body:   file,
	})

	if err != nil {
		panic(err)
	}
}

func uploadFilesByFilePath(filePath string, options *UploadOptions) {
	for _, filePath := range strings.Split(filePath, ",") {
		filePath = strings.Trim(filePath, " ")
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		fmt.Println("uploading......" + filePath)
		uploadFile(file, options)
	}
	fmt.Println("Upload Done")
}
