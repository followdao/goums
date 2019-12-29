package postgresconfig

import (
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewPostgresConfig(t *testing.T) {
	as := assert.New(t)

	dbHost := "10.0.0.13"
	cfg := NewPostgresConfig(WithHost(dbHost), WithLogLevel(pgx.LogLevelDebug))
	as.Equal(cfg.Host, dbHost)
	//	vtils.SaveTOML(p, "testdata/postgres-config.toml")
}

func TestLoad(t *testing.T) {
	as := assert.New(t)

	dbHost := "10.0.0.13"
	cfg, err := Load("testdata/postgres-config.toml")
	as.NoError(err)
	as.Equal(cfg.Host, dbHost)
}
