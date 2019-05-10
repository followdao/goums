package main

import (
	"fmt"
	"net/http"

	"github.com/tsingson/go-ums/pkg/web/xgin"
)

func main() {
	var hs *xgin.HttpServer
	hs = xgin.NewHttpServer()
	r := hs.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	server := http.Server{
		Addr:    ":3001",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
