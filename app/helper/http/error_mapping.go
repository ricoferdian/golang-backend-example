package http

const (
	errForbiddenMsg    = "Maaf, anda tidak memiliki wewenang untuk mengakses fitur ini"
	errForbiddenReason = "User has not authenticated or token not provided"

	errInvalidMsg    = "Maaf, terjadi kesalahan"
	errInvalidReason = "Invalid request parameters"

	errServerMsg    = "Maaf, terjadi kesalahan"
	errServerReason = "Internal server error, please contact our team about this issue"

	authFailedMsg    = "Kredensial tidak ditemukan"
	authFailedReason = "Credentials does not match or user not found"

	tokenExpiredMsg    = "Harap masuk ke akun anda kembali"
	tokenExpiredReason = "Token has been invalidated due to expiration"

	userAlreadyExistMsg    = "Nama pengguna sudah digunakan"
	userAlreadyExistReason = "Token has been invalidated due to expiration"
)

var errorMapping = map[string]ErrorResponse{
	StatusForbidden: {
		Code:       StatusForbidden,
		ErrMessage: errForbiddenMsg,
		ErrReason:  errForbiddenReason,
	},
	StatusInvalidRequest: {
		Code:       StatusInvalidRequest,
		ErrMessage: errInvalidMsg,
		ErrReason:  errInvalidReason,
	},
	StatusServerError: {
		Code:       StatusServerError,
		ErrMessage: errServerMsg,
		ErrReason:  errServerReason,
	},
	StatusAuthFailed: {
		Code:       StatusAuthFailed,
		ErrMessage: authFailedMsg,
		ErrReason:  authFailedReason,
	},
	StatusTokenExpired: {
		Code:       StatusTokenExpired,
		ErrMessage: tokenExpiredMsg,
		ErrReason:  tokenExpiredReason,
	},
	StatusUserIdentifierExist: {
		Code:       StatusUserIdentifierExist,
		ErrMessage: userAlreadyExistMsg,
		ErrReason:  userAlreadyExistReason,
	},
}

func GetErrResponse(statusCode string) ErrorResponse {
	if val, ok := errorMapping[statusCode]; ok {
		return val
	}

	return ErrorResponse{
		Code:       StatusNotFound,
		ErrMessage: "Terjadi kesalahan",
		ErrReason:  "Tidak diketahui",
	}
}

func GetErrResponseWithReason(statusCode string, errReason string) ErrorResponse {
	errResp := GetErrResponse(statusCode)
	errResp.ErrReason = errReason
	return errResp
}
