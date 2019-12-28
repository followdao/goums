package services

import (
	"os"
	"strconv"
	"testing"

	"github.com/tsingson/go-ums/pkg/utils"

	"github.com/stretchr/testify/assert"
)

var (
	as *AccountStore
)

func TestMain(m *testing.M) {
	as = New()
	os.Exit(m.Run())
}

func TestAccountStore_Register(t *testing.T) {
	assertions := assert.New(t)

	var tests = []struct {
		email    []byte
		password []byte
	}{
		{[]byte("t@1.1"), []byte("123456")},
	}

	for _, test := range tests {
		ac, er := as.Register(test.email, test.password)
		assertions.NoError(er)
		assertions.Equal(ac.Email, test.email)
	}
}

func BenchmarkAccountStore_Register(b *testing.B) {
	var test = struct {
		email    []byte
		password []byte
	}{[]byte("t@1.1"), []byte("123456")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = as.Register(utils.ByteBuilder(test.email, []byte(strconv.Itoa(i))), utils.ByteBuilder(test.password, []byte(strconv.Itoa(i))))
	}
}

func TestAccountStore_Register_Error(t *testing.T) {
	var test = struct {
		email    []byte
		password []byte
	}{[]byte("t@1.1"), []byte("123456")}

	_, err := as.Register(test.email, test.password)
	assert.Equal(t, err, ErrDuplicated)

}

/**
func TestAccountStore_Login(t *testing.T) {
	assertions := assert.New(t)

	var tests = []struct {
		email    string
		password string
	}{
		{"t@1.1", "123456"},
		{"email@1.2", "123452"},
	}

	// for _, test := range tests {
	//
	// 	_, _ = as.Register(test.email, test.password)
	//
	// }
	for _, test := range tests {
		_, er := as.Login(test.email, test.password)
		assertions.NoError(er)
	}
}

func BenchmarkAccountStore_Login(b *testing.B) {

	var test = struct {
		email    string
		password string
	}{"t@1.1", "123456"}

	for i := 0; i < b.N; i++ {

		_, _ = as.Register(StrBuilder(test.email, strconv.Itoa(i)), StrBuilder(test.password, strconv.Itoa(i)))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		_, _ = as.Login(StrBuilder(test.email, strconv.Itoa(i)), StrBuilder(test.password, strconv.Itoa(i)))
	}
}


*/
