package delivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kora-backend/app/helper/http"
	"kora-backend/internal/common/constants"
	"kora-backend/internal/entity"
	"strconv"
	"time"
)

func validateRegisterParam(c *gin.Context) (*entity.UserEntity, error) {
	userIdentity := c.PostForm("user_identity")
	passIdentifier := c.PostForm("password_identifier")
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	if userIdentity == "" || passIdentifier == "" || firstName == "" || lastName == "" {
		return nil, errors.New("user credentials are incomplete")
	}
	userType, err := strconv.Atoi(c.Request.Form.Get("user_type"))
	if err != nil {
		return nil, err
	}
	if userType == 0 {
		return nil, errors.New("user type not defined")
	}
	return &entity.UserEntity{
		UserIdentity:       userIdentity,
		PasswordIdentifier: passIdentifier,
		FirstName:          firstName,
		LastName:           lastName,
		UserType:           int16(userType),
	}, nil
}

func (api UserAuthHandler) authUserRegisterHandler(c *gin.Context) {
	startTime := time.Now()
	ctx := c.Request.Context()

	userData, err := validateRegisterParam(c)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	data, err := api.userAuthUC.Register(ctx, *userData)
	if err != nil || data == nil {
		errmsg := err.Error()
		if errmsg == constants.ErrSqlUserIdentityExist {
			http.WriteErrorResponseByCode(c, startTime, http.StatusUserIdentifierExist)
			return
		}
		http.WriteErrorResponseByCode(c, startTime, http.StatusServerError)
		return
	}
	http.WriteSuccessResponse(c, startTime, &data)
	return
}
