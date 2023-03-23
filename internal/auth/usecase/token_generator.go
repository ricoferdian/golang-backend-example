package usecase

import (
	"kora-backend/internal/auth/helper"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
	"strconv"
)

func (u UserAuthUseCaseImpl) generateTokenResponse(userData model.RbacUserModel) (*entity.AuthUserResponseEntity, error) {
	token, expiry, err := u.jwtModule.GetSignedToken(helper.UserModelToEntity(userData))
	if err != nil {
		return nil, err
	}
	response := entity.AuthUserResponseEntity{
		TokenData: entity.AuthTokenEntity{
			AccessToken: token,
			ExpiryTime:  strconv.FormatInt(expiry, 10),
		},
		UserData: helper.UserModelToEntity(userData),
	}
	return &response, nil
}
