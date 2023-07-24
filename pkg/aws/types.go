package aws

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/pkg/aws/s3"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AWSCredentials struct {
	region    string
	keyID     string
	secretKey string
}

type AWSClient struct {
	Cred      *AWSCredentials
	Session   *session.Session
	S3Module  s3.S3ClientInterface
	AWSConfig *helper.AWSConfig
}

func NewAWSModule(config *helper.AWSConfig) (*AWSClient, error) {
	cred, err := getSecretEnv()
	if err != nil {
		return nil, err
	}
	sess, err := getAWSConn(cred)
	if err != nil {
		return nil, err
	}
	s3Module := s3.NewS3Module(sess)

	module := &AWSClient{
		Cred:      cred,
		Session:   sess,
		S3Module:  s3Module,
		AWSConfig: config,
	}
	return module, nil
}
