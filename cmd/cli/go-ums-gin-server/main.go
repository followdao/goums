package main

import (
	"github.com/tsingson/go-ums/pkg/web/xgin"
)

func main() {
	var hs *xgin.HttpServer
	hs = xgin.NewHttpServer()
	r := hs.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
