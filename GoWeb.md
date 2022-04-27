# 1. 基础知识

## 1.1 简易web程序

1. 创建好项目的空文件之后，初始化一个模块

```
go mod init github.com/solenovex/web

/Users/qinjianquan/goweb/go.mod //在go mod 文件中查看

module github.com/solenovex/web //初始化时声明的模块

go 1.17
```

2.创建一个简单的web应用程序

/Users/qinjianquan/goweb/main.go

```
func main() {

	//1.注册一个函数，让其可以对web请求进行响应
	//处理web请求，我们一般用http包的下列函数，第一个参数相当于路由地址，"/"是根地址，表示响应所有的请求
	//第二个参数是个函数，一个是接口：用来写响应，一个是结构体的指针：包括传入请求的所有信息
	//这个函数是个回调函数，当请求到达时，函数就会执行（函数本身是参数）
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//接口w有write方法，所以可以调用
		w.Write([]byte("hello world"))
	})
	//2.设置web服务器，启动这个server
	//监听访问localhost:8080的请求，第二个参数默认调用Handler接口里面的方法（可以理解为一个路由器）
	//意思就是收到请求后，将所访问的路径分别匹配
	http.ListenAndServe("localhost:8080", nil) //DefaultServeMux
}
```

3.运行

```
go run main.go //终端运行
http://localhost:8080/ //web访问
hello world //结果页面
```

## 1.2 Handle请求

### 1.2.1 创建webserver的两种方式

在web编程中，处理请求的单词有两个Handle和Process，前者更多指的是响应，后者更多的是指处理过程

 go语言使用Handler来接受web请求，每接收到一个请求，go语言就会创建一个goroutine，具体而言是由DefaultServeMux来完成的

go语言中创建web server的方式有很多中，使用http.ListenAndServe（）是其中的一种，当第一个参数为空时，就是所有网络接口的80端口，第二个参数如果是nil，那就是DefaultServeMux

```
func main() {
	
	//第一种方式
	//http.ListenAndServe("localhost:8080", nil)
	//第二种方式：如下等同于上面的过程，但是分开写稍微更加灵活一点
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	server.ListenAndServe()

}
```

但是上述web应用只能走http，不能走https，要走要使用配套的两个函数，我们后面再讲

```
server.ListenAndServeTLS()
http.ListenAndServeTLS()
```

### 1.2.2 Handler

**handler**

handler 是一个接口，定义如下，也就是说任何类型，只要实现了ServeHTTP(w ResponseWriter, r *Request)这个方法，它就是一个handler

```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

```
var defaultServeMux ServeMux //DefaultServeMux其实也是一个handler
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) //它也是实现了ServeHTTP(w ResponseWriter, r *Request)这个方法
```

实现一个自定义的handler,所有请求都用打印“hello world”来处理

```
type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func main() {

	mh := myHandler{}

	//第一种方式
	//http.ListenAndServe("localhost:8080", nil)
	//第二种方式：如下等同于上面的过程，但是分开写稍微更加灵活一点

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}
	server.ListenAndServe()
}
```

**Handler函数 - http.Handle**

handler函数指的是行为与handler类似的函数，也就是函数签名和ServeHTTP方法的签名一样

事实上，在一个web应用中，使用多个handler对访问不同路径的请求予以处理，才是合理的操作

如何实现这一点？

首先不指定Server struct里面的Handler 字段值，这样就可以使用DefaultServeMux

接着使用http.Handle将某个Handler附加到DefaultServeMux

DefaultServeMux是ServeMux的指针变量，ServeMux有一个Handle方法，http包还有一个Handle函数，调用它实际上就是调用的是Handle方法

```
func Handle(pattern string, handler Handler)
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

先让http.Serve参数中handler 为nil

接着使用http.Handle向DefaultServeMux注册

```
//先自定义一个类型
type helloHandler struct{}

//为它实现ServeHTTP方法后就会将它变成一个handler
func (mh *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

//再自定义一个类型
type aboutHandler struct{}

//同样的，实现ServeHTTP方法
func (mh *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}

func main() {

	mh := helloHandler{} //初始化变量
	a := aboutHandler{}  //初始化变量

	//第一种方式
	//http.ListenAndServe("localhost:8080", nil)
	//第二种方式：如下等同于上面的过程，但是分开写稍微更加灵活一点

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	//向DefaultServeMux注册，此时再过来符合条件的请求就会经由DefaultServeMux分发给该注册过的Handle
	//其实就是把我们自定义的handler加入到DefaultServeMux分配的备选handler集合之中
	http.Handle("/hello", &mh) //注册
	http.Handle("/about", &a)  //同上也是注册

	server.ListenAndServe()
}
```

现在我们就能在浏览器访问指定的路径（http://localhost:8080/about 和 http://localhost:8080/hello），并可以获取写入的内容

**Handler函数 - http.HandleFunc**

http.HandleFunc 原理：它是一个函数类型，它可以将某个具有适当签名的函数f适配成一个Handler，而这个Handler具有方法f

```
//先自定义一个类型
type helloHandler struct{}

//为它实现ServeHTTP方法后就会将它变成一个handler
func (mh *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

//再自定义一个类型
type aboutHandler struct{}

//同样的，实现ServeHTTP方法
func (mh *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func main() {

	mh := helloHandler{} //初始化变量
	a := aboutHandler{}  //初始化变量

	//第一种方式
	//http.ListenAndServe("localhost:8080", nil)
	//第二种方式：如下等同于上面的过程，但是分开写稍微更加灵活一点

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	//1.使用Handle向DefaultServeMux注册，此时再过来符合条件的请求就会经由DefaultServeMux分发给该注册过的Handle
	//其实就是把我们自定义的handler加入到DefaultServeMux分配的备选handler集合之中
	http.Handle("/hello", &mh) //注册
	http.Handle("/about", &a)  //同上也是注册

	//2.使用HandleFunc向DefaultServeMux注册
	//在签名中定义func
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home!"))
	})
	//在外面定义，直接传函数名
	http.HandleFunc("/welcome", welcome)
	
	//3.也可以使用http.HandlerFunc()将函数转换为为一个handler，其实就是将函数名称变为ServeHTTP
	//type HandlerFunc func(ResponseWriter, *Request)
	//
	//// ServeHTTP calls f(w, r).
	//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	//	f(w, r)
	//}
	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()
}
```

