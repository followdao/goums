package liftgslbconfig

// LiftGslbConfig config
type LiftGslbConfig struct {
	AddressList    []string
	TerminalNotify NotifySubject
}

// NotifySubject notify subject
type NotifySubject struct {
	Subject string
	Name    string
}

var defaultLiftConfig = &LiftGslbConfig{
	AddressList: []string{"127.0.0.1:9292"},
	TerminalNotify: NotifySubject{
		Subject: "terminal",
		Name:    "terminal",
	},
}

// LiftConfigOption options
type LiftConfigOption func(*LiftGslbConfig)

// WithAddressList new server port in string
func WithAddressList(list []string) LiftConfigOption {
	return func(o *LiftGslbConfig) {
		o.AddressList = list
	}
}

// NewLiftGslbConfig new config
func NewLiftGslbConfig(opts ...LiftConfigOption) *LiftGslbConfig {
	p := defaultLiftConfig
	for _, o := range opts {
		o(p)
	}
	return p
}
