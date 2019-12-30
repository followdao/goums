package umsclient

import (
	"context"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc/keepalive"

	"emperror.dev/errors"

	"google.golang.org/grpc"

	"github.com/tsingson/logger"

	"github.com/tsingson/goums/apis/flatums"
	"github.com/tsingson/goums/pkg/vtils"
)

// UmsClient gRPC client
type UmsClient struct {
	aaaServiceClient flatums.AaaServiceClient
	Log              *logger.ZapLogger
	Debug            bool
}

// NewAaaClient   new epg client
func NewAaaClient(ctx context.Context, address string, debug bool,
	log *logger.ZapLogger) (client *UmsClient, err error) {
	if log == nil {
		err = errors.New("need log")
		return client, err
	}

	ctxTTL, cancel := context.WithTimeout(ctx, time.Duration(30)*time.Second)
	defer cancel()

	conn, er2 := grpc.DialContext(ctxTTL, address,
		grpc.WithBlock(), // 客户端将连接到 GPRC 服务, 直到连接成功
		grpc.WithInsecure(),
		// grpc.WithCodec(flat.FlatbuffersCodec{}),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(flatbuffers.FlatbuffersCodec{})),
		// client side
		grpc.WithInitialWindowSize(grpcInitialWindowSize),
		grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
		grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: grpcBackoffMaxDelay}),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                grpcKeepAliveTime,
			Timeout:             grpcKeepAliveTimeout,
			PermitWithoutStream: true,
		}))
	if er2 != nil {
		return client, er2
	}

	client = &UmsClient{
		aaaServiceClient: flatums.NewAaaServiceClient(conn),
		Log:              log,
		Debug:            debug,
	}

	return client, nil
}

// Import send a list of terminal to import into server DB
func (a *UmsClient) Import(ctx context.Context, v *flatums.TerminalListT) error {
	result, err := a.aaaServiceClient.Import(ctx,
		v.Builder(),
		grpc.WaitForReady(true))
	if err != nil {
		return err
	}
	if result.Code() != int64(0) {
		return errors.New(vtils.B2S(result.Message()))
	}
	return nil
}

// Active send a terminal to active
func (a *UmsClient) Active(ctx context.Context, v *flatums.TerminalRequestT) (*flatums.AccessResult, error) {
	result, err := a.aaaServiceClient.Active(ctx,
		v.Builder(),
		grpc.WaitForReady(true))
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Close  close
func (a *UmsClient) Close() error {
	// return a.conn.Close()
	return nil
}
