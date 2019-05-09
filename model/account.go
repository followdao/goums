package model

import (
	"time"
)

// Account  account  define
type Account struct {
	ID          int64  // 全局唯一
	Email       string // email length >=5
	Password    string // password length >=6
	CreateAt    time.Time
	AccessToken string // accessToken length >=32
}

type OperationResult struct {
	Code int    // code = 200, success;  code=500, error
	Msg  string // return operation name when success,  return error message when error
}

// AccountOperation account operation interface define
type AccountOperation interface {
	Register(email, password string) (account *Account, result OperationResult)
	Exists(email string) (exists bool, result OperationResult)
	Login(email, password string) (token string, result OperationResult)
	Logout(token string) (result OperationResult)
	Auth(token string) (pass bool, result OperationResult)
	Verify(token string) (pass bool, result OperationResult)
}
