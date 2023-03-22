package http

const (
	errForbiddenMsg    = "Maaf, anda tidak memiliki wewenang untuk mengakses fitur ini"
	errForbiddenReason = "User has not authenticated or token not provided"

	errInvalidMsg    = "Maaf, terjadi kesalahan"
	errInvalidReason = "Invalid request parameters"
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
