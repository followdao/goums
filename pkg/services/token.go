package services

import (
	"strconv"
	"time"
)

func GenerateToken() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
