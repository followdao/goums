package services

import (
	"strings"
	"sync"
	"time"

	"github.com/orcaman/concurrent-map"
	"golang.org/x/xerrors"

	"github.com/tsingson/go-ums/model"
)

type AccountStore struct {
	lock  sync.Mutex
	total int
	// accountList     cmap.ConcurrentMap
	accountListMail cmap.ConcurrentMap
}

// New initial a AccountStore
func New() *AccountStore {
	// accountList := cmap.New()
	accountListMail := cmap.New()
	return &AccountStore{
		// accountList:     accountList,
		accountListMail: accountListMail,
	}
}

// Register account register
func (as *AccountStore) Register(email, password string) (account *model.Account, err error) {

	// check parameter
	{
		// TODO: verify parameter
	}

	// check email is duplicated or noet
	var ok bool
	ok = as.Exists(email)
	if ok {
		return nil, ErrDuplicated
	}
	// add to store
	uid := GenerateID()
	account = &model.Account{
		ID:        uid,
		Email:     email,
		Password:  password,
		Role:      model.Role_Guest,
		Status:    model.StatusWaitForEmailVerify,
		CreatedAt: timeNow(),
		UpdatedAt: timeNow(),
	}
	// update account profile
	{
		as.UpsetEmail(account)
	}
	account.Password = ""
	return
}

func (as *AccountStore) UpsetEmail(ac *model.Account) {
	cb := func(exists bool, valueInMap interface{}, newValue interface{}) interface{} {
		nv := newValue.(*model.Account)
		if !exists {
			return []*model.Account{nv}
		}
		res := valueInMap.([]*model.Account)
		return append(res, nv)
	}
	as.accountListMail.Upsert(ac.Email, ac, cb)

}

// Exists check account is duplicated or not
func (as *AccountStore) Exists(email string) bool {
	// check parameter
	{

	}
	//
	return as.accountListMail.Has(email)

}

// Login account login
func (as *AccountStore) Login(email, password string) (token string, err error) {
	// check parameter
	{

	}

	// check email is duplicated or noet
	ok := as.Exists(email)
	if ok {
		err = ErrAccountMissing
		return "", err
	}

	// generate token and update account
	v, _ := as.accountListMail.Get(email)
	ac := v.(*model.Account)
	if strings.EqualFold(ac.Password, password) {
		token = GenerateToken()
		ac.AccessToken = token
		ac.UpdatedAt = timeNow()
		// update account
		{
			as.UpsetEmail(ac)
		}
		err = nil
		return token, err
	}
	err = xerrors.New("wrong password")
	return
}

// Logout  account logout
func (as *AccountStore) Logout(token string) (err error) {
	err = nil
	return
}

// Auth account auth
func (as *AccountStore) AuthToken(token string) (pass bool, err error) {
	return true, nil
}

// Verify account's access token versify
func (as *AccountStore) Verify(token string) (pass bool, err error) {
	return true, nil
}

func timeNow() int64 {
	return time.Now().UnixNano()
}
