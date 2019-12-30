package proxyconfig

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestNewProxyConfig(t *testing.T) {
	p := NewProxyConfig()
	serverPort := ":80"
	assert.Equal(t, p.ServerPort, serverPort)
}
