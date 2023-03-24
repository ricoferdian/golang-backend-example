package jwtauth

import (
	"kora-backend/app/helper"
)

type JwtAuthModule struct {
	config    *helper.JWTConfig
	secretKey string
}

type JWTUserAuthPayload struct {
	ClaimAudience string
	ClaimIssuer   string
	TokenExpiry   int64
	UserData      JWTUserPayload
}

type JWTUserPayload struct {
	UserID       int64  `json:"user_id"`
	UserIdentity string `json:"user_identity"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserType     int16  `json:"user_type"`
}

func NewJwtAuthModule(config *helper.JWTConfig) (*JwtAuthModule, error) {
	secretKey, err := getSecretEnv()
	if err != nil {
		return nil, err
	}
	return &JwtAuthModule{
		config:    config,
		secretKey: secretKey,
	}, nil
}
