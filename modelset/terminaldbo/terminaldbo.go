package terminaldbo

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/tsingson/logger"

	"github.com/tsingson/goums/dbv4"
	"github.com/tsingson/goums/dbv4/postgresconfig"
)

// TerminalDbo  cmsdbo dbo
type TerminalDbo struct {
	// ctx   context.Context
	// conn  *pgx.Conn
	pool  *pgxpool.Pool
	log   *logger.ZapLogger
	debug bool
}

// NewTerminalDbo connect cmsdbo db
func NewTerminalDbo(ctx context.Context, cfg *postgresconfig.PostgresConfig, log *logger.ZapLogger) (*TerminalDbo, error) { // 数据库配置
	// CMS
	// conn, err := dbv4.Connect(ctx, cfg, afterConnect, log)
	// if err != nil {
	// 	return nil, err
	// }

	cfg.SetLogger(log)

	pool, er1 := dbv4.ConnectPool(ctx, cfg, afterConnect, log)
	if er1 != nil {
		return nil, er1
	}
	cmsDb := &TerminalDbo{
		// conn:  conn,
		pool:  pool,
		log:   log,
		debug: true,
	}
	return cmsDb, nil
}

// Close  close pgx connection
func (s *TerminalDbo) Close(ctx context.Context) error {
	return nil // s.conn .Close(ctx)
}

// Acquire  get connect
func (s *TerminalDbo) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	return s.pool.Acquire(ctx)
}
