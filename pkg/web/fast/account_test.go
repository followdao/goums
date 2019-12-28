package fast

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/tsingson/go-ums/pkg/utils"

	json "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"

	"github.com/valyala/fasthttp"

	"github.com/tsingson/go-ums/model"
)

var hs *HTTPServer

func TestMain(m *testing.M) {

	// _ = flag.Set("alsologtostderr", "true")
	// _ = flag.Set("log_dir", "/tmp")
	// _ = flag.Set("v", "3")
	// flag.Parse()

	var addr = ":3001"
	_ = FastServer(addr)
	os.Exit(m.Run())

	/**
	  fmt.Println("begin")
	      m.Run()
	      fmt.Println("end")
	*/

}

func TestHttpSERV_RegisterHandler(t *testing.T) {
	var err error
	// var r *router.Router
	// {
	//	r = router.New()
	//	r.POST("/register", hs.RegisterHandler)
	// }
	//
	// ln := fasthttputil.NewInmemoryListener()
	// s := &fasthttp.Server{
	//	Handler: r.Handler,
	// }
	//
	// var g run.Group
	// g.Add(func() error {
	//	return s.Serve(ln)
	// }, func(e error) {
	//	_ = s.Shutdown()
	// })
	//
	// err = g.Run()
	// assert.NoError(t, err)

	// defer ln.Close()

	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	// host := ln.Addr().String()

	url := "http://localhost:3001/register"

	// req.SetRequestURI("http://" + host + "/register")
	req.SetRequestURI(url)

	req.Header.SetContentType("application/json; charset=utf-8")
	// 	req.Header.Add("User-Agent", "Test-Agent")
	req.Header.Add("Accept", "application/json")

	var body []byte

	var tests = []struct {
		email    string
		password string
	}{
		{"t@1.1", "123456"},
		{"t@1.2", "123452"},
	}

	// t.Parallel()

	for _, test := range tests {
		var register = &model.AccountRequest{
			Email:    test.email,
			Password: test.password,
		}

		body, _ = json.Marshal(register)

		req.Header.SetMethod("POST")
		req.SetBody(body)

		var timeOut = time.Duration(5 * time.Second)
		err = fasthttp.DoTimeout(req, resp, timeOut)

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
	// var err error
	// var r *router.Router
	// {
	//
	//	r = router.New()
	//	r.POST("/register", hs.RegisterHandler)
	// }
	//
	// ln := fasthttputil.NewInmemoryListener()
	// s := &fasthttp.Server{
	//	Handler: r.Handler,
	// }
	//
	// var g run.Group
	// g.Add(func() error {
	//	return s.Serve(ln)
	// }, func(e error) {
	//	_ = s.Shutdown()
	// })
	//
	// err = g.Run()
	// assert.NoError(b, err)
	//
	// // defer ln.Close()
	//
	// c := &fasthttp.Client{
	//	Dial: func(addr string) (net.Conn, error) {
	//		return ln.Dial()
	//	},
	// }
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	// host := ln.Addr().String()
	//
	// req.SetRequestURI("http://" + host + "/register")

	url := "http://localhost:3001/register"

	// req.SetRequestURI("http://" + host + "/register")

	var body []byte

	var test = struct {
		email    string
		password string
	}{"t@1.1", "123456"}

	b.ResetTimer()
	// b.RunParallel(func(pb *testing.PB) {
	//	for pb.Next() {
	//		Fibonacci(5000000)
	//	}
	// })
	for i := 0; i < b.N; i++ {
		req.Reset()
		resp.Reset()

		req.SetRequestURI(url)

		req.Header.SetContentType("application/json; charset=utf-8")
		// 	req.Header.Add("User-Agent", "Test-Agent")
		req.Header.Add("Accept", "application/json")

		var register = &model.AccountRequest{
			Email:    utils.StrBuilder(test.email, strconv.Itoa(i)),
			Password: utils.StrBuilder(test.password, strconv.Itoa(i)),
		}

		body, _ = json.Marshal(register)

		req.Header.SetMethod("POST")
		req.SetBody(body)

		var timeOut = time.Duration(5 * time.Second)
		err := fasthttp.DoTimeout(req, resp, timeOut)
		if err != nil {
			break
		}

	}
}
