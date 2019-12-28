package services

import (
	"bytes"
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/tsingson/goutils"

	"github.com/orcaman/concurrent-map"
	"golang.org/x/xerrors"

	"github.com/tsingson/go-ums/model"
)

// AccountStore   storage of account
type AccountStore struct {
	total           int
	accountListMail cmap.ConcurrentMap
	accountCheck    *fastcache.Cache
}

// New initial a AccountStore
func New() *AccountStore {
	accountListMail := cmap.New()
	ch := fastcache.New(1024 * 16)
	return &AccountStore{
		// accountList:     accountList,
		accountListMail: accountListMail,
		accountCheck:    ch,
	}
}

// Register account register
func (as *AccountStore) Register(email, password []byte) (account *model.AccountProfile, err error) {

	// check parameter
	{
		// TODO: verify parameter
	}

	// check email is duplicated or noet

	ok := as.Exists(email)
	if ok {
		return nil, ErrDuplicated
	}
	// add to store
	uid := GenerateID()
	ac := &model.Account{
		ID:        uid,
		Email:     email,
		Role:      model.RoleGuest,
		Status:    model.StatusWaitForEmailVerify,
		CreatedAt: timeNow(),
		UpdatedAt: timeNow(),
	}
	account = &model.AccountProfile{
		ID:        uid,
		Email:     goutils.B2S(email),
		Role:      model.RoleGuest,
		Status:    model.StatusWaitForEmailVerify,
		CreatedAt: timeNow(),
		UpdatedAt: timeNow(),
	}
	// update account profile
	{
		as.UpsetEmail(ac)
		as.total++
	}

	return
}

// UpsetEmail  update or insert account profile via email
func (as *AccountStore) UpsetEmail(ac *model.Account) {
	cb := func(exists bool, valueInMap interface{}, newValue interface{}) interface{} {
		nv := newValue.(*model.Account)
		if !exists {
			return []*model.Account{nv}
		}
		res := valueInMap.([]*model.Account)
		return append(res, nv)
	}
	as.accountListMail.Upsert(goutils.B2S(ac.Email), ac, cb)
	as.accountCheck.Set(ac.Email, []byte("1"))

}

// Exists check account is duplicated or not
func (as *AccountStore) Exists(email []byte) bool {
	// check parameter
	{

	}
	//
	return as.accountCheck.Has(email)

}

// Login account login
func (as *AccountStore) Login(email, password []byte) (token []byte, err error) {
	// check parameter
	{

	}

	// check email is duplicated or noet
	ok := as.Exists(email)
	if ok {
		err = ErrAccountMissing
		return token, err
	}

	// generate token and update account
	v, _ := as.accountListMail.Get(goutils.B2S(email))
	ac := v.(*model.Account)
	if bytes.Equal(ac.Password, password) {
		token = GenerateToken()
		ac.AccessToken = token
		ac.UpdatedAt = timeNow()
		// update account
		{
			as.UpsetEmail(ac)
			as.total++
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

// AuthToken account auth
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
