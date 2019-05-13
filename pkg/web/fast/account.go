package fast

import (
	"time"

	json "github.com/json-iterator/go"
	"github.com/tsingson/goutils"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"

	"github.com/tsingson/go-ums/model"
	"github.com/tsingson/go-ums/pkg/services"
)

type HttpServer struct {
	as *services.AccountStore
}

// var once sync.Once

func NewHttpServer() *HttpServer {
	as := services.New()
	return &HttpServer{
		as: as,
	}
}

func (hs *HttpServer) RegisterHandler(ctx *fasthttp.RequestCtx) {
	var err error

	// verify token
	// {
	// 	token := ctx.Request.Header.Peek("token")
	// 	err = verify(token)
	// 	if err != nil {
	// 		msg := error.Error(err)
	// 		ctx.Error(msg, fasthttp.StatusInternalServerError)
	// 		return
	// 	}
	// }
	// handle payload
	{
		transactionID := ctx.Request.Header.Peek("transactionID")
		payload := ctx.PostBody()
		var reqObj = &model.AccountRequest{}

		err = json.Unmarshal(payload, reqObj)
		if err != nil {
			errorResult(ctx, transactionID, err)
			return
		}

		var p fastjson.Parser
		v, err := p.ParseBytes(ctx.PostBody())
		if err != nil {
			errorResult(ctx, transactionID, err)
			return
		}
		// tid := v.GetStringBytes("transactionID")
		email := v.GetStringBytes("email")
		password := v.GetStringBytes("password")

		var ac *model.Account
		ac, err = hs.as.Register(goutils.B2S(email), goutils.B2S(password))
		if err != nil {
			errorResult(ctx, transactionID, err)
			return
		}

		var out []byte
		out, err = json.Marshal(ac)
		if err != nil {
			errorResult(ctx, transactionID, err)
			return
		}
		ctx.SetBody(out)
	}

}

func (hs *HttpServer) LoginHandler(ctx *fasthttp.RequestCtx) {
	var err error
	transactionID := ctx.Request.Header.Peek("transactionID")
	// handle payload
	{
		payload := ctx.PostBody()
		var reqObj = &model.AccountRequest{}

		err = json.Unmarshal(payload, reqObj)
		if err != nil {
			errorResult(ctx, transactionID, err)
			return
		}

		var p fastjson.Parser
		v, err := p.ParseBytes(ctx.PostBody())
		if err != nil {
			errorResult(ctx, transactionID, err)
			return
		}
		// tid := v.GetStringBytes("transactionID")
		email := v.GetStringBytes("email")
		password := v.GetStringBytes("password")

		var token string
		token, err = hs.as.Login(goutils.B2S(email), goutils.B2S(password))
		if err != nil {
			errorResult(ctx, transactionID, err)
			return
		}
		ctx.SetBodyString(token)
	}

}

func errorResult(ctx *fasthttp.RequestCtx, tid []byte, err error) {
	var result = model.Result{}
	result.Code = 500
	result.Msg = error.Error(err)
	result.TimeStamp = time.Now().UnixNano()
	ctx.Response.Header.SetBytesV("tid", tid)
	ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
	out, _ := json.Marshal(result)
	ctx.SetBody(out)

}
