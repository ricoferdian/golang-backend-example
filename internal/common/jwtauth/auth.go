package jwtauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"kora-backend/internal/common/constants"
	"kora-backend/internal/entity"
	"os"
	"strings"
	"time"
)

// getSecretEnv used to get jwt secret key from environment variable.
func getSecretEnv() (string, error) {
	env := os.Getenv("KORA_JWT_SECRET_KEY")
	if env == "" {
		return "", errors.New("Unable to get JWT secret key environment variable")
	}
	return env, nil
}

// GenerateToken used to generate jwt token.
func (jwt *JwtAuthModule) GenerateToken(header string, payload JWTUserAuthPayload) (string, error) {
	// create a new hash of type sha256. We pass the secret key to it
	h := hmac.New(sha256.New, []byte(jwt.secretKey))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))
	// We then Marshal the payload which is a map. This converts it to a string of JSON.
	payloadstr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error generating Token")
		return string(payloadstr), err
	}
	payload64 := base64.StdEncoding.EncodeToString(payloadstr)

	// Now add the encoded string.
	message := header64 + "." + payload64

	// We have the unsigned message ready.
	unsignedStr := header + string(payloadstr)

	// We write this to the SHA256 to hash it.
	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	//Finally we have the token
	tokenStr := message + "." + signature
	return tokenStr, nil
}

// ValidateToken helps in validating the token
func (jwt *JwtAuthModule) ValidateToken(token string) (bool, *entity.AuthenticatedUserEntity, error) {
	// JWT has 3 parts separated by '.'
	splitToken := strings.Split(token, ".")
	// if length is not 3, we know that the token is corrupt
	if len(splitToken) != 3 {
		return false, nil, nil
	}

	// decode the header and payload back to strings
	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, nil, err
	}
	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, nil, err
	}
	//again create the signature
	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(jwt.secretKey))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Println(signature)

	// if both the signature don’t match, this means token is wrong
	if signature != splitToken[2] {
		return false, nil, nil
	}
	// Check payload
	authPayload := JWTUserAuthPayload{}
	err = json.Unmarshal(payload, &authPayload)
	if err != nil {
		return false, nil, err
	}
	// Check token expiry
	currentTime := time.Now().Unix()
	if authPayload.TokenExpiry < currentTime {
		return false, nil, errors.New(constants.ErrTokenExpired)
	}

	userEntity := jwtPayloadToUserEntity(authPayload.UserData)
	// This means the token matches
	return true, &userEntity, nil
}

func (jwt *JwtAuthModule) GetSignedToken(user entity.UserEntity) (string, int64, error) {
	// we make a JWT Token here with signing method of ES256 and claims.
	// claims are attributes.
	// aud - audience
	// iss - issuer
	// exp - expiration of the Token
	expiry := time.Now().Add(time.Second * time.Duration(jwt.config.ClaimExpirySec)).Unix()
	payload := JWTUserAuthPayload{
		ClaimAudience: jwt.config.ClaimAudience,
		ClaimIssuer:   jwt.config.ClaimIssuer,
		TokenExpiry:   expiry,
		UserData:      userEntityToUserJwtPayload(user),
	}
	// here we provide the shared secret. It should be very complex.
	// Also, it should be passed as a System Environment variable

	tokenString, err := jwt.GenerateToken(jwt.config.HeaderKey, payload)
	if err != nil {
		return "", 0, err
	}
	return tokenString, expiry, nil
}
