package liftconfig

// LiftConfig config
type LiftConfig struct {
	AddressList []string
	Subject     string
	Name        string
}

var defaultLiftConfig = &LiftConfig{
	AddressList: []string{"127.0.0.1:9292"},
	Subject:     "terminal",
	Name:        "terminal",
}

// NotifySubject
type NotifySubject struct {
	Subject string
	Name    string
}

// LiftConfigOption options
type LiftConfigOption func(*LiftConfig)

// WithAddressList new server port in string
func WithAddressList(list []string) LiftConfigOption {
	return func(o *LiftConfig) {
		o.AddressList = list
	}
}

// WithSubject new server port in string
func WithSubject(subject string) LiftConfigOption {
	return func(o *LiftConfig) {
		o.Subject = subject
		o.Name = subject
	}
}

// NewLiftConfig new config
func NewLiftConfig(opts ...LiftConfigOption) *LiftConfig {
	p := defaultLiftConfig
	for _, o := range opts {
		o(p)
	}
	return p
}
