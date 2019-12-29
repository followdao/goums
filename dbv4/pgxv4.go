package dbv4

import (
	"context"
	"math"
	"time"

	"emperror.dev/errors"

	"github.com/RussellLuo/timingwheel"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	"github.com/tsingson/goums/dbv4/postgresconfig"
	"github.com/tsingson/logger"
)

// PrepareConfig prepare configuration from vkconfig
func prepareConfig(cfg *postgresconfig.PostgresConfig) *pgx.ConnConfig {
	config, _ := pgx.ParseConfig("")
	config.Host = cfg.Host
	config.Port = cfg.Port
	config.User = cfg.User
	config.Password = cfg.Password
	config.Database = cfg.Database
	config.LogLevel = pgx.LogLevelDebug

	// if cfg.Debug {
	config.LogLevel = pgx.LogLevelDebug
	// } else {
	// 	config.LogLevel = pgx.LogLevelError
	// }
	//
	if len(cfg.RuntimeParams) > 0 {
		config.RuntimeParams = cfg.RuntimeParams
	} else {
		config.RuntimeParams = map[string]string{"application_name": "vkst"}
	}
	return config
}

// Connect connect to database in pgx v4
func Connect(ctx context.Context, pCfg *postgresconfig.PostgresConfig,
	afterConnect func(context.Context, *pgx.Conn) error, log *logger.ZapLogger) (conn *pgx.Conn, err error) {
	cfg := prepareConfig(pCfg)
	cfg.Logger = log.PgxLogger()
	cfg.LogLevel = pgx.LogLevelError

	conn, err = pgx.ConnectConfig(ctx, cfg)

	if err != nil {
		log.Error("Unable to create connection", zap.Error(err))
		return nil, err
	}

	if afterConnect != nil {
		err = afterConnect(ctx, conn)
		if err != nil {
			_ = conn.Close(ctx)
			log.Error("数据库插入或删除TVOD SQL语句预编译出错", zap.Error(err))
			return nil, err
		}
	}

	// conn.ConnInfo().RegisterDataType(pgtype.DataType{
	// 	Value: &pgxuuid.UUID{},
	// 	Name:  "uuid",
	// 	OID:   2950,
	// })

	return conn, nil
}

// ConnectString  connect to database in pgx v4
func ConnectString(ctx context.Context, pCfg string,
	afterConnect func(context.Context, *pgx.Conn) error, log *logger.ZapLogger) (conn *pgx.Conn, err error) {
	var cfg *pgx.ConnConfig
	cfg, err = pgx.ParseConfig(pCfg)
	if err != nil {
		return nil, err
	}
	cfg.Logger = log.PgxLogger()
	cfg.LogLevel = pgx.LogLevelError

	conn, err = pgx.ConnectConfig(ctx, cfg)

	if err != nil {
		log.Error("Unable to create connection", zap.Error(err))
		return nil, err
	}

	if afterConnect != nil {
		err = afterConnect(ctx, conn)
		if err != nil {
			_ = conn.Close(ctx)
			log.Error("数据库插入或删除TVOD SQL语句预编译出错", zap.Error(err))
			return nil, err
		}
	}

	// conn.ConnInfo().RegisterDataType(pgtype.DataType{
	// 	Value: &pgxuuid.UUID{},
	// 	Name:  "uuid",
	// 	OID:   2950,
	// })

	return conn, nil
}

// ConnectLoop connect with time out
func connectLoop(ctx context.Context, cfg *postgresconfig.PostgresConfig,
	afterConnectMap func(ctx2 context.Context, conn *pgx.Conn) error, log *logger.ZapLogger,
	timeout time.Duration) (conn *pgx.Conn, err error) {
	// log := l.Named("ConnectLoop")

	tw := timingwheel.NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	count := int(math.RoundToEven(timeout.Seconds()))
	running := make(chan time.Time, 1)

	for i := 0; i < count; i++ {
		tw.AfterFunc(time.Duration(i)*time.Second, func() {
			running <- time.Now()
		})
	}

	timeoutExceeded := make(chan time.Time, 1)
	tw.AfterFunc(timeout, func() {
		timeoutExceeded <- time.Now()
	})

	for {
		select {
		case <-timeoutExceeded:
			err = errors.Errorf("db connection failed after %s timeout", timeout)
			return nil, err

		case <-running:
			conn, err = Connect(ctx, cfg, afterConnectMap, log)
			if err == nil {
				log.Info("connect to DB success")
				return
			}
			//
			log.Error("failed to connect DB", zap.Error(err))
		}
	}
}

// ConnectPool connect pool to postgres in pgx v4
func ConnectPool(ctx context.Context, pCfg *postgresconfig.PostgresConfig,
	afterConnect func(context.Context, *pgx.Conn) error, log *logger.ZapLogger) (pool *pgxpool.Pool, err error) {
	pgxLogger := log.PgxLogger()
	cfg := prepareConfig(pCfg)
	// cfg.Logger = pgxLogger

	var poolConfig *pgxpool.Config

	if len(pCfg.PoolConnections) > 0 {
		poolConfig, _ = pgxpool.ParseConfig("pool_max_conns=" + pCfg.PoolConnections)
	} else {
		poolConfig, _ = pgxpool.ParseConfig("pool_max_conns=42")
	}
	// poolConfig := &pgxpool.VkConfig{
	// 	ConnConfig:        cfg,
	// 	MaxConns:          4,
	// 	HealthCheckPeriod: 5 * time.Minute,
	// 	MaxConnLifetime:   24 * 31 * 6 * time.Hour,
	// }

	poolConfig.ConnConfig = cfg
	poolConfig.HealthCheckPeriod = 500 * time.Millisecond
	// poolConfig.MaxConnLifetime = 100 * time.Millisecond
	poolConfig.MaxConnLifetime = 24 * 31 * time.Hour
	poolConfig.AfterConnect = afterConnect
	poolConfig.ConnConfig.Logger = pgxLogger
	poolConfig.MaxConns = 4
	// poolConfig.BeforeAcquire
	poolConfig.ConnConfig.LogLevel = pgx.LogLevelError

	pool, err = pgxpool.ConnectConfig(ctx, poolConfig)

	return
}
