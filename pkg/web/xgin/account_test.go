// go test -httptest.serve=127.0.0.1:8000 .
package xgin

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/tsingson/go-ums/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/oklog/run"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"github.com/tsingson/go-ums/model"
)

var hs *HTTPServer

func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)
	hs = NewHTTPServer()
	r := hs.SetupRouter()

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
		panic("can't run gin")
	}

	os.Exit(m.Run())
}

func TestHttpServer_RegisterHandler(t *testing.T) {

	var err error
	// defer ln.Close()

	c := &fasthttp.Client{}
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI("http://localhost:3001/register")

	req.Header.SetContentType("application/json; charset=utf-8")
	// 	req.Header.Add("User-Agent", "Test-Agent")
	req.Header.Add("Accept", "application/json")

	var body []byte

	var tests = []struct {
		email    []byte
		password []byte
	}{{[]byte("t@1.1"), []byte("123456")}}

	for _, test := range tests {
		var register = &model.AccountRequest{
			Email:    test.email,
			Password: test.password,
		}

		body, _ = json.Marshal(register)

		req.Header.SetMethod("POST")
		req.SetBody(body)

		var timeOut = time.Duration(5 * time.Second)
		err = c.DoTimeout(req, resp, timeOut)

		assert.NoError(t, err)
		// println("Error:", err.Error())
		assert.Equal(t, resp.StatusCode(), fasthttp.StatusOK)

		var ac *model.Account
		err = json.Unmarshal(resp.Body(), &ac)
		assert.NoError(t, err)
		// println("Error:", err.Error())
		assert.Equal(t, test.email, ac.Email)
	}

}

func BenchmarkHttpServer_RegisterHandler(b *testing.B) {
	var err error

	// defer ln.Close()

	c := &fasthttp.Client{}
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI("http://localhost:3001/register")

	req.Header.SetContentType("application/json; charset=utf-8")
	// 	req.Header.Add("User-Agent", "Test-Agent")
	req.Header.Add("Accept", "application/json")

	var body []byte

	var test = struct {
		email    []byte
		password []byte
	}{[]byte("t@1.1"), []byte("123456")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		var register = &model.AccountRequest{
			Email:    utils.ByteBuilder(test.email, []byte(strconv.Itoa(i))),
			Password: utils.ByteBuilder(test.password, []byte(strconv.Itoa(i))),
		}

		body, _ = json.Marshal(register)

		req.Header.SetMethod("POST")
		req.SetBody(body)

		var timeOut = time.Duration(5 * time.Second)
		err = c.DoTimeout(req, resp, timeOut)
		if err != nil {
			break
		}

	}
}
