package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"go.uber.org/zap"

	"github.com/tsingson/goums/grpc/config/gslbconfig"
	"github.com/tsingson/goums/grpc/umsserver"
	"github.com/tsingson/goums/pkg/vtils"
)

func main() {
	runtime.MemProfileRate = 0
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)

	fmt.Println("----------------------------------------------------")
	fmt.Println("--- ", gslbconfig.VERSION, " console debug version ---")
	fmt.Println("----------------------------------------------------")
	fmt.Println(" ")

	path, _ := vtils.GetCurrentExecDir()
	filename := path + "/" + gslbconfig.ConfigFilename
	cfg, err := gslbconfig.Load(filename)
	if err != nil {
		fmt.Println("load config file ", path+"/"+gslbconfig.ConfigFilename, " error %v", err)
		os.Exit(-1)
	}

	// litter.Dump(cfg)
	log := cfg.Log
	cfg.Debug = true

	stopSignal := make(chan struct{})

	// =============================
	log.Info("trying to start daemon")

	// =============================
	var s *umsserver.UmsServer
	ctx := context.Background()
	s, err = umsserver.NewUmsServer(ctx, cfg)
	if err != nil {
		fmt.Printf("Unable to connect to database, Err: %v ", err)
		os.Exit(1)
	}

	err = s.Serve(ctx)
	if err != nil {
		log.Error("Unable to connect to database",
			zap.Error(err))
		fmt.Printf("Unable to connect to database, Err: %v ", err)
		os.Exit(1)
	}

	<-stopSignal
}
