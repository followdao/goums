package umsserver

import (
	"context"

	"github.com/tsingson/logger"
	"go.uber.org/zap"

	"github.com/tsingson/goums/grpc/config/gslbconfig"
	"github.com/tsingson/goums/modelset/terminaldbo"
)

// UmsServer pgx postgres database access interface
type UmsServer struct {
	cfg         *gslbconfig.GslbConfig
	terminalDbo *terminaldbo.TerminalDbo
	liveDbo     *terminaldbo.TerminalDbo
	log         *logger.ZapLogger
	debugMode   bool
}

// NewUmsServer initial a grpc server
func NewUmsServer(ctx context.Context, cfg *gslbconfig.GslbConfig) (*UmsServer, error) {
	log := cfg.Log.Log.Named("NewNotify")

	terminalDbo, er3 := terminaldbo.NewTerminalDbo(ctx, cfg.UmsPostgresConfig, cfg.Log)
	if er3 != nil {
		log.Error("NewTerminalDbo error",
			zap.Error(er3))
		return nil, er3
	}

	t := &UmsServer{
		cfg:         cfg,
		log:         cfg.Log,
		debugMode:   cfg.Debug,
		terminalDbo: terminalDbo,
	}

	return t, nil
}
