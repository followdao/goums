package umsserver

import (
	"context"
	"strconv"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tsingson/goums/apis/go/goums/terminal"
	"github.com/tsingson/goums/grpc/session"
)

// Active  login
func (s *UmsServer) Active(ctx context.Context, in *terminal.TerminalRequest) (*flatbuffers.Builder, error) {
	tid := time.Now().Unix()
	log := s.log.Named("Import " + strconv.FormatInt(int64(tid), 10))

	if ctx.Err() == context.Canceled {
		log.Error("client cancel", zap.Error(ctx.Err()))
		return nil, status.Errorf(codes.Canceled, "Import canceled")
	}

	va := in.UnPack()

	re, err := s.terminalDbo.Active(ctx, va.SerialNumber, va.ActiveCode, va.ApkType)
	if err != nil {
		log.Error("import error", zap.Error(err))
		return terminal.ResultBuilder(tid, int64(1), "import error"), err
	}

	rt := &terminal.AccessResultT{
		Me: &terminal.AccessProfileT{
			UserID:     strconv.FormatInt(re.UserID, 10),
			ActiveDate: time.Unix(re.ActiveDate, 0).String(),
			// TODO:  replace session expired time
			Expiration: session.GenerateExpiration(),
		},
		// TODO:  replace token generator
		Token: session.GenerateToken(),
	}

	return rt.Builder(), nil
}
