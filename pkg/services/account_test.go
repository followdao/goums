package services

import (
	"os"
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

func TestAccountStore_Login(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		email    string
		password string
	}{
		{"t@1.1", "123456"},
	}

	for _, test := range tests {
		ac, er := as.Register(test.email, test.password)
		assert.NoError(er)
		assert.Equal(ac.Email, test.email)
	}
}

func BenchmarkAccountStore_Login(b *testing.B) {
	var test = struct {
		email    string
		password string
	}{"t@1.1", "123456"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = as.Register(test.email, test.password)
	}
}
