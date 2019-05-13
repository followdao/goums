package main

import (
	"fmt"
	"net/http"

	"github.com/oklog/run"

	"github.com/tsingson/go-ums/pkg/web/xgin"
)

func main() {

	hs := xgin.NewHttpServer()
	r := hs.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	server := http.Server{
		Addr:    ":3001",
		Handler: r,
	}

	var g run.Group
	g.Add(func() error {
		return server.ListenAndServe()
	}, func(e error) {
		_ = server.Close()
	})
	err := g.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
