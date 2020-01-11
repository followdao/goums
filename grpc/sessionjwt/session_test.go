package sessionjwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	as := assert.New(t)

	s := NewSessionStore()
	userID := "111111"
	secret := "1"
	role := "access"
	u, err := NewSession(userID, role, secret, 2)
	as.NoError(err)
	s.Insert(u)

	u, jb, er1 := s.Login(userID)
	as.NoError(er1)

	u, jb, er1 = s.Login(userID)
	as.NoError(er1)

	// u , jb, err = s.Login(userID)
	// as.NoError(err)
	if er1 == nil {
		// time.Sleep( 2 *time.Second)
		jw, err := ParseVerifyToken(jb)
		as.NoError(err)
		err = s.Valid(jw)
		as.NoError(err)
		if err == nil {
			fmt.Println("--------------")
		}
	}
	// u.LogoutAllSessions()
	s.Update(u)
	_, jb1, er2 := s.Login(userID)
	as.NoError(er2)

	_, jb1, er2 = s.Login(userID)
	as.NoError(er2)
	// litter.Dump(u1 )
	// litter.Dump(user)
	if er2 == nil {
		time.Sleep(2 * time.Second)
		_, err := ParseVerifyToken(jb1)
		as.Error(err)
		// err = s.Valid(jw)
		// as.Error(err)
	}
	jw2, er3 := ParseVerifyToken(jb)
	as.Error(er3)
	err = s.Valid(jw2)
	as.Error(err)
}
