package s3

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

func (r *S3ChoreographerContentRepository) UploadChoreographerContent(choreographerID int64, fileName string, fileReader io.Reader) (string, error) {
	path := r.getChoreographerUUID(choreographerID) + fmt.Sprintf("choreographer-%s", fileName)
	return r.upload(path, fileReader)
}

func (r *S3ChoreographerContentRepository) getChoreographerUUID(id int64) string {
	choreoUUID := md5.Sum([]byte(strconv.FormatInt(id, 10)))
	choreoPath := fmt.Sprintf("choreographer-%x/", choreoUUID)
	path := fmt.Sprintf("%s/%s", r.s3Config.ContentConfig.Path, choreoPath)

	return path
}

func (r *S3ChoreographerContentRepository) upload(path string, fileReader io.Reader) (string, error) {
	err := r.awsClient.S3Module.UploadFile(r.s3Config.BucketName, path, fileReader)
	if err != nil {
		return "", err
	}
	return r.awsClient.GetCloudFrontDistributionURL(path), nil
}