为什么要使用两种注册方式呢？因为我们看到这样就不用修改原来的函数了，直接把它转换为接口里面的方法

## 1.3 内置的handlers

Go语言可以自定义handler以及将现有符合条件的函数转换为handler来处里访问特定路径的web请求。另外Go语言也内置了五种类型的handler

```
func main() {
	//1. func NotFoundHandler() Handler,给每个请求都返回 "404 page not found"
	http.NotFoundHandler()
	//2. func RedirectHandler(url string, code int) Handler,将要访问的url页面通过跳转码跳转到另一个url页面
	//常见的有StatusMovedPermanently, StatusFound or StatusSeeOther，分别是301，302，303
	http.RedirectHandler("/welcome", 301)
	//3. func StripPrefix(prefix string, h Handler) Handler,将请求的url去掉指定的前缀，然后调用另一个handler，如果不符合，则404
	//有点像中间件，修饰了另外一个handler
	http.StripPrefix("/welcome", http.NotFoundHandler())
	//4.func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
	//time.Duration 实际上是int64的一个别名，这个handler也相当于一个修饰器
	http.TimeoutHandler(http.NotFoundHandler(), 3, "timeout")
	//5. func FileServer(root FileSystem) Handler,参数是一个接口，得传入一个实现了接口的类型
	http.FileServer(http.Dir("/welcome"))
}
```

使用web server测试http.FileServer()，先准备一份html和css文件放入独立文件夹但与main函数放在同一目录下

```
func main() {

	//第一种方式
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "index"+r.URL.Path)
	//})
	//
	//http.ListenAndServe(":8080", nil)
	//第二种方式，使用 http.FileServer()
	http.ListenAndServe(":8080", http.FileServer(http.Dir("index")))
}
```

## 1.4 request

本节讲述对http请求的处理过程

### 1.4.1 HTTP消息

**HTTP消息**

HTTP Request和HTTP Response，结构相同

```
请求（响应）行

0个或者多个header

空行

可选的消息体
```

```
1.请求行
格式：
请求方式  请求资源  请求协议/版本
GET  虚拟路径/login.html HTTP/1.1

	请求方式：
	 HTTP中目前有9种请求方式，常用的有2种
		 GET：
			1.请求参数在请求行中，在URL后面
			2.请求的url长度有限制的
			3.不太安全
		 POST：
		 	1.请求参数封装在请求体中
		 	2.请求的url长度没有限制
		 	3.相对安全

2.请求头
	这部分的作用就是浏览器告诉服务器自身的一些信息
	Host:表示请求的主机
	User-Agent:浏览器告诉服务器我访问你使用的浏览器版本信息 
		可以在服务器端获取该头的信息，解决浏览器的兼容性问题
	Referer：告诉服务器我（当前的请求）从哪 里来的
			用来防盗链和做统计工作

3.请求空行
	就是用于分割POST请求的请求头和请求体的

4.请求体(真正的的内容所在)
	封装POST请求消息的请求参数的
```

**Request**

go语言中，net/http包提供了用于表示HTTP消息的结构

```
type Request struct {
	//几个重要的字段
	URL //通用形式：scheme：//[userinfo@]host/path[?query][#fragment] //[?query]查询字符串[#fragment]
	Header
	Body
	Form,PsotForm,MultipartForm
	...
}
另外,还可以使用Request的一些方法访问请求中的cookie，URL，User Agent等信息
request既可以用在服务端，也可以用在客户端
```

**URL Fragment**

```
URL Fragment
如果请求从浏览器发出，则无法从URL中提取Fragment字段值，因为浏览器会在发送时去掉它
```

```
func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.Fragment)
	})

	server.ListenAndServe()

}
访问http://localhost:8080/url#fragement无返回值
curl localhost:8080/url#frag //使用curl也无返回值
//curl 是一种命令行工具，作用是发出网络请求，然后获取数据，显示在"标准输出"（stdout）上面
```

**Request Header**

```
Request Header
header 是一个map，用来表述HTTP Header 里面的Key-Value对
key string,value []string,value里面第一个元素就是新的header值
使用append为指定的key增加一个header值

获取header
r.Header 返回的是map
r.Header["Accept-Encoding"] //返回的是[]string
r.Header.Get("Accept-Encoding") //返回的是string，使用逗号隔开
```

```
func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}

	//http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, r.URL.Fragment)
	//})
	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintln(w, r.URL.Fragment)
		fmt.Fprintln(w, r.Header)
		fmt.Fprintln(w, r.Header["Accept-Encoding"])
		fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))

	})

	server.ListenAndServe()
}
```

不同方式获取header得到的结果，如下是使用vs code插件 REST Client发送的请求

```
map[Accept-Encoding:[gzip, deflate] Connection:[close] User-Agent:[vscode-restclient]]
[gzip, deflate]
gzip, deflate
```

**Request Body**

获取body得到的结果，如下是使用vs code插件 REST Client发送的请求

```
Request Body

Body io.ReadCloser body类型是一个接口，接口里面也是两个接口,使用Read方法读取body内容

type ReadCloser interface {
    Reader
    Closer
}
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}
```

```
HTTP/1.1 200 OK
Date: Thu, 21 Apr 2022 05:30:07 GMT
Content-Length: 70
Content-Type: text/plain; charset=utf-8
Connection: close

{
  "name": "Sample",
  "time": "Wed,21 Oct 2015 18:27:50 UTC"
}
```

**URL Query** 

```
URL Query 

RawQuery 会提供实际查询的字符串,在问号后面表示，有两对值，如：id=123&thread_id=456
例如：http：//www.example.com/post?id=123&thread_id=456
还可以通过Request的Form字段得到key-value对

http：//www.example.com/post?id=123&thread_id=456
在URL中使用查询字符串向后端发送数据，这种通常是GET请求
r.URL.RawQuery 会提供实际查询的原始字符串，如id=123&thread_id=456
r.URL.Query会提供查询字符串对应的map[string][]string
```

