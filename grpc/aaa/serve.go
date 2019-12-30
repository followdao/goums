package aaa

import (
	"github.com/oklog/run"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"go.uber.org/zap"

	"github.com/tsingson/goums/grpc/aaa/aaaconfig"
	"github.com/tsingson/goums/grpc/config"
	"github.com/tsingson/goums/grpc/config/proxyconfig"
)

// Serve run
func (s *Aaa) Serve() (err error) {
	log := s.log.Log.Named("aaa gslb server")

	s.initRouter()
	// reuse port
	s.ln, err = reuseport.Listen("tcp4", config.AAAPort)

	if err != nil {
		log.Error("connect error",
			zap.String("port", config.AAAPort),
			zap.Error(err))
		return err
	}

	// run fasthttp serv
	s.server = &fasthttp.Server{
		Handler:        s.router.Handler,
		Name:           aaaconfig.ServerName,
		ReadBufferSize: proxyconfig.BufferSize,
		// MaxConnsPerIP:         proxyconfig.MaxConnsPerIP,
		// MaxRequestsPerConn:    proxyconfig.MaxRequestsPerConn,
		MaxRequestBodySize: proxyconfig.MinRequestBodySize, // MaxRequestBodySize: 100<<20, // 100MB
		Concurrency:        proxyconfig.Concurrency,
		// TCPKeepalive:       true,
		Logger: s.log,
		// MaxConnsPerIP:      3,
		// MaxRequestsPerConn: 3,
		// 	TCPKeepalivePeriod: 10 * time.Second,
	}

	var g run.Group
	{
		g.Add(func() error {
			return s.server.Serve(s.ln)
		}, func(e error) {
			_ = s.ln.Close()
		})
	}
	// fmt.Println("-------------------------- stbaaa HTTP")
	return g.Run()
	// return s.server.Serve(s.ln)
}
