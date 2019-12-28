package main

import (
	"os"

	"github.com/tsingson/go-ums/pkg/web/fast"
)

func main() {
	var err error
	var addr = ":3001"
	err = fast.FastServer(addr)
	if err != nil {
		os.Exit(-1)
	}
}
