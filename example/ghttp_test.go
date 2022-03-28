package exampl

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/swxctx/ghttp"
)

func TestSimpleRequest(t *testing.T) {
	req := ghttp.Request{
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
}

func TestQueryStruct(t *testing.T) {
	// request-> Get http://127.0.0.1:8080?name=xc&password=xc
	type User struct {
		Name     string
		Password string
	}
	user := User{
		Name:     "xc",
		Password: "xc",
	}
	res, err := ghttp.Request{
		Url:   "http://127.0.0.1:8080",
		Query: user,
	}.Do()
	if err != nil {
		log.Panic(err)
	}
	log.Println(res)
}

func TestQueryJson(t *testing.T) {
	// Get http://127.0.0.1:8080?name=xc
	type User struct {
		Name     string `json:"name"`
		Password string `json:"-"`
		Sex      string `json:"sex,omitempty"`
	}
	user := User{
		Name:     "xc",
		Password: "xc",
		Sex:      "",
	}
	res, err := ghttp.Request{
		Url:   "http://127.0.0.1:8080",
		Query: user,
	}.Do()
	if err != nil {
		log.Panic(err)
	}
	log.Println(res)
}

func TestQuerySquash(t *testing.T) {
	// Get http://127.0.0.1:8080?id=1&name=xc&password=xc&sex=1
	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Sex      string `json:"sex"`
	}
	type GoPher struct {
		Id   int32 `json:"id"`
		User `json:",squash"`
	}
	gopher := GoPher{
		User: User{
			Name:     "xc",
			Password: "xc",
			Sex:      "1",
		},
		Id: 1,
	}
	fmt.Println(gopher)
	res, err := ghttp.Request{
		Url:   "http://127.0.0.1:8080",
		Query: gopher,
	}.Do()
	if err != nil {
		log.Panic(err)
	}
	log.Println(res)
}

func TestHeader(t *testing.T) {
	// Post http://127.0.0.1:8080?name=xc&password=xc
	type User struct {
		Name     string
		Password string
	}
	user := User{
		Name:     "xc",
		Password: "xc",
	}
	req := &ghttp.Request{
		Method:      "POST",
		Url:         "http://127.0.0.1:8080",
		Query:       user,
		ContentType: "application/json",
	}
	req.AddHeader("X-Custom", "haha")
	res, err := req.Do()
	if err != nil {
		log.Panic(err)
	}
	log.Println(res)
}

func TestCookie(t *testing.T) {
	res, err := ghttp.Request{
		Url: "http://www.baidu.com",
	}.WithCookie(&http.Cookie{Name: "c1", Value: "v1"}).Do()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
}

func TestSetTimeOut(t *testing.T) {
	res, err := ghttp.Request{
		Url:     "http://www.baidu.com",
		Timeout: 100 * time.Millisecond,
	}.Do()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
}

func TestResToJson(t *testing.T) {
	// 解析json，转换为相应结构体
	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Sex      string `json:"sex"`
	}
	var user User
	res, err := ghttp.Request{
		Url:     "http://127.0.0.1:8080",
		Timeout: 100 * time.Millisecond,
	}.Do()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res.Body.FromToJson(&user))
}

func TestProxy(t *testing.T) {
	res, err := ghttp.Request{
		Url:     "http://127.0.0.1:8080",
		Timeout: 100 * time.Millisecond,
		Proxy:   "http://127.0.0.1:8088",
	}.Do()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
}

func TestDebug(t *testing.T) {
	res, err := ghttp.Request{
		Url:       "http://127.0.0.1:8080",
		Timeout:   100 * time.Millisecond,
		ShowDebug: true,
	}.Do()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
}

func TestGzip(t *testing.T) {
	res, err := ghttp.Request{
		Method:      "POST",
		Url:         "http://www.baidu.com",
		Compression: ghttp.Gzip(),
	}.Do()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
}

func TestDeflate(t *testing.T) {
	res, err := ghttp.Request{
		Method:      "POST",
		Url:         "http://www.baidu.com",
		Compression: ghttp.Deflate(),
	}.Do()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
}

func TestForm(t *testing.T) {
	parmas := url.Values{}
	parmas.Set("idcard", "123")
	parmas.Set("name", "123")
	req := ghttp.Request{
		Method:    "POST",
		Url:       "http://www.baidu.com",
		ShowDebug: true,
		Body:      parmas,
	}
	req.AddHeader("Content-Type", "application/x-www-form-urlencoded")
	res, err := req.Do()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
}
