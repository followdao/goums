package main

import (
	"net"
	"os"

	"github.com/fasthttp/router"
	"github.com/oklog/run"
	"github.com/valyala/fasthttp"

	"github.com/tsingson/go-ums/pkg/web/fast"
)

const (
	// ServerName  apk name
	ServerName = "account service"
	// Version  version of current program
	Version = "0.1.0"
	// MaxHTTPConnect max connect limit of fast http
	MaxHTTPConnect = 30000
	// BufferSize buffer size of fast http  in http payload
	BufferSize = 1024 * 4
)

func main() {
	var err error
	var addr = ":3001"
	err = webServer(addr)
	if err != nil {
		os.Exit(-1)
	}
}

func webServer(addr string) (err error) {

	var hs *fast.HttpServer
	// initial fasthttp server with account store in memory
	hs = fast.NewHttpServer()
	// set up router
	var r *router.Router
	{
		r = router.New()
		r.POST("/register", hs.RegisterHandler)
	}

	var ln net.Listener
	ln, err = net.Listen("tcp", addr)
	if err != nil {

		panic("tcp connect error")
	}
	s := &fasthttp.Server{
		Handler:               r.Handler,
		Name:                  ServerName,
		ReadBufferSize:        BufferSize,
		MaxRequestsPerConn:    3,
		TCPKeepalive:          true,
		LogAllErrors:          true,
		NoDefaultServerHeader: true,
		MaxRequestBodySize:    1024 * 4, // MaxRequestBodySize: 100<<20, // 100MB
		Concurrency:           MaxHTTPConnect,
		// 	Logger:                zaplog,
		// 	MaxConnsPerIP:      3,
		// TCPKeepalivePeriod:   10 * time.Second,
		// MaxKeepaliveDuration: 50 * time.Second,
	}
	var g run.Group
	{
		g.Add(func() error {
			// TODO: add log here
			return s.Serve(ln)
		}, func(e error) {
			_ = ln.Close()
		})

	}

	return g.Run()

}