```
func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}

	//1.#fragment
	//http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, r.URL.Fragment)
	//})

	//2.Header
	// http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
	// 	//fmt.Fprintln(w, r.URL.Fragment)
	// 	fmt.Fprintln(w, r.Header)
	// 	fmt.Fprintln(w, r.Header["Accept-Encoding"])
	// 	fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))

	// })

	//3.Body
	//http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
	//	length := r.ContentLength //获取内容的长度
	//	body := make([]byte, length)
	//	r.Body.Read(body) //获取body，  并将其读到上述bytes中
	//
	//	fmt.Fprintln(w, string(body))
	//})

	//4.URL
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query() //返回的是一个map
		fmt.Println(query)

		id := query["id"] //返回的是id []string
		log.Println(id)

		name := query.Get("name") //返回的是[]string中的第一个元素
		log.Println(name)
	})

	server.ListenAndServe()
}
```

```
http://localhost:8080/home?id=123&name=Dave&id=456&name=Nick
```

## 1.5 Form

**1. 使用表单发送POST请求**

```
HTML表单
<form action = "/process" method = "post"> //action对应的是处理的路径，method 对应的是请求类型
    <input type = "text" name="first_name"/> //input 是输入框
    <input type = "text" name="last_name"/> //输入的数据是以name-value对形式存放的，name是现有值，value是输入值
    <input type = "sumit"/> //当我们点击提交按钮之后，这个HTML表单里面的数据（两对数据）就会以name-value对的形式以post请求发送出去
    //数据内容放在POST请求的Body里面
</form>
```

name-value对在Body里面的格式

通过POST发送的name-value数据对的格式可以通过表单的Content Type来指定，也就是enctype属性

enctype 属性的默认值：application/x-www-form-urlencoded

```
<form action = "/process" method = "post" enctype="application/x-www-form-urlencoded"> //action对应的是处理的路径，method 对应的是请求类型
    <input type = "text" name="first_name"/> //input 是输入框
    <input type = "text" name="last_name"/> //输入的数据是以name-value对形式存放的，name是现有值，value是输入值
    <input type = "sumit"/> //当我们点击提交按钮之后，这个HTML表单里面的数据（两对数据）就会以name-value对的形式以post请求发送出去
    //数据内容放在POST请求的Body里面
</form>
```

浏览器被要求至少要支持：application/x-www-form-urlencoded 、multipart/form-data(它也是一个Content Type)

如果是HTML5（超文本标记语言第五次重大更新产品）的话，还要求支持text/plain

如果encrypt 是application/x-www-form-urlencoded,那么浏览器会将表单数据编码到查询字符串里面。如：

first_name=sau%20sheong&last_name=chang

如果entype是multipart/form-data,那么

​	每一个name-value对都会被转化为一个MIME消息部分

​	每一个部分都有自己的Content Type和Content Disposition，大概形式如下

```
Content Disposition:form-data;name="first_name"
sau sheong

Content Disposition:form-data;name="last_name"
chang
```

如何选择这两种encype？

简单文本：表单URL编码

大量数据，如上传文件：multipart-MIME

​	甚至可以把二进制数据通过选择Base64编码，来当作文本进行发送

表单的method还可以是GET

```
<form action = "/process" method = "get" > 
    <input type = "text" name="first_name"/> 
    <input type = "text" name="last_name"/> 
    <input type = "sumit"/> 
</form>
```

GET请求没有Body，所有的数据都通过URL的name- value对来发送

**Request 的Form字段**

Request上的函数可以从URL或和Body中提取数据，通过如下字段：

​	Form

​	PostForm

​	MultipartForm

Form 里面的数据是key- value对

通常的做法是

​	先调用ParseForm或者ParseMultipartForm来解析Request

​	解析完之后再访问上述三个字段

将表单编写到html中

```
<!DOCTYPE html>
<html lang="en">
<head>
     <meta charset="UTF-8">
     <meta http-equiv="X-UA-Compatible" content="IE=edge">
     <meta name="viewport" content="width=device-width, initial-scale=1.0">
     <title>Document</title>
</head>
<body>
<form action="http://localhost:8081/processes" method="post" enctype="application/x-www-form-urlencoded">
     <input type="text" name="first_name"/>
     <input type="text" name="last_name"/>
     <input type="submit"/>
</form>
     
</body>
</html>
```

先运行，然后再从浏览器打开表单路径，填写表单信息

```
func main() {

	server := http.Server{
		Addr: "localhost:8081",
	}

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "form"+r.URL.Path) 因为在表单里面写了地址了，所以不用再写访问路径
		r.ParseForm() //先解析表单
		fmt.Fprintln(w, r.Form)
	})

	server.ListenAndServe()
}
```

返回结果，其是Form就是一个Map[string] [] string

```
map[first_name:[yivei] last_name:[shi]]
```

**Request 的PostForm字段**

如果只想得到value，可以使用r.Form["first_name"] ,这会返回“firsr_name”所对应的value（[] string）,但是如果url中有同样的key，则返回的slice既包含表单里的值，也包含URL中的值（默认表单中的值在前面）。使用PostForm可以只获取表单中的值

```
<form action="http://localhost:8081/process?first_name=Nick" method="post" enctype="application/x-www-form-urlencoded">
     <input type="text" name="first_name"/>
     <input type="text" name="last_name"/>
     <input type="submit"/>
</form>
```

```
func main() {

	server := http.Server{
		Addr: "localhost:8081",
	}

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "form"+r.URL.Path) 因为在表单里面写了地址了，所以不用再写访问路径
		r.ParseForm() //先解析表单
		fmt.Fprintln(w, r.PostForm["first_name"])
	})

	server.ListenAndServe()
}
```

输出结果，Nick不会被输出

```
[yivei]
```

**Request 的MultipartForm字段**

以上两个字段都不支持multipart/form-data

同理要使用这个需要先解析

```
<form action="http://localhost:8081/process?first_name=Nick" method="post" enctype="multipart/form-data">
     <input type="text" name="first_name"/>
     <input type="text" name="last_name"/>
     <input type="submit"/>
</form>
```

```
func main() {

	server := http.Server{
		Addr: "localhost:8081",
	}

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "form"+r.URL.Path) 因为在表单里面写了地址了，所以不用再写访问路径
		//r.ParseForm() //先解析表单
		r.ParseMultipartForm(1024) //此处指字节长度

		// fmt.Fprintln(w, r.PostForm["first_name"]) //map[first_name:[yivei] last_name:[shi]]
		// fmt.Fprintln(w, r.Form["first_name"])     // [yivei Nick]
		// fmt.Fprintln(w, r.PostForm["first_name"]) //[yivei]
		fmt.Fprintln(w, r.MultipartForm) // &{map[first_name:[yivei] last_name:[shi]] map[]}

	})
	server.ListenAndServe()
}
```

