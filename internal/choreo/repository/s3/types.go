package s3

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreo"
	"github.com/Kora-Dance/koradance-backend/pkg/aws"
)

type S3ChoreoContentRepository struct {
	awsClient *aws.AWSClient
	s3Config  *helper.S3BucketConfig
}

func NewS3ChoreoContentRepository(awsClient *aws.AWSClient, s3Config *helper.S3BucketConfig) choreo.S3ChoreoContentRepo {
	return &S3ChoreoContentRepository{
		awsClient: awsClient,
		s3Config:  s3Config,
	}
}
