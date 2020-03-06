## Ghttp
[![Build Status](https://travis-ci.org/swxctx/ghttp.svg?branch=master)](https://travis-ci.org/swxctx/ghttp)
[![Go Report Card](https://goreportcard.com/badge/github.com/swxctx/ghttp)](https://goreportcard.com/report/github.com/swxctx/ghttp)
[![GoDoc](http://godoc.org/github.com/swxctx/ghttp?status.svg)](http://godoc.org/github.com/swxctx/ghttp)

### install
```
$ go get -u github.com/swxctx/ghttp
```

### 序言  
Gohttp主要是对golang-http请求的一些简要封装，使用ghttp可以很方便的对http的请求、响应等做操作。目前具备如下几种功能：  
- 支持Http常用请求  
- 支持query及body参数  
- 参数支持interface结构  
- 支持灵活定制参数
- 支持灵活添加Http请求头
- 支持cookie
- 支持超时设置
- 支持响应数据转换为Struct
- 支持代理设置  
具体使用方法可以参照Test文件，里面列举了几种使用方法及情况。

### GET
Ghttp默认为Get请求，基本请求如下所示：  

```
res, err := ghttp.Request{
	Url:   "http://127.0.0.1:8080",
}.Do()
```
### POST  
Ghttp使用POST请求与golang请求一致，基本请求如下所示：  

```
res, err := ghttp.Request{
	Method: "POST",
	Url:   "http://127.0.0.1:8080",
}.Do()
```
### 请求参数
Ghttp参数支持interface类型，允许直接将Struct作为参数赋值传递，同时支持tags配置指定参数名称以及忽略参数空值，使用如下所示。  
#### 使用Struct作为参数  
直接使用struct作为参数时，默认元素小写后作为url参数，下面例子请求后生成url如request所示：  

```  
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
```
#### Tags配置参数
Ghttp支持将参数在结构体中配置为指定名称，关键字"-"表示此参数忽略不会拼接到url中，"omitempty"关键词表示该字段为空时不做拼接，可参考下面例子生成的url。  

``` 
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
```
#### 结构体嵌套  
Ghttp支持结构体嵌套的方式拼接参数，这应该也是最为常见的一种方式，例子如下：  
```json
``` 
### Header
Ghttp支持Head添加处理，如下所示：

```
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
```

### Cookie

```
res, err := ghttp.Request{
	Url:     "http://www.baidu.com",
	Timeout: 100 * time.Millisecond,
}.Do()
```
### 响应数据结构转换-Struct

```
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
```

### Proxy

```
res, err := ghttp.Request{
	Url:     "http://127.0.0.1:8080",
	Timeout: 100 * time.Millisecond,
	Proxy:   "http://127.0.0.1:8088",
}.Do()
```
