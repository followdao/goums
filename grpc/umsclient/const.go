package umsclient

import (
	"time"
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
	grpcMaxSendMsgSize        = 1 << 24
	grpcMaxCallMsgSize        = 1 << 24
	grpcKeepAliveTime         = time.Second * 15
	grpcKeepAliveTimeout      = time.Second * 45
	grpcBackoffMaxDelay       = time.Second * 5
)
