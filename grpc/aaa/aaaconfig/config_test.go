package aaaconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tsingson/goums/grpc/config"
	"github.com/tsingson/goums/pkg/vtils"
)

func TestDefault(t *testing.T) {
	c := Default()
	assert.Equal(t, c.AaaConfig.ServerPort, config.AAAPort)
	_ = vtils.SaveTOML(c, "testdata/"+ConfigFilename)
}

func TestLoad(t *testing.T) {
	c, err := Load("testdata/" + ConfigFilename)
	assert.NoError(t, err)
	if err == nil {
		assert.Equal(t, c.LiftConfig.AddressList[0], "127.0.0.1:9292")
	}
}
