package postgresconfig

import (
	"github.com/tsingson/multiconfig"

	"github.com/imdario/mergo"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap/zapcore"

	"github.com/tsingson/logger"
)

// PostgresConfig  postgres configuration
type PostgresConfig struct {
	User            string            `json:"User"`
	Password        string            `json:"Password"`
	Database        string            `json:"Database"`
	Port            uint16            `json:"Port"`
	Host            string            `json:"Host"`
	Logger          pgx.Logger        `json:"-"`
	LogLevel        pgx.LogLevel      `json:"LogLevel"`
	RuntimeParams   map[string]string `json:"RuntimeParams"`
	PoolConnections string
	Debug           bool `json:"Debug"`
}

var defaultLog = logger.New(logger.WithDebug(),
	logger.WithDays(31),
	logger.WithLevel(zapcore.DebugLevel),
	logger.WithStoreInDay())

var defaultPostgresConfig = &PostgresConfig{
	User:            "postgres",
	Password:        "postgres",
	Database:        "vktest",
	Port:            5432,
	Host:            "127.0.0.1", //  "docker.for.mac.localhost",
	PoolConnections: "4",         // 42
	LogLevel:        pgx.LogLevelError,
	Logger:          defaultLog.PgxLogger(),
	RuntimeParams: map[string]string{
		"standard_conforming_strings": "on",
	},
}

// ProxyOption options
type PostgresConfigOption func(*PostgresConfig)

// WithHost new server port in string
func WithHost(host string) PostgresConfigOption {
	return func(o *PostgresConfig) {
		o.Host = host
	}
}

// WithDatabase with
func WithDatabase(dbname string) PostgresConfigOption {
	return func(o *PostgresConfig) {
		o.Database = dbname
	}
}

// WithUserPassword options
func WithUserPassword(user, password string) PostgresConfigOption {
	return func(o *PostgresConfig) {
		o.User = user
		o.Password = password
	}
}

// WithLogger new server port in string
func WithLogger(log pgx.Logger) PostgresConfigOption {
	return func(o *PostgresConfig) {
		o.Logger = log
	}
}

// WithLogLevel new server port in string
func WithLogLevel(level pgx.LogLevel) PostgresConfigOption {
	return func(o *PostgresConfig) {
		o.LogLevel = level
	}
}

// WithDebug debug
func WithDebug(l bool) PostgresConfigOption {
	return func(o *PostgresConfig) {
		o.Debug = l
	}
}

// WithRuntimeParams options
func WithRuntimeParams(r map[string]string) PostgresConfigOption {
	return func(o *PostgresConfig) {
		for k, v := range r {
			o.RuntimeParams[k] = v
		}
	}
}

// NewPostgresConfig  new postgres config
func NewPostgresConfig(opts ...PostgresConfigOption) *PostgresConfig {
	p := new(PostgresConfig)
	_ = mergo.Merge(p, defaultPostgresConfig)

	// statement_cache_mode=describe
	//	p.RuntimeParams["statement_cache_mode"] = "describe"

	for _, o := range opts {
		o(p)
	}
	return p
}

// SetLogger set logger
func (p *PostgresConfig) SetLogger(log *logger.Logger) {
	p.Logger = log.PgxLogger()
}

// SetLogger set logger
func (p *PostgresConfig) SetLogLevel(level pgx.LogLevel) {
	p.LogLevel = level
}

// Load initial running environment
func Load(configFileName string) (cfg *PostgresConfig, err error) {
	m := multiconfig.NewWithPath(configFileName) // supports TOML, JSON and YAML
	cfg = new(PostgresConfig)
	err = m.Load(cfg) // Check for error
	if err != nil {
		// panic("-----------------读取配置文件出错啦---------------")
		return
	}
	m.MustLoad(cfg) // Panic's if there is any error
	return
}
