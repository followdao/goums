package gslbconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := NewGslbConfig()
	assert.Equal(t, cfg.Version, VERSION)
	// vtils.SaveTOML(cfg, "testdata/"+ConfigFilename)
}

func TestLoad(t *testing.T) {
	host := "127.0.0.1"
	cfg, err := Load("testdata/" + ConfigFilename)
	assert.NoError(t, err)

	if err == nil {
		assert.Equal(t, cfg.UmsPostgresConfig.Host, host)
		assert.Equal(t, cfg.Version, VERSION)
		// cfg.Log.Info("---------")
	}
}
