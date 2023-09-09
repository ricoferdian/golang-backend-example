package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/internal/common/handler"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
	"time"
)

var (
	TypeChoreo       = 1
	TypeChoreoDetail = 2

	mapPrimaryKey = map[int]string{
		TypeChoreo:       "choreo_id",
		TypeChoreoDetail: "choreo_detail_id",
	}
	mapFileCategory = map[int][]int{
		TypeChoreo:       {constants.FileCategoryThumbnailImage, constants.FileCategoryVideo},
		TypeChoreoDetail: {constants.FileCategoryThumbnailImage, constants.FileCategoryVideo, constants.FileCategoryTestVideo},
	}
)

func (api ChoreoHandler) getChoreoListHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	difficulty, err := strconv.Atoi(c.Request.URL.Query().Get("difficulty"))
	if err != nil {
		log.Println("[ChoreoHandler] error parse query param difficulty", err)
	}
	choreoID, err := strconv.Atoi(c.Request.URL.Query().Get("choreo_id"))
	if err != nil {
		log.Println("[ChoreoHandler] error parse query param choreo_id", err)
	}
	choreographerID, err := strconv.Atoi(c.Request.URL.Query().Get("choreographer_id"))
	if err != nil {
		log.Println("[ChoreoHandler] error parse query param choreographer_id", err)
	}
	price, err := strconv.Atoi(c.Request.URL.Query().Get("price"))
	if err != nil {
		log.Println("[ChoreoHandler] error parse query param price", err)
	}
	filter := entity2.ChoreoFilterEntity{Difficulty: difficulty, ChoreoID: int64(choreoID), Price: int64(price), ChoreographerID: int64(choreographerID)}

	data, err := api.getChoreoListWithOptionalAuth(c, ctx, filter)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}

func (api ChoreoHandler) getChoreoListWithOptionalAuth(c *gin.Context, ctx context.Context, filter entity2.ChoreoFilterEntity) ([]entity2.ChoreographyEntity, error) {

	authData, isOk := c.Value(constants.CtxAuthUserData).(*entity2.AuthenticatedUserEntity)
	if !isOk {
		return api.choreoUC.GetChoreoList(ctx, filter)
	}
	return api.choreoUC.GetChoreoListWithUserContent(ctx, authData.UserID, filter)
}

func (api ChoreoHandler) getChoreoByIDHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	choreoID, err := strconv.Atoi(c.Request.URL.Query().Get("choreo_id"))
	if err != nil {
		log.Println("[ChoreoHandler] error parse query param choreo_id", err)
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
	}

	data, err := api.getChoreoByIDWithOptionalAuth(c, ctx, int64(choreoID))
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}

func (api ChoreoHandler) getChoreoByIDWithOptionalAuth(c *gin.Context, ctx context.Context, choreoID int64) (entity2.ChoreographyEntity, error) {

	authData, isOk := c.Value(constants.CtxAuthUserData).(*entity2.AuthenticatedUserEntity)
	if !isOk {
		return api.choreoUC.GetChoreoByID(ctx, choreoID)
	}
	return api.choreoUC.GetChoreoByIDWithUserContent(ctx, authData.UserID, choreoID)
}

func (api ChoreoHandler) getFile(c *gin.Context) (io.Reader, string, error) {
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

func (api ChoreoHandler) choreoFileUploadHandler(c *gin.Context, key int, handle func(ctx context.Context, choreoID int64, fileCategory int, fileName string, fileReader io.Reader) (interface{}, error)) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	file, filename, err := api.getFile(c)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	choreoID, err := strconv.Atoi(c.Request.URL.Query().Get(mapPrimaryKey[key]))
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	fileCategory, err := strconv.Atoi(c.Request.URL.Query().Get("file_category"))
	if choreoID == 0 || fileCategory == 0 || err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	err = validateFileCategory(key, fileCategory)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	data, err := handle(ctx, int64(choreoID), fileCategory, filename, file)
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

func (api ChoreoHandler) uploadChoreoContent(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return api.choreoFileUploadHandler(c, TypeChoreo, api.choreoUC.UploadChoreoContent)
}

func (api ChoreoHandler) uploadChoreoDetailContent(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return api.choreoFileUploadHandler(c, TypeChoreoDetail, api.choreoUC.UploadChoreoDetailContent)
}

func (api ChoreoHandler) getChoreoDataReq(c *gin.Context) (choreoData entity2.ChoreographyEntity, err error) {
	// Get choreo json data from request body
	err = json.NewDecoder(c.Request.Body).Decode(&choreoData)
	if err != nil {
		return choreoData, err
	}
	if err != nil {
		return choreoData, err
	}
	if choreoData.AdditionalInfo == "" {
		choreoData.AdditionalInfo = "{}"
	}

	return choreoData, nil
}

func (api ChoreoHandler) getChoreoDetailDataReq(c *gin.Context) (choreoData entity2.ChoreographyDetailEntity, err error) {
	// Get choreo json data from request body
	err = json.NewDecoder(c.Request.Body).Decode(&choreoData)
	if err != nil {
		return choreoData, err
	}
	return choreoData, nil
}

func (api ChoreoHandler) upsertChoreoHandler(c *gin.Context, handle func(ctx context.Context, choreo entity2.ChoreographyEntity) (result entity2.ChoreographyEntity, err error)) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	req, err := api.getChoreoDataReq(c)
	if err != nil {
		metricsErr = err
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	choreo, err := handle(ctx, req)
	if err != nil {
		metricsErr = err
		http.WriteErrorResponseObj(c, startTime, http.StatusServerError, http.ErrorResponse{
			Code:       http.StatusServerError,
			ErrMessage: err.Error(),
			ErrReason:  err.Error(),
		})
		return
	}
	http.WriteSuccessResponse(c, startTime, choreo)
	return
}

func (api ChoreoHandler) upsertChoreoDetailHandler(c *gin.Context, handle func(ctx context.Context, detail entity2.ChoreographyDetailEntity) (result entity2.ChoreographyDetailEntity, err error)) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	req, err := api.getChoreoDetailDataReq(c)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	choreo, err := handle(ctx, req)
	if err != nil {
		http.WriteErrorResponseObj(c, startTime, http.StatusServerError, http.ErrorResponse{
			Code:       http.StatusServerError,
			ErrMessage: err.Error(),
			ErrReason:  err.Error(),
		})
		return
	}
	http.WriteSuccessResponse(c, startTime, choreo)
	return
}

func (api ChoreoHandler) insertChoreoHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return api.upsertChoreoHandler(c, api.choreoUC.InsertChoreo)
}

func (api ChoreoHandler) insertChoreoDetailHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return api.upsertChoreoDetailHandler(c, api.choreoUC.InsertChoreoDetail)
}

func (api ChoreoHandler) updateChoreoHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return api.upsertChoreoHandler(c, api.choreoUC.UpdateChoreo)
}

func (api ChoreoHandler) updateChoreoDetailHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return api.upsertChoreoDetailHandler(c, api.choreoUC.UpdateChoreoDetail)
}

func (api ChoreoHandler) deleteChoreoByID(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return handler.GenericDeleteHandler(c, api.handlerCfg.Timeout, "choreo_id", api.choreoUC.DeleteChoreoByID)
}

func (api ChoreoHandler) deleteChoreoDetailByID(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return handler.GenericDeleteHandler(c, api.handlerCfg.Timeout, "choreo_detail_id", api.choreoUC.DeleteChoreoDetailByID)
}