返回结果

```
&{map[first_name:[yivei] last_name:[shi]] map[]} //一个结构体中两个map
```

**FormValue和PostFormValue方法**

FormValue获取的是[]string的第一个值

PostFormValue用法一样，但是只能读取PostForm

**上传文件**

```
multipart/form-data //最常见的应用就是上传文件
```

```
<!--<form action="http://localhost:8081/process?first_name=Nick" method="post" enctype="multipart/form-data">-->
<form action="http://localhost:8081/process?hello=world&thread=123" method="post" enctype="multipart/form-data">
<!--     <input type="text" name="first_name"/>-->
<!--     <input type="text" name="last_name"/>-->
     <input type="text" name="hello" value="sau sheong"/>
     <input type="text" name="post" value="456"/>
<!--     //上传文件功能-->
     <input type="file" name="uploaded"/>
     <input type="submit"/>
</form>
```

```
func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}

	//http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
	//	//http.ServeFile(w, r, "form"+r.URL.Path) 因为在表单里面写了地址了，所以不用再写访问路径
	//	//r.ParseForm() //先解析表单
	//	//r.ParseMultipartForm(1024) //此处指字节长度
	//
	//	// fmt.Fprintln(w, r.PostForm["first_name"]) //map[first_name:[yivei] last_name:[shi]]
	//	// fmt.Fprintln(w, r.Form["first_name"])     // [yivei Nick]
	//	// fmt.Fprintln(w, r.PostForm["first_name"]) //[yivei]
	//	//fmt.Fprintln(w, r.MultipartForm) // &{map[first_name:[yivei] last_name:[shi]] map[]}
	//	//fmt.Fprintln(w, r.FormValue("first_name")) //不用解析，使用FormValue直接调用 //Nick
	//
	//})
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

//写一个handler专门处理上传的文件
func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024) //解析表单

	fileHandler := r.MultipartForm.File["uploaded"][0] //访问File字段
	file, err := fileHandler.Open()                    //获取文件

	//file, _, err := r.FormFile("uploaded")

	if err != nil {
		data, err := ioutil.ReadAll(file) //读取文件到byte切片里
		if err == nil {
			fmt.Fprintln(w, string(data)) //打印
		}
	}
}
```

Form方法，它更加的简单

```
func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(1024) //解析表单
	//
	//fileHandler := r.MultipartForm.File["uploaded"][0] //访问File字段
	//file, err := fileHandler.Open()                    //获取文件

	file, _, err := r.FormFile("uploaded") //该方法能够直接获取文件而不用解析（默认是获取slice的第一个值），比较适用只上传一个文件

	if err != nil {
		data, err := ioutil.ReadAll(file) //读取文件
		if err == nil {
			fmt.Fprintln(w, string(data)) //打印
		}
	}
}
```

Forms - MultipartReader（）

读取Form的值
Form，PostForm，FormValue()，PostFormValue(), FormFile ()，MultiPartReader()

```
func (r *Request) MultipartReader() (*multipart.Reader, error) {
	if r.MultipartForm == multipartByReader {
		return nil, errors.New("http: MultipartReader called twice")
	}
	if r.MultipartForm != nil {
		return nil, errors.New("http: multipart handled by ParseMultipartForm")
	}
	r.MultipartForm = multipartByReader
	return r.multipartReader(true)
}
```

如果是multipart/form-data或multipart 混合的POST请求

MultipartReader返回一个MIME multipart reader

否则返回nil 和一个错误

可以使用该函数来代替ParseMultipartForm来把请求的Body作为Stream进行处理

​	不是把表单作为一个对象来处理，不是一次性获得整个map

​	逐个检查来自表单的值，然后每次处理一个

**POST请求**

JSON Body 

不是所有的请求都来自HTML的Form

如客户端框架（如Angular等）会以不同的方式对POST 请求进行编码

jQuery通常使用application/x-www-form-urlencoded

Angular 是application/json

ParseForm 方法无法处理application/json

## 1.6 ResponseWriter

从服务器向客户端返回响应时需要使用ResponseWriter

它是一个接口，handler用来返回响应

```
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}
```

支撑ResponseWriter的幕后struct 是非导出的http.response

```
type ResponseWriter interface { //ResponseWriter 是一个接口，非导出的*response也实现了它里面的所有方法，所以*response实现了这个接口，也就是说这个接口实际上也是一个指针，都是按照引用进行传递的
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}
```

写入到ResponseWriter

**Write方法**

它里面的Write方法会写入到HTTP响应的Body里面

如果调用Write方法，但是header里未设定content type，那么数据的前512字节会用于检测content type

```
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/write", writerExample)
	server.ListenAndServe()
}

func writerExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head>Go Web</head> //head里面只写了内容
<body><h1>Hello World</h1></body> 
</html>`
	w.Write([]byte(str))
}

运行 
qinjianquan@MacBook-Pro-10 part6.response % go run .

使用curl访问

curl -i localhost:8080/write
HTTP/1.1 200 OK //使用的协议是HTTP/1.1
Date: Mon, 25 Apr 2022 04:24:58 GMT
Content-Length: 68
Content-Type: text/html; charset=utf-8 //推断出来的

<html>
<head>Go Web</head>
<body><h1>Hello World</h1></body>
</html>%                             
```

**WriterHeader方法**

参数是状态码，并将其作为响应的状态码返回

在调用Write方法之前会隐式的调用，如果显式调用，那主要用于返回错误

调用之后，仍然可以写入到ResponeWriter，但不可修改head

```
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	//http.HandleFunc("/write", writerExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	server.ListenAndServe()
}


func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "no such service, try next door")
}
```

执行上面的访问步骤，可以看到响应了body

```
qinjianquan@MacBook-Pro-10 ~ % curl -i localhost:8080/writeheader
HTTP/1.1 501 Not Implemented
Date: Mon, 25 Apr 2022 04:41:07 GMT
Content-Length: 31
Content-Type: text/plain; charset=utf-8

