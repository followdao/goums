package services

import (
	"github.com/tsingson/go-ums/model"
)

// SessionStore
type SessionStore struct {
	sessionList map[string]*session
}

type session struct {
	AccountID        int64
	Role             model.Role
	AccessToken      string
	metadata         map[string]*[]byte
	accessTimestamp  int64
	expiredTimestamp int64
	refreshTimestamp int64
}
