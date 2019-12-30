package terminaldbo

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"

	"github.com/tsingson/logger"

	"github.com/tsingson/goums/apis/flatums"
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

func TestTerminalDbo_UpdateTerminal(t *testing.T) {
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

	var c int64
	c, err = terminalDbo.UpdateTerminal(ctx, id, true, 2, 2)

	assert.NoError(t, err)
	if err == nil {
		fmt.Println(c)
	}
}

func TestTerminalDbo_Active(t *testing.T) {
	ctx := context.Background()
	terminalDbo, err := NewTerminalDbo(ctx, cfg, log)
	assert.NoError(t, err)

	sn := vtils.RandString(16)
	co := vtils.RandString(16)

	var userID int64
	userID, err = terminalDbo.InsertTerminal(ctx, sn, co)
	assert.NoError(t, err)

	apkType := "test"

	var id *flatums.TerminalProfileT
	id, err = terminalDbo.Active(ctx, sn, co, apkType)
	assert.NoError(t, err)
	assert.Equal(t, id.UserID, userID)
}

func TestTerminalDbo_Notify(t *testing.T) {
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

	terminalDbo.Notify(ctx)

	var c int64
	c, err = terminalDbo.UpdateTerminal(ctx, id, true, 2, 2)

	assert.NoError(t, err)
	if err == nil {
		fmt.Println(c)
	}

	time.Sleep(5 * time.Second)
}

func TestTerminalDbo_UmsNotify(t *testing.T) {
	ctx := context.Background()
	terminalDbo, err := NewTerminalDbo(ctx, cfg, log)
	assert.NoError(t, err)
	terminalDbo.UmsNotify(ctx)
}