no such service, try next door
```

**Header方法**

返回的是headers 的map，可以进行修改，体现在返回给客户端的HTTP响应里

```go
func main() {
   server := http.Server{
      Addr: "localhost:8080",
   }
   //http.HandleFunc("/write", writerExample)
   //http.HandleFunc("/writeheader", writeHeaderExample)
   http.HandleFunc("/redirect", headerExample)
   server.ListenAndServe()
}

func headerExample(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("location", "https://www.google.com") //第一个参数是key，第二个参数是value
   w.WriteHeader(302)//注意此处应该放在修改header之后
}
```

```
curl -i localhost:8080/redirect   
HTTP/1.1 302 Found
Location: https://www.google.com
Date: Mon, 25 Apr 2022 04:49:51 GMT
Content-Length: 0
```

设置header content type 为 application/json

```
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	//http.HandleFunc("/write", writerExample)
	//http.HandleFunc("/writeheader", writeHeaderExample)
	//http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //设置content-type
	post := &Post{
		User:    "Sau sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post) //将struct转换为它的json编码，实际上是一个byte slice
	w.Write(json)                 //使用Write方法返回回去
}
```

```
qinjianquan@MacBook-Pro-10 part6.response % curl -i localhost:8080/json 
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 25 Apr 2022 05:15:33 GMT
Content-Length: 58

{"User":"Sau sheong","Threads":["first","second","third"]}%    
```

**内置response**

NotFound函数，包装了一个404状态码和一个额外的信息，告诉请求者请求的资源找不到了

ServeFile函数，从文件系统提供文件，返回给请求者

ServeContent函数，它可以把实现了io.ReaderSeeker接口的任何东西里面的内容返回给请求者，还可以请求资源的一部分，ServeFile和io.Copy则不行

Redirect函数，告诉客户端重定向到另一个URL

## 1.7 模版

### 1.7.1 **模版简介**

模版是预先设计好的HTML页面，可以被模版引擎反复使用，产生HTML页面

Go标准库提供了text/template，html/template两个模版库

​	大多数Go的web框架都使用这些库作为默认的模版引擎

**模版与模版引擎**

模版引擎可以合并模版与上下文数据，产生最终的HTML

**两种理想的模版引擎**

无逻辑模版引擎

​	通过占位符，动态数据被替换到模版中

​	不做逻辑处理，只做字符串替换

​	完全由handler处理

​	目标是展示层与逻辑的完全分离

逻辑嵌入模版引擎

​	编程语言被嵌入到模版中

​	在运行时，由模版引擎来执行，也包含替换功能

​	功能强大

​	逻辑代码遍布handler和模版，难以维护

**Go的模版引擎**	

主要使用的是text/template，HTML相关部分使用了html/template，是和混合体

模版可以完全无逻辑，但又具有足够的嵌入特性

和大多数模版引擎一样，Go Web的模版位于无逻辑和嵌入逻辑之间的某个地方

**Go模版引擎的工作原理**

在web应用中，通常是由handler来出发模版引擎

handler调用模版引擎，并将使用的模版传递给引擎（通常是一组模版文件和动态数据）

模版引擎生成HTML，并将其写入到ResponseWriter

ResponseWriter再将它加入到HTTP响应中

**关于模版**

模版必须是可读的文本格式，扩展名随意。对于web应用通常就是HTML

​	里面会嵌入一些命令（叫做action）

action位于{{.}}之间

​	.就是一个action

​	可以命令模版引擎将其替换成一个值

模版样例

```
<!DOCTYPE html>
<html lang="en">
<head>
     <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
     <title>Go Web Programming</title>
</head>
<body>
     {{.}} //action
</body>
</html>
```

**如何使用模版引擎**

1.解析模版源（可以是字符串或者模版文件）,从而创建一个解析好的模版的struct

2.执行这些解析好的模版，并传入ResponseWriter和数据

​	触发模版引擎组合解析好的模版和数据，来产生最终的HTML，并将它传递给ResponseWriter

```
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport",content="width=device-width, initial-scale=1.0">
    <title>Template</title>
