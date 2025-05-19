package aws_func

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3BucketName = "linkbush-images"

// UploadToS3 uploads a file to the specified S3 bucket.
func UploadToS3(file multipart.File, key string, client *s3.Client) error {
	// Create a buffer to read the file content
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Create an S3 client
	// cfg, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	return fmt.Errorf("failed to load default config: %w", err)
	// }
	// client := s3.NewFromConfig(cfg)

	// Upload the file to S3
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(S3BucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}

	return nil
}
