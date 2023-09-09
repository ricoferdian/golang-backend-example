package usecase

import (
	"github.com/Kora-Dance/koradance-backend/internal/helper"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
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
