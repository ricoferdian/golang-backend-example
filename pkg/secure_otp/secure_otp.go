package secure_otp

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"io"
	"log"
	"os"
	"time"
)

const (
	maxOtpDigits      = 4
	otpCacheKeyPrefix = "kora:authentication:secure_otp:%s"
)

// getSecretEnv used to get jwt secret key from environment variable.
func getSecretEnv() (string, error) {
	env := os.Getenv("KORA_OTP_SECRET_KEY")
	if env == "" {
		return "", errors.New("Unable to get OTP secret key environment variable")
	}
	return env, nil
}

func (m *SecureOtpModule) getKey(id string) string {
	return fmt.Sprintf(otpCacheKeyPrefix, id)
}

func (m *SecureOtpModule) setOtpCache(ctx context.Context, identity string, data SecureOtpCacheData) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return m.redisCli.Set(ctx, m.getKey(identity), body, 2*time.Minute).Err()
}

func (m *SecureOtpModule) getOtpCache(ctx context.Context, identity string) (SecureOtpCacheData, error) {
	var data SecureOtpCacheData
	val, err := m.redisCli.Get(ctx, m.getKey(identity)).Result()
	if val == "" {
		return data, nil
	}
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (m *SecureOtpModule) delOtpCache(ctx context.Context, identity string) error {
	return m.redisCli.Del(ctx, m.getKey(identity)).Err()
}

func (m *SecureOtpModule) generateCode() string {
	b := make([]byte, maxOtpDigits)
	n, err := io.ReadAtLeast(rand.Reader, b, maxOtpDigits)
	if n != maxOtpDigits {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (m *SecureOtpModule) isValid(secret, passcode string) bool {
	return secret == passcode
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func (m *SecureOtpModule) GenerateOtp(ctx context.Context, request entity.SecureOtpRequest) (string, error) {
	otpCode := m.generateCode()

	data := SecureOtpCacheData{
		Secret:   otpCode,
		Identity: request.ReceiverIdentity,
	}
	err := m.setOtpCache(ctx, request.ReceiverIdentity, data)
	if err != nil {
		return "", err
	}
	return otpCode, nil
}

func (m *SecureOtpModule) ValidateOtp(ctx context.Context, request entity.SecureOtpRequest, passcode string) (string, error) {
	cacheData, err := m.getOtpCache(ctx, request.ReceiverIdentity)
	if err != nil {
		return "", err
	}
	if cacheData.Secret == "" {
		return "", nil
	}

	if m.isValid(cacheData.Secret, passcode) {
		err := m.delOtpCache(ctx, request.ReceiverIdentity)
		if err != nil {
			log.Println("err delete otp cache due to ", err)
		}
		return cacheData.Identity, nil
	}

	return "", nil
}