</head>
<body>
{{ . }}
</body>
</html>
```

运行并在浏览器中访问

```
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html") //text/template库，解析模版
	t.Execute(w, "Hello World")              //执行传入ResponseWriter 和 上下文数据
}
```

浏览器返回Hello World

### 1.7.2 解析和执行模版

解析模版是三个函数：

ParseFiles：解析模版文件，并创建一个解析好的模版struct，后续可以被执行，这个函数是Template struct 上方法的简便调用，调用ParseFiles后，会创建一个新的模版，模版的名字是文件名

new函数会创建一个空的模版，名字就是传进去的名字

ParseFiles的参数可变，但是只返回一个模版

​	当解析多个文件时，第一个文件作为返回的模版（名，内容），其余的作为map，供后续执行使用	

```
func main() {
	//第一种方法
	//t,_ := template.ParseFiles("tmpl01.html") //调用解析函数，先把模版文件里面的内容读成字符串，调用模版上面的parse方法解析字符串
	//第二种方法
	t := template.New("tmpl01.html")
	t, _ = t.ParseFiles("tmpl01.html")
}
```

ParseGlob：

使用模式匹配来解析特定文件

```
t,_ := template.ParseGlob("*.html") //*是通配符，会找到所有文件名，模版名为第一个文件名
```

Parse

可以解析字符串模版，其它方式最终都会调用Parse

**Lookup方法**

通过模版名来寻找模版，如果没找到就返回nil，它是template上面的方法

**Must函数**

可以包裹一个函数，返回一个模版的指针和一个错误

​	如果错误不为nil，那么就panic

**执行模版**

两种方法，都是template上面的方法

```
func (t *Template) Execute(wr io.Writer, data interface{}) error //使用单模版，模版集：只会使用第一个模版
```

```
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error //适用于模版集
```

例如

```
func process(w http.ResponseWriter, r *http.Request) {

	//t, _ := template.ParseFiles("tmpl.html") //text/template库，解析模版
	//t.Execute(w, "Hello World")              //执行传入ResponseWriter 和 上下文数据

	//执行模版的两个方法
	t, _ := template.ParseFiles("t1.html")
	t.Execute(w, "Hello world")

	ts, _ := template.ParseFiles("t1.html", "t2.html")
	ts.ExecuteTemplate(w, "t2.html", "Hello World")

}
```

### 1.7.3 实操案例

```
func main() {

	//调用加载模版代码
	templates := loadTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Path[1:]      //截取后边的部分，获取文件名
		t := templates.Lookup(fileName) //通过文件名寻找模版
		if t != nil {
			err := t.Execute(w, nil) //执行模版
			if err != nil {
				log.Fatalln(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/img/", http.FileServer(http.Dir("wwwroot")))

	http.ListenAndServe("localhost:8080", nil)

}
```

```
//解析模版
func loadTemplates() *template.Template {
	result := template.New("templates")                 //创建一个空的模版
	template.Must(result.ParseGlob("templates/*.html")) //程序刚开始启动时就应该加载，失败的话终止程序,模版名就是文件名
	return result
}
```

### 1.7.4 Action

Action是Go模版中嵌入的命令，. 就是一个比较重要的action，代表了传入模版的数据

Action分为五类：

​	**条件类**

```
<body>
{{ if.}}
Number is greater than 5!
{{ else}}
Number is 5 or less!
{{end}}
</body>
```

```
func process(w http.ResponseWriter, r *http.Request) {

	//Action
	t, _ := template.ParseFiles("tmpl02.html")
	rand.Seed(time.Now().Unix())    //从1970年1月日到现在的秒数
	t.Execute(w, rand.Intn(10) > 5) //生成随机数的范围是0-9，判断是否大于5
}
```

​	**迭代/ 遍历类**

```
 <!--// . 是传入模版的数据-->
    {{ range.}}
    <!--// . 是 遍历的元素-->
    <li> {{.}}</li>
    <!--    //回落机制-->
    {{else}}
    <li>Nothing to show</li>
    {{end}}
```

​	**设置类的Ation**

```
 <div> The dot is {{.}} </div>
    <div>
        {{with "world"}}
        Now the dot is set to {{.}}
        {{end}}
    </div>
    <div>The dot is {{.}} again</div>
```

```
//Action
	//设置
	t, _ := template.ParseFiles("tmpl04.html")
	t.Execute(w, "hello")
```

​	回落机制

```
<div> The dot is {{.}} </div>
    <div>
         <!-- {{with "world"}}
        Now the dot is set to {{ . }}
        {{ end }} -->
        {{with ""}}
        Now the dot is set to {{ . }}
        {{else}}
        The dot is still {{ . }}
        {{ end }}
    </div>
    <div>The dot is {{.}} again</div>
```

​	**包含类的Action**

main函数中的内容以及两个html文件

```
//包含
	t,_ := template.ParseFiles("tmpl05.html","tmpl06.html")
	t.Execute(w, "hello world")
```

```
<body>
   <div>This is tmpl05.html</div>
   <div>This is the value of the dot in tmpl05.html - [{{.}}]</div>
   <hr/>
   {{template "tmpl06.html"}}
   <hr/>
   <div>This is the html05.html after</div>
</body>
```

```
<div style="background-color: yellow;">
    This is tmpl06.html<br>
    This is the value of the dot in tmpl06.html - [{{.}}]
</div>
```

浏览器结果

```
This is tmpl05.html
This is the value of the dot in tmpl05.html - [hello world]
This is tmpl06.html
This is the value of the dot in tmpl06.html - []
This is the html05.html after
```

```
  <div>This is tmpl05.html</div>
   <div>This is the value of the dot in tmpl05.html - [{{.}}]</div>
   <hr/>
   {{template "tmpl06.html" . }} //传入内容
   <hr/>
   <div>This is the html05.html after</div>
```

输出结果

```
This is tmpl05.html
This is the value of the dot in tmpl05.html - [hello world]
This is tmpl06.html
This is the value of the dot in tmpl06.html - [hello world]
This is the html05.html after
```

**定义类的Action**

### 1.7.5 函数与管道

参数；就是在模版里用到的值

​	可以是bool，整数...

​	也可以是struct，struct的字段，数组的key等

​	其他

**在Action中设置变量**

变量以$开头

​	$variable := value

遍历的时候就可能涉及设置变量

**管道**

是按顺序连接到一起的参数、函数和方法

​	和Unix的管道类似

如：{{p1|p2|p3}},pi要么是参数要么是函数

管道允许我们把参数的输出发给下一个参数，下一个参数由管道（｜）分隔开

```
{{12.3456 | printf "%.2f"}}
```

输出

12.35

**函数**

参数可以是一个函数

Go模版引擎提供了一些基本内置函数，但是功能有限，如fmt.Sprint的各类变体等

但是开发者可以自定义函数，要求是：

​	输入的参数的数量可以任意多个

​	返回值只能有一个或者一个返回值+一个错误

内置函数

define、template、block

html \ js \urlquery :主要用于转义字符串，防止安全问题

但是如果使用的是html/template这个库（是Web模版），基本用不上，因为其会自动对数据进行转义

index；后面一个集合+一个索引，代表集合中的某个元素

print/printf/println：是fmt包S开头的几个函数的变体

len 求集合长度

with 设置

**如何编写自定义函数**

第一步：创建一个FuncMap（map类型）

​	key是函数名

​	value是函数

第二步：把Func附加到模版

```
func process(w http.ResponseWriter, r *http.Request) {
//编写自定义的函数
	funcMap := template.FuncMap{"fdate": formatDate} //放入FuncMap
	t := template.New("tmpl08.html").Funcs(funcMap)  //创建新模版
	t.ParseFiles("tmpl08.html")                      //解析
	t.Execute(w, time.Now())//执行，操作顺序非常重要
}
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

```

```
<body>
    {{ . | fdate}} //等同于 {{fdate .}} //但是管道更加强大和灵活，建议使用管道
</body>
```

### 1.7.6 模版组合

jsp、php中都一个类似的概念就是layout模版

它是网页中的固定的部分，它可以被多个网页重复使用

**如何制作layout模版**

layout模版的初衷是保持一部分不变，其它的可变

正确做法，专门建一个layout模版，里面包含content页面

```
<!DOCTYPE html>
<html lang="en">
<head>
     <meta charset="UTF-8">
     <meta http-equiv="X-UA-Compatible" content="IE=edge">
     <meta name="viewport" content="width=device-width, initial-scale=1.0">
     <title>Layout</title>
     <link herf="https://cdn.bootcdn.net/ajax/libs/twitter-bootstrap/4.5.3/css/bootstrap.min.css"> //添加固定的部分
</head>
<body>
     <nav class="navbar navbar-dark bg-primary">//添加固定的部分
         <a class="navbar-brand" herf="#">Navbar</a>
     </nav>
    {{ template "content.html" .}}  //每个模版都需要提供content.html来定自己的内容,但是问题在要让每个内容文件都同名，所以我们要定义
</body>
</html>
```

重定义某个内容模版的模版名为content.html

```
{{define "content.html"}}//这里的模版名只要和公共模版中的名字对应上即可，可以是任何字符

<h1>Home</h1>

{{end}}
```

```
{{define "layout"}} //把layout也改一下

<!DOCTYPE html>
<html lang="en">
<head>
     <meta charset="UTF-8">
     <meta http-equiv="X-UA-Compatible" content="IE=edge">
     <meta name="viewport" content="width=device-width, initial-scale=1.0">
     <title>Layout</title>
     <link herf="https://cdn.bootcdn.net/ajax/libs/twitter-bootstrap/4.5.3/css/bootstrap.min.css">
</head>
<body>
     <nav class="navbar navbar-dark bg-primary">
         <a class="navbar-brand" herf="#">Navbar</a>
     </nav>
    {{ template "content.html" .}} 
</body>
</html>

{{end}}
```

运行，浏览器可以正常输出内容

```
Navbar
Home
Hello World
```

再增加一个About页面

```
{{define "content"}}

<h1>About</h1>

{{end}}
```

```
func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("layout.html", "home.html")
		if err != nil {
			log.Panic("解析失败", err)
		}
		t.ExecuteTemplate(w, "layout", "Hello World")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("layout.html", "about.html")
		if err != nil {
			log.Panic("解析失败", err)
		}
		t.ExecuteTemplate(w, "layout", "")
	})

	http.ListenAndServe("localhost:8080", nil)
}
```

浏览器访问结果

```
Navbar
About
```

blockAction 后的文件可以不存在，template必须存在

改为block

```
 <!-- {{ template "content" .}} -->
    {{block "content" .}}
    test
    {{end}}
