package siteback

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"os"
)

// Upload sur S3
func S3(f io.Reader, bucket string, filename string) (string, error) {
	cfg := aws.NewConfig().
		WithRegion("fr-par").
		WithEndpoint("https://s3.fr-par.scw.cloud").
		WithCredentials(credentials.NewStaticCredentials(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), ""))
	sess := session.Must(session.NewSession(cfg))
	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	return result.Location, nil
}
