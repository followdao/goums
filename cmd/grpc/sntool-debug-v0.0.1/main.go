package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/integrii/flaggy"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tsingson/logger"

	"github.com/tsingson/goums/grpc/config/gslbconfig"
	"github.com/tsingson/goums/grpc/serial"
	"github.com/tsingson/goums/grpc/umsclient"
)

func main() {
	runtime.MemProfileRate = 0
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)

	fmt.Println("----------------------------------------------------")
	fmt.Println("--- ", gslbconfig.VERSION, " console debug version ---")
	fmt.Println("----------------------------------------------------")
	fmt.Println(" ")

	log := logger.New(logger.WithStoreInDay(),
		logger.WithDebug(),
		logger.WithDays(31),
		logger.WithLevel(zapcore.DebugLevel))

	var fi string
	address := "127.0.0.1:8099"

	flaggy.String(&fi, "f", "file", "导入 excel 文件全路径")
	flaggy.String(&address, "a", "address", "grpc 地址, 默认为 "+address)
	flaggy.Parse()

	fi = strings.TrimSpace(fi)
	if len(fi) == 0 {
		log.Log.Error("excel 文件名称为空")
		fmt.Printf("excel 文件名称为空")
		os.Exit(-1)
	}
	in, er2 := serial.ReadList(fi)
	if er2 != nil {
		log.Log.Error("read terminal list error",
			zap.Error(er2))
		fmt.Printf("Error %v \n", er2)
		os.Exit(-1)
	}

	// --------------

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5*time.Second)))
	defer cancel()

	c, er1 := umsclient.NewAaaClient(ctx, address, false, log)
	if er1 != nil {
		log.Log.Error("NewStbGrpcClient error",
			zap.Error(er1))
		fmt.Printf("Error %v \n", er1)
		os.Exit(-1)
	}

	err := c.Import(ctx, in)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Log.Error("Import error",
					zap.Error(err))
				fmt.Printf(" error %v", err)
				os.Exit(-1)
			}
		}

		log.Fatalf("client.Import err: %v", err)
	} else {
		log.Log.Info("excels import success",
			zap.String("filename", fi))
		fmt.Println(" excels import success ")
	}
}
