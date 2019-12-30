package aaa

import (
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/tsingson/logger"

	"github.com/tsingson/goums/grpc/aaa/aaaconfig"
)

var s *Aaa

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	cfg := aaaconfig.Default()

	cfg.Debug = true

	log := logger.New(logger.WithStoreInDay(),
		logger.WithDebug(),
		logger.WithDays(31),
		logger.WithLevel(zapcore.DebugLevel))
	var err error
	s, err = NewAaa(cfg, log)

	if err != nil {
		//
		fmt.Printf("ERRR %v /n", err)
	}
	// setup(true)
	os.Exit(m.Run())
}
