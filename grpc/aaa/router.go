package aaa

import (
	"github.com/tsingson/goums/grpc/config"
)

// InitRouter init
func (s *Aaa) initRouter() {
	s.router.GET("/", s.hello())

	s.router.GET("/healthcheck", s.healthCheck())
	s.router.GET("/resetcache", s.reset())

	s.router.POST(config.ActiveURI, recovery(s.activate()))
}

// design and code by tsingson
