package aaa

import (
	"github.com/valyala/fasthttp"

	"github.com/tsingson/goums/pkg/vtils"
)

// healthCheck
func (s *Aaa) hello() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var ip []byte
		// var country

		ipString := ctx.Request.Header.PeekBytes([]byte("Client-IP"))
		// countryString :=  ctx.Request.Header.PeekBytes([]byte("Client-Country"))

		if len(ipString) > 0 {
			ip = ipString
		} else {
			ip = vtils.S2B(ctx.RemoteIP().String())
		}
		// if len(countryString) > 0 {
		// 	country = countryString
		// }
		//

		ctx.SetContentType("text/plain; charset=utf8")
		ctx.SetBody(ip)
		return
	}
}

// healthCheck
func (s *Aaa) healthCheck() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("text/plain; charset=utf8")
		ctx.SetBody([]byte("1"))
		return
	}
}

// healthCheck
func (s *Aaa) reset() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		s.cache.Clear()
		ctx.SetContentType("text/plain; charset=utf8")
		ctx.SetBody([]byte("reset cache"))
		return
	}
}

// design and code by tsingson
