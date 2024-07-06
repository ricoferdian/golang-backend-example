package handler

import (
	"context"
	"errors"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
	"time"
)

var (
	TypeChoreo        = 1
	TypeChoreoDetail  = 2
	TypeChoreographer = 3

	mapPrimaryKey = map[int]string{
		TypeChoreo:        "choreo_id",
		TypeChoreoDetail:  "choreo_detail_id",
		TypeChoreographer: "choreographer_id",
	}
	mapFileCategory = map[int][]int{
		TypeChoreo:        {constants.FileCategoryThumbnailImage, constants.FileCategoryVideo},
		TypeChoreoDetail:  {constants.FileCategoryThumbnailImage, constants.FileCategoryVideo, constants.FileCategoryTestVideo},
		TypeChoreographer: {constants.FileCategoryThumbnailImage},
	}
)

func getFile(c *gin.Context) (io.Reader, string, error) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, "", err
	}
	filename := header.Filename

	return file, filename, err
}

func validateFileCategory(key int, fileCategory int) error {
	for _, category := range mapFileCategory[key] {
		if category == fileCategory {
			return nil
		}
	}
	return errors.New("invalid file category")
}

func GenericFileUploadHandler(c *gin.Context, key int, timeout int, handle func(ctx context.Context, choreoID int64, fileCategory int, fileName string, fileReader io.Reader) (interface{}, error)) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(timeout))
	defer cancel()

	startTime := time.Now()

	file, filename, err := getFile(c)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	primaryID, err := strconv.Atoi(c.Request.URL.Query().Get(mapPrimaryKey[key]))
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	fileCategory, err := strconv.Atoi(c.Request.URL.Query().Get("file_category"))
	if primaryID == 0 || fileCategory == 0 || err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	err = validateFileCategory(key, fileCategory)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	data, err := handle(ctx, int64(primaryID), fileCategory, filename, file)
	if err != nil {
		http.WriteErrorResponseObj(c, startTime, http.StatusServerError, http.ErrorResponse{
			Code:       http.StatusServerError,
			ErrMessage: err.Error(),
			ErrReason:  err.Error(),
		})
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}

func GenericDeleteHandler(c *gin.Context, timeout int, key string, handle func(ctx context.Context, id int64) error) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(timeout))
	defer cancel()

	startTime := time.Now()

	id, err := strconv.Atoi(c.Request.URL.Query().Get(key))
	if id == 0 || err != nil {
		log.Println("[DeleteHandler] error parse query param primary key", err)
		http.WriteErrorResponseObj(c, startTime, http.StatusInvalidRequest, http.ErrorResponse{
			Code:       http.StatusInvalidRequest,
			ErrMessage: err.Error(),
			ErrReason:  err.Error(),
		})
	}
	err = handle(ctx, int64(id))
	if err != nil {
		http.WriteErrorResponseObj(c, startTime, http.StatusServerError, http.ErrorResponse{
			Code:       http.StatusServerError,
			ErrMessage: err.Error(),
			ErrReason:  err.Error(),
		})
		return
	}
	http.WriteSuccessResponse(c, startTime, nil)
	return
}
