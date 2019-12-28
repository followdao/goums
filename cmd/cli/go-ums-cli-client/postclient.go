package main

import (
	"fmt"
	"time"

	json "github.com/json-iterator/go"
	"github.com/sanity-io/litter"
	"github.com/tsingson/goutils"

	"github.com/valyala/fasthttp"

	"github.com/tsingson/go-ums/model"
	"github.com/tsingson/go-ums/pkg/services"
)

func postRegister(url, email, password string) (err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	req.SetRequestURI(url)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.Add("User-Agent", "Test-Agent")
	// req.Header.Add("Accept", "application/json")

	req.Header.SetMethod("POST")

	var body []byte

	var register = &model.AccountRequest{
		Email:    email,
		Password: password,
	}
	transactionID := services.GenerateIDString()

	fmt.Println("-------------------> http request object ")
	litter.Dump(register)

	body, err = json.Marshal(register)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	req.Header.SetMethod("POST")
	req.Header.Set("transactionID", transactionID)
	req.SetBody(body)

	var timeOut = time.Duration(5 * time.Second)
	err = fasthttp.DoTimeout(req, resp, timeOut)

	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Println("-------------------> http response")
	if resp.StatusCode() == fasthttp.StatusOK {
		var ac *model.AccountProfile
		err = json.Unmarshal(resp.Body(), &ac)
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
		litter.Dump(resp.StatusCode())
		litter.Dump(ac)
	} else if resp.StatusCode() == fasthttp.StatusInternalServerError {
		litter.Dump(resp.StatusCode())
		litter.Dump(goutils.B2S(resp.Body()))
		return
	}
	return
}