```

```
http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("layout.html")
		if err != nil {
			log.Panic("解析失败", err)
		}
		err = t.ExecuteTemplate(w, "layout", "")
		log.Println(err)
	})
```

```
Navbar //浏览器返回内容
test
```

**逻辑运算符**

eq/ne 相等或者不相等

It/gt 小于或大于

le/ge 小于等于或大于等于

and or not 与或非

```
{{ define "content" }}

<h1>Home</h1>

<h2>{{if eq . "Hello World"}}
     !!!
     {{ end }}
</h2>

{{end}}
```

## 2 Go数据库编程

### 2.1 链接SQL数据库

#### **2.1.1 sql.DB简介**

想要连接到SQL数据库，首先需要加载目标数据库的驱动，驱动里面包括了与该数据库的交互逻辑

使用sql.Open(),它会得到一个指向sql.DB这个struct的指针

sql.DB用来操作数据库，它代表0个或者多个底层连接的池，这些连接由sql包来维护，sql包会自动创建和释放这些连接（无需手动关闭）

​	goroutine可以并发安全使用它

注意：Open() 函数并不会连接到数据库，甚至不会验证参数，它只是把后续连接到数据库所必需的structs给设置好了

而真正的连接是在被需要的时候才进行懒设置的

sql.DB是用来处理数据库的而不是实际的连接

这个抽象包含了数据库连接的池，而且会对此进行维护

使用时，可以定义它的全局变量进行使用，也可以将它传递到函数或者方法里

#### 2.1.2 **在docker中启用sql数据库**

在使用Go连接数据库之前，需要确保本地的数据库已经注册并启用，为了节省步骤，在此不直接在本地安装数据库，而是在docker中启用一个数据库镜像，以下流程都是在mac上执行

1.下载安装docker

https://www.docker.com/get-started/

2.启动docker

```
docker run -d -p 80:80 docker/getting-started
```

启动成功显示如下

```
4b53e2ebd76b   docker/getting-started           "/docker-entrypoint.…"   43 minutes ago   Up 43 minutes   0.0.0.0:80->80/tcp       sweet_bhaskara
```

3.注册并启动数据库

为了简便起见，启用数据库的时候只需设置连接密码即可，用户名默认为sa

```
docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=15Willis," -p 1433:1433 --name sql1 --hostname sql1 -d mcr.microsoft.com/mssql/server
```

启动成功显示如下

```
6ed090d89ae3   mcr.microsoft.com/mssql/server   "/opt/mssql/bin/perm…"   34 minutes ago   Up 34 minutes   0.0.0.0:1433->1433/tcp   sql1
```

在容器内连接

```
docker exec -it 5cafffe65f4a /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P 15Willis, 
1> //出现此标志符说明连接成功
```

注意：在连接数据库之前请务必保证本地的数据库已经启动，无论是直接下载启动数据库还是在docker中启动一个数据库镜像容器，此操作不执行会导致在golang中无法连接数据库

#### 2.1.3 在go中连接数据库

```
import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb" //微软sqlserver数据库的驱动，引入时它的init函数进行了自我注册
	"log"
)

var db *sql.DB //

const (
	server   = "localhost"
	port     = 1433
	user     = "sa"
	password = "xxxxx" //使用上一节中的密码
	database = "master" //使用默认的master即可
)

