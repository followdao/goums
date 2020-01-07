package umsserver

import (
	"context"
	"strconv"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tsingson/goums/apis/flatums"

	"github.com/tsingson/goums/apis/go/goums/terminal"
)

// Import  login
func (s *UmsServer) Import(ctx context.Context, in *flatums.TerminalList) (*flatbuffers.Builder, error) {
	// var names []string
	tid := time.Now().Unix()
	log := s.log.Named("Import " + strconv.FormatInt(int64(tid), 10))

	if ctx.Err() == context.Canceled {
		log.Error("client cancel", zap.Error(ctx.Err()))
		return nil, status.Errorf(codes.Canceled, "Import canceled")
	}
	_, err := s.terminalDbo.InsertList(ctx, in.UnPack())
	if err != nil {
		log.Error("import error", zap.Error(err))
		return terminal.ResultBuilder(tid, int64(1), "import error"), err
	}
	return terminal.ResultBuilder(tid, int64(0), "import success"), nil
}
