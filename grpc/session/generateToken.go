package session

import (
	"github.com/tsingson/goums/pkg/vtils"
)

// GenerateToken  generate access token
// TODO:  rewrite
func GenerateToken() string {
	return vtils.RandString(32)
}

// GenerateExpiration  generate session expiration time
// TODO:  rewrite
func GenerateExpiration() string {
	return vtils.RandString(8)
}
