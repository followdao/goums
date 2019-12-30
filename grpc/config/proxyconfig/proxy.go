package proxyconfig

// Proxy  aaastb proxy setting struct define
type ProxyConfig struct {
	ServerPort string `json:"ServerPort"`
	IP         string `json:"IP"`
	DomainName string `json:"DomainName"`

	FullRUL            string `json:"FullURL"`
	StaticFullFilePath string `json:"StaticFullFilePath"`

	// log               *logger.ZapLogger
}

var defaultProxy = &ProxyConfig{
	IP:         "*",
	ServerPort: ":80",
	DomainName: "http://127.0.0.1:80",
	FullRUL:    "http://127.0.0.1:80",
}

// ProxyOption options
type ProxyOption func(*ProxyConfig)

// WithServerPort new server port in string
func WithServerPort(port string) ProxyOption {
	return func(o *ProxyConfig) {
		o.ServerPort = port
	}
}

// WithServerPort new server port in string
// func WithLogger(log *logger.ZapLogger) ProxyOption {
// 	return func(o *Proxy) {
// 		o.log = log
// 	}
// }

// NewProxyConfig new proxy
func NewProxyConfig(opts ...ProxyOption) *ProxyConfig {
	p := defaultProxy
	for _, o := range opts {
		o(p)
	}
	return p
}

type option func(f *ProxyConfig) option

// Option sets the options specified.
// It returns an option to restore the last arg's previous value.
func (f *ProxyConfig) Option(opts ...option) (previous option) {
	for _, opt := range opts {
		previous = opt(f)
	}
	return previous
}

// Verbosity sets Foo's verbosity level to v.
func WithPort(v string) option {
	return func(f *ProxyConfig) option {
		previous := f.ServerPort
		f.ServerPort = v
		return WithPort(previous)
	}
}

func DoSomethingVerbosely(foo *ProxyConfig, port string) {
	// Could combine the next two lines,
	// with some loss of readability.
	prev := foo.Option(WithPort(port))
	defer foo.Option(prev)
	// ... do some stuff with foo under high verbosity.
}
