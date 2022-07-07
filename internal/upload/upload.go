package upload

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadFile(file *os.File, options *UploadOptions) {
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

func UploadFiles(files []*os.File, options *UploadOptions) {
	for _, file := range files {
		UploadFile(file, options)
	}
}
