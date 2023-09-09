package s3

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreographer"
	"github.com/Kora-Dance/koradance-backend/pkg/aws"
)

type S3ChoreographerContentRepository struct {
	awsClient *aws.AWSClient
	s3Config  *helper.S3BucketConfig
}

func NewS3ChoreographerContentRepository(awsClient *aws.AWSClient, s3Config *helper.S3BucketConfig) choreographer.S3ChoreographerContentRepo {
	return &S3ChoreographerContentRepository{
		awsClient: awsClient,
		s3Config:  s3Config,
	}
}
