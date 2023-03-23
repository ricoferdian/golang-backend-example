package jwtauth

import (
	"kora-backend/internal/entity"
)

func userEntityToUserJwtPayload(user entity.UserEntity) JWTUserPayload {
	return JWTUserPayload{
		UserID:       user.UserID,
		UserIdentity: user.UserIdentity,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserType:     user.UserType,
	}
}

func jwtPayloadToUserEntity(user JWTUserPayload) entity.AuthenticatedUserEntity {
	return entity.AuthenticatedUserEntity{
		UserID:       user.UserID,
		UserIdentity: user.UserIdentity,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserType:     user.UserType,
	}
}
