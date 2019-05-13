package services

import (
	"strconv"
	"time"
)

func GenerateID() int64 {
	return time.Now().UnixNano()
}

func GenerateIDString() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
