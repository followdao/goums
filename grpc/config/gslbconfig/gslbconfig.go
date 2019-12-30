package gslbconfig

import (
	"github.com/tsingson/multiconfig"
	"go.uber.org/zap/zapcore"

	"github.com/tsingson/logger"

	"github.com/tsingson/goums/dbv4/postgresconfig"
	"github.com/tsingson/goums/grpc/config/liftgslbconfig"
)

// ConfigFilename  default config name
const (
	ConfigFilename = "gslbsync-config.toml"
	VERSION        = "v0.0.1"
)

// RPCConfig gRPC config
type RPCConfig struct {
	Port string
	Host string
}

// SyncConfig sync config
type GslbConfig struct {
	Name    string
	Version string

	RPCConfig         RPCConfig
	UmsPostgresConfig *postgresconfig.PostgresConfig
	LiftConfig        *liftgslbconfig.LiftGslbConfig
	Log               *logger.ZapLogger
	Debug             bool
}

var defaultLog = logger.New(logger.WithStoreInDay(),
	logger.WithDebug(),
	logger.WithDays(31),
	logger.WithLevel(zapcore.DebugLevel))

var umsPostgres = postgresconfig.NewPostgresConfig(postgresconfig.WithLogger(defaultLog.PgxLogger()),
	postgresconfig.WithDatabase("goums"))

var defaultSycConfig = &GslbConfig{
	Name:              "umsserver",
	Version:           VERSION,
	UmsPostgresConfig: umsPostgres,
	LiftConfig:        liftgslbconfig.NewLiftGslbConfig(),
	RPCConfig: RPCConfig{
		Port: ":8099",
		Host: "127.0.0.1",
	},
	Log:   defaultLog,
	Debug: false,
}

// LiftConfigOption options
type LiftConfigOption func(*GslbConfig)

// WithAddressList new server port in string
func WithDebug() LiftConfigOption {
	return func(o *GslbConfig) {
		o.Debug = true
	}
}

// NewGslbConfig new config
func NewGslbConfig(opts ...LiftConfigOption) *GslbConfig {
	p := defaultSycConfig
	for _, o := range opts {
		o(p)
	}
	return p
}

// New new sync config
func New(cfg1 *postgresconfig.PostgresConfig, cfg2 *liftgslbconfig.LiftGslbConfig) *GslbConfig {
	p := defaultSycConfig
	p.UmsPostgresConfig = cfg1
	p.LiftConfig = cfg2
	return p
}

// Load initial running environment
func Load(configFileName string) (*GslbConfig, error) {
	// d := defaultSycConfig
	m := multiconfig.NewWithPath(configFileName) // supports TOML, JSON and YAML
	cfg := new(GslbConfig)
	err := m.Load(cfg) // Check for error
	if err != nil {
		// panic("-----------------读取配置文件出错啦---------------")
		return nil, err
	}
	m.MustLoad(cfg) // Panic's if there is any error

	cfg.Log = defaultLog
	cfg.UmsPostgresConfig.Logger = defaultLog.PgxLogger()

	// _ = mergo.Merge(d, cfg)

	return cfg, nil
}
