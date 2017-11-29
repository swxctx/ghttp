package ghttp_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/Swxctx/ghttp"
)

func TestRequest(t *testing.T) {
	req := &ghttp.Request{
		Url:           "https://www.baidu.com",
		Method:        "GET",
		XForwardedFor: "127.0.0.1",
		ContentType:   "application/json",
	}
	resp, err := req.Do()
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	result, err := resp.Body.FromToString()
	if err != nil {
		log.Panic(err)
	}
	log.Println("resp->string:", result)

	respbs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	log.Println(respbs)
}
