package aaa

import (
	"net"

	"emperror.dev/errors"
	"github.com/fasthttp/router"
	lift "github.com/liftbridge-io/go-liftbridge"
	"github.com/valyala/fasthttp"

	"github.com/tsingson/logger"

	"github.com/tsingson/bytecache"
	"github.com/tsingson/bytecache/fastc"

	"github.com/tsingson/goums/grpc/aaa/aaaconfig"
	"github.com/tsingson/goums/pkg/liftclient"
)

// StbServer http server struct
type Aaa struct {
	Cfg            *aaaconfig.AAAConfig
	TerminalClient lift.Client

	ln     net.Listener
	server *fasthttp.Server

	router *router.Router

	cache     bytecache.ByteCache
	debugMode bool

	log *logger.ZapLogger
}

// NewAaa  new server
func NewAaa(cfg *aaaconfig.AAAConfig, log *logger.ZapLogger) (*Aaa, error) {
	if cfg == nil {
		return nil, errors.New("config error")
	}

	terminalClient, er1 := liftclient.Setup(cfg.LiftConfig.AddressList, cfg.LiftConfig.TerminalNotify.Subject)
	if er1 != nil {
		return nil, er1
	}

	a := &Aaa{
		Cfg:            cfg,
		log:            log,
		TerminalClient: terminalClient,

		router:    router.New(),
		cache:     fastc.NewVictoriaCache(fastc.WithPath("aaa")),
		debugMode: cfg.Debug,
	}

	//
	return a, nil
}
