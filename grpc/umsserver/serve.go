package umsserver

import (
	"context"
	"net"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	mid "github.com/grpc-ecosystem/go-grpc-middleware"
	gz "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/tsingson/goums/apis/go/goums/terminal"
)

// Serve  grpc server
func (s *UmsServer) Serve(ctx context.Context) (err error) {
	log := s.log.Log.Named("serve")
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []gz.Option{
		gz.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}

	rpc := grpc.NewServer(
		grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             3 * time.Second,
			PermitWithoutStream: true,
		}),
		mid.WithUnaryServerChain(
			tags.UnaryServerInterceptor(),
			gz.UnaryServerInterceptor(log, opts...),
			// grpc_auth.UnaryServerInterceptor(auther.Authenticate),
		),
		mid.WithStreamServerChain(
			tags.StreamServerInterceptor(),
			gz.StreamServerInterceptor(log, opts...),
			// grpc_auth.StreamServerInterceptor(auther.Authenticate),
		),
	)

	// defer server.Close(ctx)
	terminal.RegisterAaaServiceServer(rpc, s)

	ln, er2 := net.Listen("tcp", s.cfg.RPCConfig.Port)
	if er2 != nil {
		log.Error("failed to listen", zap.Error(er2))
		return er2
	}
	err = rpc.Serve(ln)
	if err == nil {
		log.Info("UMS grpc server start success")
	} else {
		log.Info("UMS grpc server start fail------------------ ")
		return
	}
	return nil
}
