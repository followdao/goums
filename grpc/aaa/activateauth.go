package aaa

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

// activate for api /api/activate
func (s *Aaa) activate() func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		log := s.log.Named("activate-" + strconv.FormatInt(int64(ctx.ID()), 10))
		log.Info("----------------------------------------------------")

		return
	}
}
