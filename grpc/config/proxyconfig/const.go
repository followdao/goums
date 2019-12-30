package proxyconfig

import (
	"time"

	"github.com/valyala/fasthttp"
)

const (
	// MaxRequestBodySize: 100<<20, // 100MB
	// MaxRequestBodySize: p.Config.MaxRequestBodySize, //  100 << 20, // 100MB // 1024 * 4, // MaxRequestBodySize:
	LogFileNameTimeFormat = "2006-01-02-15"
	ExpiredDate           = "2019-12-30 00:00:00"

	MaxFttpConnect     = 30000
	ReadBufferSize     = 1024 * 2
	MaxConnsPerIP      = 5
	MaxRequestsPerConn = 100
	MaxRequestBodySize = 1024 * 1024 * 4
	MinRequestBodySize = 1024 * 2
	MaxConnsPerHost    = fasthttp.DefaultMaxConnsPerHost * 4
	Concurrency        = 5000
	BufferSize         = 1024 * 4
	UploadFileField    = "filename"
	AcceptJson         = "application/json"
	AcceptRest         = "application/vnd.pgrst.object+json"
	ContentText        = "text/plain; charset=utf8"
	ContentRest        = "application/vnd.pgrst.object+json; charset=utf-8"
	ContentJson        = "application/json; charset=utf-8"
	cacheSize          = 1024 * 1204 * 512
	IdleTimeout        = time.Duration(15) * time.Second
)
