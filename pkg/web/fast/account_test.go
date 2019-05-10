package fast

import (
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	json "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"

	"github.com/tsingson/go-ums/model"
)

var hs *httpServer

func TestMain(m *testing.M) {
	hs = NewHttpServer()
	os.Exit(m.Run())
}

func TestHttpSERV_RegisterHandler(t *testing.T) {
	var err error
	var r *router.Router
	{

		r = router.New()
		r.POST("/register", hs.RegisterHandler)
	}

	ln := fasthttputil.NewInmemoryListener()
	s := &fasthttp.Server{
		Handler: r.Handler,
	}
	serverStopCh := make(chan struct{})
	go func() {
		if err = s.Serve(ln); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		close(serverStopCh)
	}()

	// defer ln.Close()

	c := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	host := ln.Addr().String()

	req.SetRequestURI("http://" + host + "/register")

	req.Header.SetContentType("application/json; charset=utf-8")
	// 	req.Header.Add("User-Agent", "Test-Agent")
	req.Header.Add("Accept", "application/json")

	var body []byte

	var tests = []struct {
		email    string
		password string
	}{
		{"t@1.1", "123456"},
		{"email@1.2", "123452"},
	}

	for _, test := range tests {
		var register = &model.AccountRequest{
			Email:    test.email,
			Password: test.password,
		}

		body, err = json.Marshal(register)
		if err != nil {

		}

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
	var r *router.Router
	{

		r = router.New()
		r.POST("/register", hs.RegisterHandler)
	}

	ln := fasthttputil.NewInmemoryListener()
	s := &fasthttp.Server{
		Handler: r.Handler,
	}
	serverStopCh := make(chan struct{})
	go func() {
		if err = s.Serve(ln); err != nil {
			b.Fatalf("unexpected error: %s", err)
		}
		close(serverStopCh)
	}()

	// defer ln.Close()

	c := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	host := ln.Addr().String()

	req.SetRequestURI("http://" + host + "/register")

	req.Header.SetContentType("application/json; charset=utf-8")
	// 	req.Header.Add("User-Agent", "Test-Agent")
	req.Header.Add("Accept", "application/json")

	var body []byte

	var test = struct {
		email    string
		password string
	}{"t@1.1", "123456"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		var register = &model.AccountRequest{
			Email:    strBuilder(test.email, strconv.Itoa(i)),
			Password: strBuilder(test.password, strconv.Itoa(i)),
		}

		body, err = json.Marshal(register)
		if err != nil {

		}

		req.Header.SetMethod("POST")
		req.SetBody(body)

		var timeOut = time.Duration(5 * time.Second)
		err = c.DoTimeout(req, resp, timeOut)
		if err != nil {
			break
		}

	}
}

func strBuilder(args ...string) string {
	var str strings.Builder

	for _, v := range args {
		str.WriteString(v)
	}
	return str.String()
}
