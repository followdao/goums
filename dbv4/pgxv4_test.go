package dbv4

import (
	"context"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap/zapcore"

	"github.com/tsingson/logger"

	"github.com/tsingson/goums/dbv4/postgresconfig"
)

var conn *pgx.Conn

func TestMain(m *testing.M) {
	runtime.MemProfileRate = 0
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)

	// logFullPath := "/Users/qinshen/go/bin/log/"

	log := logger.New()
	//
	logger.SetLevel(zapcore.DebugLevel)

	cfg := postgresconfig.NewPostgresConfig()

	ctx := context.Background()
	var err error
	conn, err = Connect(ctx, cfg, afterConnect, log)
	if err != nil {
		log.Log.Error("connect error", zap.Error(err))
		os.Exit(-1)
	}

	// setup(true)
	os.Exit(m.Run())
}

func TestConnect(t *testing.T) {
	// logFullPath := "/Users/qinshen/go/bin/log/"

	cfg := &postgresconfig.PostgresConfig{
		User:          "postgres",
		Password:      "postgres",
		Database:      "vktest",
		Port:          5432,
		Host:          "127.0.0.1",
		RuntimeParams: nil,
	}

	log := logger.New()
	//

	logger.SetLevel(zapcore.DebugLevel)
	ctx := context.Background()
	var err error
	conn, err = Connect(ctx, cfg, afterConnect, log)
	assert.NoError(t, err, "connect error  or not")
	assert.NotNil(t, conn, "connect success")
	conn.Close(ctx)
}

func TestConnectLoop(t *testing.T) {
	// logFullPath := "/Users/qinshen/go/bin/log/"

	cfg := postgresconfig.NewPostgresConfig(postgresconfig.WithDebug(true))

	log := logger.New()
	//

	logger.SetLevel(zapcore.DebugLevel)
	ctx := context.Background()
	var err error
	timeOut := 5 * time.Second
	conn, err = connectLoop(ctx, cfg, afterConnect, log, timeOut)
	assert.NoError(t, err, "connect error  or not")
	assert.NotNil(t, conn, "connect success")
	conn.Close(ctx)
}

func TestConnectPool(t *testing.T) {
	// logFullPath := "/Users/qinshen/go/bin/log/"

	cfg := postgresconfig.NewPostgresConfig(postgresconfig.WithDebug(true))
	log := logger.New()
	//

	logger.SetLevel(zapcore.DebugLevel)
	ctx := context.Background()
	var err error
	var pool *pgxpool.Pool
	pool, err = ConnectPool(ctx, cfg, afterConnect, log)
	assert.NoError(t, err, "connect error  or not")
	assert.NotNil(t, pool, "connect success")
	pool.Close()
}

const max = "42"

func BenchmarkConnectPool(b *testing.B) {
	cfg := postgresconfig.NewPostgresConfig(postgresconfig.WithDebug(true))
	cfg.PoolConnections = max

	aft := func(ctx context.Context, c *pgx.Conn) error {
		_, err := c.Prepare(ctx, "ps1", "select $1::int8")
		return err
	}
	log := logger.New(logger.WithDebug())
	logger.SetLevel(zapcore.DebugLevel)
	ctx := context.Background()
	pool, err := ConnectPool(ctx, cfg, aft, log)
	require.NoError(b, err)

	var n int64
	i := int64(100)

	b.SetParallelism(8)
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			err = pool.QueryRow(ctx, "ps1", i).Scan(&n)
			if err != nil {
				b.Fatal(err)
			}

			if n != i {
				b.Fatalf("expected %d, got %d", i, n)
			}
		}
	})
}

func BenchmarkConnectPool1(b *testing.B) {
	cfg := postgresconfig.NewPostgresConfig(postgresconfig.WithDebug(true))
	cfg.PoolConnections = max

	aft := func(ctx context.Context, c *pgx.Conn) error {
		_, err := c.Prepare(ctx, "ps1", "select $1::int8")
		return err
	}
	log := logger.New(logger.WithDebug())
	logger.SetLevel(zapcore.DebugLevel)
	ctx := context.Background()
	pool, err := ConnectPool(ctx, cfg, aft, log)
	require.NoError(b, err)

	var n int64

	b.SetParallelism(8)
	b.ReportAllocs()

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err = pool.QueryRow(ctx, "ps1", i).Scan(&n)
		if err != nil {
			b.Fatal(err)
		}

		if n != int64(i) {
			b.Fatalf("expected %d, got %d", i, n)
		}
	}
}

const (
	insertRawTVODSQL = "insertRawTVODSQL"
	removeRawTVODSQL = "removeRawTVOD"
	countRawTVODSQL  = "countRawTVODSQL"
)

func afterConnect(ctx context.Context, conn *pgx.Conn) (err error) {
	//
	_, err = conn.Prepare(ctx, insertRawTVODSQL,
		`INSERT INTO data.rawepg (channel_id, channel_title, title, subtitle, date, start_time, end_time, timezone ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) `)
	if err != nil {
		return
	}
	//
	_, err = conn.Prepare(ctx, removeRawTVODSQL, `delete from data.rawepg where  date=$1 and channel_id=$2`)
	if err != nil {
		return
	}

	_, err = conn.Prepare(ctx, countRawTVODSQL,
		`select count(id) as total from data.rawepg where channel_id=$1 and date=$2`)

	return nil
}
