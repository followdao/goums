package terminaldbo

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"

	"github.com/tsingson/logger"

	"github.com/tsingson/goums/dbv4/postgresconfig"
	"github.com/tsingson/goums/pkg/vtils"
)

var (
	cfg *postgresconfig.PostgresConfig
	log *logger.ZapLogger
)

// var conn *pgx.conn

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	log = logger.New()

	cfg = &postgresconfig.PostgresConfig{
		User:     "postgres",
		Database: "goums",
		Port:     5432,
		Host:     "127.0.0.1",
		// Host: "docker.for.mac.host.modelset",
		LogLevel: pgx.LogLevelDebug,
	}

	log = logger.New(logger.WithDebug(), logger.WithStoreInDay())

	os.Exit(m.Run())
}

func TestNewTerminalDbo(t *testing.T) {
	ctx := context.Background()
	terminalDbo, err := NewTerminalDbo(ctx, cfg, log)
	assert.NoError(t, err)
	terminalDbo.Close(ctx)
}

func TestTerminalDbo_InsertTerminal(t *testing.T) {
	ctx := context.Background()
	terminalDbo, err := NewTerminalDbo(ctx, cfg, log)
	assert.NoError(t, err)

	sn := vtils.RandString(16)
	co := vtils.RandString(16)

	var id int64
	id, err = terminalDbo.InsertTerminal(ctx, sn, co)
	assert.NoError(t, err)
	if err == nil {
		fmt.Println("id ", id)
	}
}

func TestTerminalDbo_Notify(t *testing.T) {
	ctx := context.Background()
	v, err := NewTerminalDbo(ctx, cfg, log)
	assert.NoError(t, err)

	assert.NoError(t, err)

	v.Notify(ctx)
}
