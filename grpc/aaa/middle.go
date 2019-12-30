package aaa

import "github.com/valyala/fasthttp"

func recovery(next func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if rvr := recover(); rvr != nil {
				ctx.Error("recover", 500)
			}
		}()
		next(ctx)
	}
}
