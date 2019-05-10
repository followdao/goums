package services

import (
	"os"
	"strconv"
	"strings"
	"testing"

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
		email    string
		password string
	}{
		{"t@1.1", "123456"},
		{"email@1.2", "123452"},
	}

	for _, test := range tests {
		ac, er := as.Register(test.email, test.password)
		assertions.NoError(er)
		assertions.Equal(ac.Email, test.email)
	}
}

func BenchmarkAccountStore_Register(b *testing.B) {
	var test = struct {
		email    string
		password string
	}{"t@1.1", "123456"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = as.Register(test.email, test.password)
	}
}

func TestAccountStore_Register_Error(t *testing.T) {
	var test = struct {
		email    string
		password string
	}{"t@1.1", "123456"}

	_, err := as.Register(test.email, test.password)
	assert.Equal(t, err, ErrDuplicated)

}

func TestAccountStore_Login(t *testing.T) {
	assertions := assert.New(t)

	var tests = []struct {
		email    string
		password string
	}{
		{"t@1.1", "123456"},
		{"email@1.2", "123452"},
	}

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

		_, _ = as.Register(strBuilder(test.email, strconv.Itoa(i)), strBuilder(test.password, strconv.Itoa(i)))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		_, _ = as.Login(strBuilder(test.email, strconv.Itoa(i)), strBuilder(test.password, strconv.Itoa(i)))
	}
}

func strBuilder(args ...string) string {
	var str strings.Builder

	for _, v := range args {
		str.WriteString(v)
	}
	return str.String()
}
