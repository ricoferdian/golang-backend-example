package secure_otp

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/redis/go-redis/v9"
)

type SecureOtpCacheData struct {
	Secret   string `json:"secret"`
	Identity string `json:"identity"`
}

type SecureOtpModule struct {
	config   *helper.SecureOtpConfig
	redisCli *redis.Client
}

func NewSecureOtpModule(config *helper.SecureOtpConfig, redisCli *redis.Client) *SecureOtpModule {
	return &SecureOtpModule{
		config:   config,
		redisCli: redisCli,
	}
}
