package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
)

func (s S3Module) UploadFile(bucketName string, fileName string, fileReader io.Reader) error {
	//upload to the s3 bucket
	_, err := s.uploaderM.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   fileReader,
	})
	if err != nil {
		return err
	}
	return nil
}