func main() {
	//数据库的连接字符串
	connStr := fmt.Sprintf("server=%s;user id =%s; password=%s;port=%d;database=%s;", server, user, password, port, database)
	db, err := sql.Open("sqlserver", connStr) //open函数的两个参数分别是数据库驱动的名称以及连接字符串，返回的第一个值是db
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background() //返回一个非空的Context，不会被取消，没有值也没有截止时间
	
	err = db.PingContext(ctx) //使用db ping数据库
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Connected")
}
```

### 2.2 单笔查询

**sql.DB上的方法**

Query：把查询语句发送到数据库，结果是0行或多行

QueryRow：只能返回一条记录

QueryContext：可以带上下文，如取消、截止时间等

QueryRowContext：

Query 和 QueryContext 返回类型都是：type Rows struct{},此类型还有一些方法

QueryRowContext 和 QueryRow返回类型一样

##  3 路由

前面针对不同的路径触发不同的handler就是一种路由。

main函数主要用于设置类的工作

controller：静态资源、把不同的请求送到不同的controller进行处理

路由结构：前置controller 后面有若干个handler

把每个handler都注册成路由

```
func registerHomeRouts() {
	http.HandleFunc("/home", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("layout.html", "home.html")
	if err != nil {
		log.Panic("解析失败", err)
	}
	t.ExecuteTemplate(w, "layout", "Hello World")
}
...
```

```
func RegisterRouts() {

	//static resources

	registerAboutRouts()
	registerContactRouts()
	registerHomeRouts()
}
```

然后在main函数中调用，此时路径参数的传递应该是隐式的

```
func main() {
   controller.RegisterRouts()
   http.ListenAndServe("localhost:8080", nil)
}
```

上述都是静态路由，也就是一个路径对应一个页面

**带参数的路由**

根据路由参数，创建一组不同的页面，路径相同，会走更接近的路径

```
func registerCompanyRouts() {
	http.HandleFunc("/companies", handleCompanies)
	http.HandleFunc("/companies/", handleCompany)
}

func handleCompanies(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "companies.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func handleCompany(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "company.html")
	pattern, _ := regexp.Compile(`/companies/(\d+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		fmt.Println(matches[0])
		companyID, _ := strconv.Atoi(matches[1])
		t.ExecuteTemplate(w, "layout", companyID)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
```

**第三方路由器**

gorilla/mux：灵活性更高、功能强大、性能相对差一点

httprouter：注重性能、功能简单

当然可以根据业务需求编写自己的路由规则

## 4.JSON

 Go struct 和json可以相互映射

类型也可以映射

读取json，解码器和编码器

如下是先解码，再编码

```
POST http://localhost:8080/companies HTTP/1.1
content-type: application/json

{
     "id" :123,
     "name" : "Google",
     "country": "USA"
}
```

```
func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			dec := json.NewDecoder(r.Body)
			company := Company{}
			err := dec.Decode(&company)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			enc := json.NewEncoder(w)
			err = enc.Encode(company)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe("localhost:8080", nil)
}
```

```
HTTP/1.1 200 OK
Date: Wed, 27 Apr 2022 07:21:47 GMT
Content-Length: 43
Content-Type: text/plain; charset=utf-8
Connection: close

{
  "id": 123,
  "name": "Google",
  "country": "USA"
}
```

还有一种方式marshal 和 unmarshal

```
func main() {

	jsonStr := `
	{
		"id" :123,
		"name" : "Google",
		"country": "USA"
	}`
	c := Company{}
	_ = json.Unmarshal([]byte(jsonStr), &c)
	fmt.Println(c)

	bytes, _ := json.Marshal(c)
	fmt.Println(string(bytes))

	bytes1, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(bytes1))

	http.ListenAndServe("localhost:8080", nil)
}
```

输出

```
{123 Google USA}
{"id":123,"name":"Google","country":"USA"}
{
  "id": 123,
  "name": "Google",
  "country": "USA"
}
```

两种方式区别，主要是str 和bytes的区别

## 5.中间件

```
//链式结构

type MyMiddleWare struct {
	Next http.Handler
}

func (m MyMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//do something
}
```

用途：所做日志，身份认证，请求超时，响应压缩

```
type AuthMiddleware struct {
	Next http.Handler
}

func (am AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if am.Next == nil {
		am.Next = http.DefaultServeMux
	}

	auth := r.Header.Get("Authorization")
	if auth != "" {
		am.Next.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
```

```
func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			ID:      123,
			Name:    "Google",
			Country: "USA",
		}

		enc := json.NewEncoder(w)
		enc.Encode(c)
	})
	http.ListenAndServe("localhost:8080", new(middleware.AuthMiddleware))
}
```

运行结果，使用两个测试来运行，一个能返回结果，一个未授权

```
### Without Auth
POST http://localhost:8080/companies HTTP/1.1


### With Auth
POST http://localhost:8080/companies HTTP/1.1
Authorization:abcd
```

## 6.请求上下文

```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

```
type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx := r.Context()

	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	r.WithContext(ctx)

	ch := make(chan struct{})
	go func() {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
```

```
func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			ID:      123,
			Name:    "Google",
			Country: "USA",
		}
		time.Sleep(4 * time.Second)

		enc := json.NewEncoder(w)
		enc.Encode(c)
	})
	http.ListenAndServe("localhost:8080", &middleware.TimeoutMiddleware{Next: new(middleware.AuthMiddleware)})
}
```

## 7 HTTPS

HTTPS 不是直接在传输层上面传输数据的，而是会添加一个层，TLS

先生成两个文件

```
qinjianquan@MacBook-Pro-10 part12.https % go run  /usr/local/go/src/crypto/tls/generate_cert.go -host localhost

2022/04/27 17:13:12 wrote cert.pem
2022/04/27 17:13:12 wrote key.pem
```

```
func main() {

	controller.RegisterRouts()
	http.ListenAndServeTLS("localhost:8080", "cert.pem", "key.pem", nil) //传入生成文件
}
```

现在就可以使用https访问了

Http2 

允许多路复用

允许对header进行压缩

协议默认安全

Server push 减少了请求和页面加载

## 8.测试

注意：测试的本质就是找一组特征值，看返回的结果是否正确

测试文件以_test.go 结尾

测试函数需要以Test开头

```
type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func (c *Company) GetCompanyType() (result string) {
	if strings.HasSuffix(c.Name, ".LTD") {
		result = "Limited Liability Company"
	} else {
		result = "others"
	}
	return
}
```

```
func TestGetCompanyTypeCorrect(t *testing.T) {
	c := Company{ID: 12345,
		Name:    "ABCD.LTD",
		Country: "China",
	}

	companyType := c.GetCompanyType()
	fmt.Println(companyType)

	if companyType != "Limited Liability Company" {
		t.Errorf("Test failed")
	}
}
```

## 9.性能分析

内存消耗情况
CPU的使用情况

阻塞的Goroutine

执行追踪
Heap/profile/goroutine...

比入监听网站访问、对数据库进行增删改查等