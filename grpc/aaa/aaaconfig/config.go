package aaaconfig

import (
	"github.com/tsingson/multiconfig"

	"github.com/tsingson/goums/grpc/config"
	"github.com/tsingson/goums/grpc/config/liftgslbconfig"
)

// UmsConfig ums URI cfg
type UmsConfig struct {
	ActiveAuthURI   string
	RegisterAuthURI string
	PlayAuthURI     string
}

// HTTPConfig  aaastb cfg
type HTTPConfig struct {
	ServerPort    string
	CacheSizeByte int
	CacheTimeOut  int64
	// AaaClusterIP  []string
}

// CdnConfig cdn cfg
type CdnConfig struct {
	VodGslb  GslbInfo
	LiveGslb GslbInfo
}

type GslbInfo struct {
	AreaID  int
	Address []string
}

// StbConfig current cfg struct
type AAAConfig struct {
	Name            string
	Version         string
	Debug           bool
	OneMachineLimit bool

	LiftConfig *liftgslbconfig.LiftGslbConfig

	UmsConfig UmsConfig
	AaaConfig HTTPConfig
}

var defaultStbConfig = &AAAConfig{
	Name:    ServerName,
	Version: VERSION,
	Debug:   true,
	AaaConfig: HTTPConfig{
		// AaaClusterIP:  []string{"127.0.0.1"},
		ServerPort:    config.AAAPort,
		CacheSizeByte: 67108864,
		CacheTimeOut:  15,
	},
	UmsConfig: UmsConfig{
		ActiveAuthURI:   "http://127.0.0.1:3001/rpc/active",
		RegisterAuthURI: "http://127.0.0.1:3001/rpc/auth",
		PlayAuthURI:     "http://127.0.0.1",
	},

	LiftConfig: liftgslbconfig.NewLiftGslbConfig(),
}

// Default default config for testing only
func Default() *AAAConfig {
	return defaultStbConfig
}

// Load load config from file
func Load(configFile string) (*AAAConfig, error) {
	// 读取配置文件
	// d := defaultStbConfig
	m := multiconfig.NewWithPath(configFile) // supports TOML, JSON and YAML
	cfg := new(AAAConfig)
	err := m.Load(cfg) // Check for error
	if err != nil {
		return nil, err
	}
	m.MustLoad(cfg) // Panic's if there is any error

	// mergo.Merge(cfg, d)

	return cfg, nil
}

// design and code by tsingson
