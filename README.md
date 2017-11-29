<script src="http://code.jquery.com/jquery-1.7.2.min.js"></script>
<script type="text/javascript">
 //是否显示导航栏
 var showNavBar = true;
 //是否展开导航栏
 var expandNavBar = true;
 
 $(document).ready(function(){
    var h1s = $("body").find("h1");
    var h2s = $("body").find("h2");
    var h3s = $("body").find("h3");
    var h4s = $("body").find("h4");
    var h5s = $("body").find("h5");
    var h6s = $("body").find("h6");

    var headCounts = [h1s.length, h2s.length, h3s.length, h4s.length, h5s.length, h6s.length];
    var vH1Tag = null;
    var vH2Tag = null;
    var vH3Tag = null;
    var vH4Tag = null;
    var vH5Tag = null;
    var vH6Tag = null;

    for(var i = 0; i < headCounts.length; i++){
        if(headCounts[i] > 0){
            if(vH1Tag == null){
                vH1Tag = 'h' + (i + 1);
            }else if(vH2Tag == null){
                vH2Tag = 'h' + (i + 1);
            }else if(vH3Tag == null){
                vH3Tag = 'h' + (i + 1);
            }else if(vH4Tag == null){
                vH4Tag = 'h' + (i + 1);
            }else if (vH5Tag == null){
                vH5Tag = 'h' + (i + 1);
            }else if(vH6Tag == null){
                vH6Tag = 'h' + (i + 1);
            }
        }
    }
    if(vH1Tag == null){
        return;
    }

    $("body").prepend('<div class="BlogAnchor">' + 
		'<span style="color:red;position:absolute;top:-6px;left:0px;cursor:pointer;" onclick="$(\'.BlogAnchor\').hide();">×</span>' +
        '<p>' + 
            '<b id="AnchorContentToggle" title="收起" style="cursor:pointer;">目录▲</b>' + 
        '</p>' + 
        '<div class="AnchorContent" id="AnchorContent"> </div>' + 
    '</div>' );

    var vH1Index = 0;
    var vH2Index = 0;
    var vH3Index = 0;
    var vH4Index = 0;
    var vH5Index = 0;
    var vH6Index = 0;
    $("body").find("h1,h2,h3,h4,h5,h6").each(function(i,item){
        var id = '';
        var name = '';
        var tag = $(item).get(0).tagName.toLowerCase();
        var className = '';
        if(tag == vH1Tag){
            id = name = ++vH1Index;
            name = id;
            vH2Index = 0;
            className = 'item_h1';
        }else if(tag == vH2Tag){
            id = vH1Index + '_' + ++vH2Index;
            name = vH1Index + '.' + vH2Index;
            className = 'item_h2';
        }else if(tag == vH3Tag){
            id = vH2Index + '_' + ++vH3Index;
            name = vH2Index + '.' + vH3Index;
            className = 'item_h3';
        }else if(tag == vH4Tag){
            id = vH3Index + '_' + ++vH4Index;
            name = vH3Index + '.' + vH4Index;
            className = 'item_h4';
        }else if(tag == vH5Tag){
            id = vH4Index + '_' + ++vH5Index;
            name = vH4Index + '.' + vH5Index;
            className = 'item_h5';
        }else if (tag == vH6Tag){
            id = vH5Index + '_' + ++vH6Index;
            name = vH5Index + '.' + vH6Index;
            className = 'item_h6';
        }
        $(item).attr("id","wow"+id);
		$(item).addClass("wow_head");
        $("#AnchorContent").css('max-height', ($(window).height() - 180) + 'px');
        $("#AnchorContent").append('<li><a class="nav_item '+className+' anchor-link" onclick="return false;" href="#" link="#wow'+id+'">'+name+" · "+$(this).text()+'</a></li>');
    });

    $("#AnchorContentToggle").click(function(){
        var text = $(this).html();
        if(text=="目录▲"){
            $(this).html("目录▼");
            $(this).attr({"title":"展开"});
        }else{
            $(this).html("目录▲");
            $(this).attr({"title":"收起"});
        }
        $("#AnchorContent").toggle();
    });
    $(".anchor-link").click(function(){
        $("html,body").animate({scrollTop: $($(this).attr("link")).offset().top}, 500);
    });
	
	var headerNavs = $(".BlogAnchor li .nav_item");
	var headerTops = [];
	$(".wow_head").each(function(i, n){
		headerTops.push($(n).offset().top);
	});
	$(window).scroll(function(){
		var scrollTop = $(window).scrollTop();
		$.each(headerTops, function(i, n){
			var distance = n - scrollTop;
			if(distance >= 0){
				$(".BlogAnchor li .nav_item.current").removeClass('current');
				$(headerNavs[i]).addClass('current');
				return false;
			}
		});
	});

	if(!showNavBar){
		$('.BlogAnchor').hide();
	}
	if(!expandNavBar){
		$(this).html("目录▼");
        $(this).attr({"title":"展开"});
		$("#AnchorContent").hide();
	}
 });
</script>
<style>
    /*导航*/
    .BlogAnchor {
        background: #f1f1f1;
        padding: 10px;
        line-height: 180%;
        position: fixed;
        right: 48px;
        top: 48px;
        border: 1px solid #aaaaaa;
    }
    .BlogAnchor p {
        font-size: 18px;
        color: #15a230;
        margin: 0 0 0.3rem 0;
        text-align: right;
    }
    .BlogAnchor .AnchorContent {
        padding: 5px 0px;
        overflow: auto;
    }
    .BlogAnchor li{
        text-indent: 0.5rem;
        font-size: 14px;
        list-style: none;
    }
	.BlogAnchor li .nav_item{
		padding: 3px;
	}
    .BlogAnchor li .item_h1{
        margin-left: 0rem;
    }
    .BlogAnchor li .item_h2{
        margin-left: 2rem;
        font-size: 0.8rem;
    }
    .BlogAnchor li .item_h3{
        margin-left: 4rem;
        font-size: 0.8rem;
    }
    .BlogAnchor li .item_h4{
        margin-left: 6rem;
        font-size: 0.8rem;
    }
    .BlogAnchor li .item_h5{
        margin-left: 8rem;
        font-size: 0.8rem;
    }
    .BlogAnchor li .item_h6{
        margin-left: 10rem;
        font-size: 0.8rem;
    }

	.BlogAnchor li .nav_item.current{
		color: white;
		background-color: #5cc26f;
	}
    #AnchorContentToggle {
        font-size: 13px;
        font-weight: normal;
        color: #FFF;
        display: inline-block;
        line-height: 20px;
        background: #5cc26f;
        font-style: normal;
        padding: 1px 8px;
    }
    .BlogAnchor a:hover {
        color: #5cc26f;
    }
    .BlogAnchor a {
        text-decoration: none;
    }
</style>

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