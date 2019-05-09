package services

import (
	"time"

	"github.com/tsingson/go-ums/model"
)

type AccountStore struct {
	accountList map[int64]*model.Account
}

// New initial a AccountStore
func New() *AccountStore {
	accountList := make(map[int64]*model.Account, 0)
	return &AccountStore{
		accountList: accountList,
	}
}

// Register account register
func (as *AccountStore) Register(email, password string) (account *model.Account, result model.OperationResult) {
	return
}

// Exists check account is duplicated or not
func (as *AccountStore) Exists(email string) (exists bool, result model.OperationResult) {
	return
}

// Login account login
func (as *AccountStore) Login(email, password string) (token string, result model.OperationResult) {
	return
}

// Logout  account logout
func (as *AccountStore) Logout(token string) (result model.OperationResult) {
	return

}

// Auth account auth
func (as *AccountStore) Auth(token string) (pass bool, result model.OperationResult) {
	return
}

// Verify account's access token versify
func (as *AccountStore) Verify(token string) (pass bool, result model.OperationResult) {
	return
}

func generateID() int64 {
	return time.Now().UnixNano()
}
