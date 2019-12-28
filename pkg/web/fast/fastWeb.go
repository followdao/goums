package fast

import (
	"net"

	"github.com/fasthttp/router"
	"github.com/oklog/run"
	"github.com/valyala/fasthttp"
)

// ServerName  apk name
const ServerName = "account service"

// Version  version of current program
// Version = "0.1.0"

// MaxHTTPConnect max connect limit of fast http
const MaxHTTPConnect int = 30000

// BufferSize buffer size of fast http  in http payload
const BufferSize int = 1024 * 4

func FastServer(addr string) (err error) {

	// initial fasthttp server with account store in memory
	hs := NewHTTPServer()
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
