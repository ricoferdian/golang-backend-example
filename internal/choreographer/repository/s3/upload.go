package s3

import (
	"io"
)

func (s S3ChoreographerContentRepository) UploadChoreographerContent(choreographerID int64, fileName string, fileReader io.Reader) (string, error) {
	//TODO implement me
	panic("implement me")
}

//
//func (r *S3ChoreoContentRepository) getChoreoUUID(id int64) string {
//	choreoUUID := md5.Sum([]byte(strconv.FormatInt(id, 10)))
//	choreoPath := fmt.Sprintf("choreo-%x/", choreoUUID)
//	path := fmt.Sprintf("%s/%s", r.s3Config.ContentConfig.Path, choreoPath)
//
//	return path
//}
//
//func (r *S3ChoreoContentRepository) upload(path string, fileReader io.Reader) (string, error) {
//	err := r.awsClient.S3Module.UploadFile(r.s3Config.BucketName, path, fileReader)
//	if err != nil {
//		return "", err
//	}
//	return r.awsClient.GetCloudFrontDistributionURL(path), nil
//}
//
//func (r *S3ChoreoContentRepository) UploadChoreoDetailContent(choreoDetailID int64, fileName string, fileReader io.Reader) (string, error) {
//	path := r.getChoreoUUID(choreoDetailID) + fmt.Sprintf("detailcontent-%s", fileName)
//	return r.upload(path, fileReader)
//}
//
//func (r *S3ChoreoContentRepository) UploadChoreoContent(choreoID int64, fileName string, fileReader io.Reader) (string, error) {
//	path := r.getChoreoUUID(choreoID) + fmt.Sprintf("content-%s", fileName)
//	return r.upload(path, fileReader)
//}
