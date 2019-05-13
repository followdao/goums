package services

import (
	"strconv"
	"time"
)

// GenerateToken generate token
func GenerateToken() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
