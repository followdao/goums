package notidy

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tsingson/logger"

	"github.com/tsingson/goums/dbv4"
	"github.com/tsingson/goums/dbv4/postgresconfig"
)

// TerminalDbo  cmsdbo dbo
type NotifyDbo struct {
	// ctx   context.Context
	// conn  *pgx.Conn
	pool  *pgxpool.Pool
	log   *logger.ZapLogger
	debug bool
}

// NewNotifyDbo connect db for postgres notify
func NewNotifyDbo(ctx context.Context, cfg *postgresconfig.PostgresConfig, log *logger.ZapLogger) (*NotifyDbo, error) { // 数据库配置

	cfg.SetLogger(log)

	pool, er1 := dbv4.ConnectPool(ctx, cfg, afterConnect, log)
	if er1 != nil {
		return nil, er1
	}
	cmsDb := &NotifyDbo{
		pool:  pool,
		log:   log,
		debug: true,
	}
	return cmsDb, nil
}

// Close  close pgx connection
func (s *NotifyDbo) Close(ctx context.Context) error {
	return nil // s.conn .Close(ctx)
}

// Acquire  get connect
func (s *NotifyDbo) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	return s.pool.Acquire(ctx)
}
