package general

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func GenerateUniqueFileName() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func GetFileFormat(fileName string) (string, error) {
	format := strings.Split(fileName, ".")
	if len(format) < 2 {
		return "", errors.New("Invalid file format")
	}
	return format[len(format)-1], nil
}
