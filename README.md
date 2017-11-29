<link rel="stylesheet" href="http://yandex.st/highlightjs/6.2/styles/googlecode.min.css">

<script src="http://code.jquery.com/jquery-1.7.2.min.js"></script>

<script src="http://yandex.st/highlightjs/6.2/highlight.min.js"></script>

<script>hljs.initHighlightingOnLoad();</script>

<script type="text/javascript">

 $(document).ready(function(){

    $("h2,h3,h4,h5,h6").each(function(i,item){

        var tag = $(item).get(0).localName;

        $(item).attr("id","wow"+i);

        $("#category").append('<a class="new'+tag+'" href="#wow'+i+'">'+$(this).text()+'</a></br>');

        $(".newh2").css("margin-left",0);

        $(".newh3").css("margin-left",20);

        $(".newh4").css("margin-left",40);

        $(".newh5").css("margin-left",60);

        $(".newh6").css("margin-left",80);

    });

 });

</script>

<div id="category"></div>
## Ghttp
### 序言
因为做的一些项目涉及到golang的http请求及响应json数据的封装，所以想要封装一些简单的方法用于对Http请求的参数封装及响应数据的解析，使之可以很方便的对响应数据进行转换到Struct；通过查阅一些资料以及借鉴一些开源的项目，经过对自身项目的适应处理后产生了Ghttp。   
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

```json
	res, err := ghttp.Request{
			Url:   "http://127.0.0.1:8080",
		}.Do()
```
### POST  
Ghttp使用POST请求与golang请求一致，基本请求如下所示：  

```json
	res, err := ghttp.Request{
			Method: "POST",
			Url:   "http://127.0.0.1:8080",
		}.Do()
```
### 请求参数
Ghttp参数支持interface类型，允许直接将Struce作为参数赋值传递，同时支持tags配置指定参数名称以及忽略参数空值，使用如下所示。  
#### 使用Struce作为参数  
直接使用struce作为参数时，默认元素小写后作为url参数，下面例子请求后生成url如request所示：  

```json  
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

```json  
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

```json
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

```json
	res, err := ghttp.Request{
		Url:     "http://www.baidu.com",
		Timeout: 100 * time.Millisecond,
	}.Do()
```
### 响应数据结构转换-Struct

```json
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

```json
res, err := ghttp.Request{
		Url:     "http://127.0.0.1:8080",
		Timeout: 100 * time.Millisecond,
		Proxy:   "http://127.0.0.1:8088",
	}.Do()
```