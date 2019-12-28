package services

import (
	"time"

	"github.com/tsingson/go-ums/pkg/utils"
)

// GenerateToken generate token
func GenerateToken() []byte {
	return utils.Int64ToBytes(time.Now().UnixNano())
}
