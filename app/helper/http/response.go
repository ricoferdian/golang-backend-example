package http

import (
	"github.com/gin-gonic/gin"
	"time"
)

type ErrorResponse struct {
	Code       string `json:"code"`
	ErrMessage string `json:"message"`
	ErrReason  string `json:"reason"`
}

type BaseResponse struct {
	Status      int            `json:"status_code"`
	ProcessTime float64        `json:"time"`
	Error       *ErrorResponse `json:"error,omitempty"`
	Data        interface{}    `json:"data"`
}

const (
	StatusSuccess             = "OK"
	StatusAuthFailed          = "AUTH_FAILED"
	StatusVerificationFailed  = "VERIFICATION_FAILED"
	StatusUserIdentifierExist = "USER_ALREADY_EXIST"
	StatusTokenExpired        = "EXPIRED_TOKEN"
	StatusInvalidRequest      = "INVALID_REQUEST"
	StatusFailedDatabase      = "ERROR_DATABASE"
	StatusServerError         = "SERVER_ERROR"
	StatusNotFound            = "NOT_FOUND"
	StatusUnauthorized        = "UNAUTHORIZED"
	StatusRequestTimeout      = "REQUEST_TIMEOUT"
)

var (
	mapStatusWithCode = map[string]int{
		StatusSuccess:            200,
		StatusAuthFailed:         200,
		StatusVerificationFailed: 200,
		StatusInvalidRequest:     400,
		StatusFailedDatabase:     500,
		StatusServerError:        500,
		StatusNotFound:           404,
		StatusUnauthorized:       401,
		StatusRequestTimeout:     408,
	}
)

func WriteResponse(c *gin.Context, pTime time.Time, httpStatus string, data interface{}, err *ErrorResponse) {
	httpCode := mapStatusWithCode[httpStatus]
	if httpCode == 0 {
		httpCode = 200
	}
	resp := BaseResponse{
		Status:      httpCode,
		ProcessTime: time.Since(pTime).Seconds(),
		Data:        data,
	}
	if err != nil {
		resp.Error = err
	}

	c.JSON(httpCode, resp)
	return
}

func WriteSuccessResponse(c *gin.Context, pTime time.Time, data interface{}) {
	WriteResponse(c, pTime, StatusSuccess, data, nil)
	return
}

func WriteErrorResponse(c *gin.Context, pTime time.Time, errStatus string, data interface{}, errMsg string, errReason string) {
	err := &ErrorResponse{
		Code:       errStatus,
		ErrMessage: errMsg,
		ErrReason:  errReason,
	}

	WriteResponse(c, pTime, errStatus, data, err)
	return
}

func WriteErrorResponseByCode(c *gin.Context, pTime time.Time, errStatus string) {
	err := GetErrResponse(errStatus)
	WriteResponse(c, pTime, errStatus, nil, &err)
	return
}

func WriteErrorResponseObj(c *gin.Context, pTime time.Time, errStatus string, errResp ErrorResponse) {
	WriteResponse(c, pTime, errStatus, nil, &errResp)
	return
}

func WriteEmptyErrorResponse(c *gin.Context, pTime time.Time, errStatus string) {
	WriteResponse(c, pTime, errStatus, nil, nil)
	return
}
