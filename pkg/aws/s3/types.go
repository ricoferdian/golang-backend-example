package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
)

type S3ClientInterface interface {
	UploadFile(bucketName string, fileName string, fileReader io.Reader) error
}

type S3Module struct {
	uploaderM   *s3manager.Uploader
	downloaderM *s3manager.Downloader
}

func NewS3Module(sess *session.Session) S3ClientInterface {
	uploaderM := s3manager.NewUploader(sess)
	downloaderM := s3manager.NewDownloader(sess)
	return &S3Module{
		uploaderM:   uploaderM,
		downloaderM: downloaderM,
	}
}
