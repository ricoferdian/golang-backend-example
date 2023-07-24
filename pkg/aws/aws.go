package aws

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func getSecretEnv() (cred *AWSCredentials, err error) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		return cred, errors.New("AWS Region is not defined")
	}
	keyID := os.Getenv("AWS_ACCESS_KEY_ID")
	if keyID == "" {
		return cred, errors.New("AWS Access Key ID is not defined")
	}
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secretKey == "" {
		return cred, errors.New("AWS Secret Access Key is not defined")
	}
	return &AWSCredentials{
		region:    region,
		keyID:     keyID,
		secretKey: secretKey,
	}, nil
}

func getAWSConn(cred *AWSCredentials) (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(cred.region),
			Credentials: credentials.NewStaticCredentials(
				cred.keyID,
				cred.secretKey,
				"",
			),
		},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (a *AWSClient) GetCloudFrontDistributionURL(s3Path string) string {
	return fmt.Sprintf("%s/%s", a.AWSConfig.CFConfig.URL, s3Path)
}
