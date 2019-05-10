package services

import (
	"strings"
	"sync"
	"time"

	"golang.org/x/xerrors"

	"github.com/tsingson/go-ums/model"
)

type AccountStore struct {
	lock            sync.Mutex
	total           int
	accountList     map[int64]*model.Account
	accountListMail map[string]*model.Account
}

// New initial a AccountStore
func New() *AccountStore {
	accountList := make(map[int64]*model.Account, 0)
	accountListMail := make(map[string]*model.Account, 0)
	return &AccountStore{
		accountList:     accountList,
		accountListMail: accountListMail,
	}
}

// Register account register
func (as *AccountStore) Register(email, password string) (account *model.Account, err error) {

	// check parameter
	{

	}

	// check email is duplicated or noet
	ok, _ := as.Exists(email)
	if ok {
		return nil, ErrDuplicated
	}
	// add to store
	uid := generateID()
	account = &model.Account{
		ID:        uid,
		Email:     email,
		Password:  password,
		Role:      model.Role_Guest,
		Status:    model.Status_WaitForEmailVerify,
		CreatedAt: timeNow(),
		UpdatedAt: timeNow(),
	}
	// update account profile
	{
		as.lock.Lock()
		as.accountList[uid] = account
		as.accountListMail[email] = account
		as.total++
		as.lock.Unlock()
	}
	return
}

// Exists check account is duplicated or not
func (as *AccountStore) Exists(email string) (exists bool, err error) {
	// check parameter
	{

	}
	//
	_, ok := as.accountListMail[email]
	if ok {
		return true, nil
	}
	return false, xerrors.New("account not exist via email")
}

// Login account login
func (as *AccountStore) Login(email, password string) (token string, err error) {
	// check parameter
	{

	}

	// check email is duplicated or noet
	ok, _ := as.Exists(email)
	if !ok {
		err = ErrAccountMissing
		return
	}

	// generate token and update account
	ac := as.accountListMail[email]
	if strings.EqualFold(ac.Password, password) {
		token = generateToken()
		ac.AccessToken = token
		ac.UpdatedAt = timeNow()
		// update account
		{
			as.lock.Lock()
			as.accountListMail[email] = ac
			as.accountList[ac.ID] = ac
			as.lock.Unlock()
		}

		return
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

func (as *AccountStore) AuthUID(id int64) (pass bool, err error) {
	// check parameter
	{

	}

	_, ok := as.accountList[id]
	if ok {
		return true, nil
	}
	return false, xerrors.New("account ID no exists")

}

// Verify account's access token versify
func (as *AccountStore) Verify(token string) (pass bool, err error) {
	return true, nil
}

func generateID() int64 {
	return time.Now().UnixNano()
}

func generateToken() string {
	return string(time.Now().UnixNano())
}

func timeNow() int64 {
	return time.Now().UnixNano()
}
