# boy

### summary

A light go web framework. Most of the functional components are implemented by interfaces.

Support custom execution nodes and Middleware. 

Boy framework like and demo, Defined some alias global variables,  easy to use.

You can as soo as possible organize a smaller or bigger framework using the corresponding packages.

The ultimate goal of writing this:

1 The framework can be as large or as small as possible to customize and modify most components.
	
2 Highly configurable. You can easily deploy multiple services on one machine. Independent of system environment variables

### Content List
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Benchmarks](#benchmarks)
- [Configration](#configration)
- [Router](#router)
- [Context](#context)
	- [Request](#request)
		- [Request Paramters](#request-paramters)
	- [Response](#response)
	- [Contexter Api](#contexter-api)
		- [Contexter Setter Getter](#contexter-setter-getter)
		- [Contexter Get Set Request](#contexter-get-set-request)
		- [Contexter Get Set Response](#contexter-get-set-response)
		- [Contexter Abort](#contexter-abort) process flow control.
	- [Save File](#save-file)
- [MidderWare](#middlerware)
- [Static](#static)
- [Custom](#custom)

### Installation

To install boy framework, you should install Go first.

Go required 1.14+

```ssh
go get -u github.com/slclub/boy
```

If you use "go mod" to manager your project. just import the packeage.

```go
import "github.com/slclub/boy"
```

If you want to use the latest code. please change to latest version in the go.mod file.

```go
require (
	github.com/slclub/boy latest
)
```

### Quick start

```go
package main                                                                                                                                                                                                                    

import (
    "github.com/slclub/boy"
    "github.com/slclub/gnet"
)

func main() {
	boy.R.GET("/example/ping", func(this gnet.Contexter) {
        this.Response().WriteString("Hello World!")
    }) 
    boy.Run()
}

```

### Benchmarks

```go
BenchmarkServer-4		 1000000	      1012 ns/op	     528 B/op	       7 allocs/op
PASS
ok		github.com/slclub/gcore	1.029s
```

### Configration

Install from [boy.Install()](https://github.com/slclub/boy)

[Config file template](https://github.com/slclub/boy/tree/master/etc_temp)

If there is no config file. It will automaticlly generate one.

You don't have to worry about it.

### Router

Import from [grouter](https://github.com/slclub/grouter)

- Restfull router added.
```go
	boy.R.

	GET(url string, ctx gnet.HandleFunc)
	POST(url string, ctx gnet.HandleFunc)
	PUT(url string, ctx gnet.HandleFunc)
	DELETE(url string, ctx gnet.HandleFunc)
	PATCH(url string, ctx gnet.HandleFunc)
	HEAD(url string, ctx gnet.HandleFunc)
	OPTIONS(url string, ctx gnet.HandleFunc)
	// can accept any http.Method request.
	ANY(url string, ctx gnet.HandleFunc)
```

-  Use Group

```go
	// 
	boy.R.Group(func(group grouter.IGroup){
		// add this group routes a middlerware.
		group.Use(func(ctx gnet.Contexter) {})
		// Deny a middlerware for these routes.
		group.Deny(gnet.HandleFunc)
		boy.R.GET(url string, ctx gnet.HandleFunc)
		boy.R.POST(url string, ctx gnet.HandleFunc)
		...
	})
```

- Use Defined code handle, 404, 405,500;

```go
	boy.R.BindCodeHandle(404, gnet.HandleFunc)
	boy.R.BindCodeHandle(405, gnet.HandleFunc)
	...
```

- [group.Router interface](https://github.com/slclub/grouter/blob/master/irouter.go)

### Context

[gnet.Contexter](https://github.com/slclub/gnet)

#### Request
	
Request object. Obtain all kinds of requested information and parameter routes in the object here

[Source Code](https://github.com/slclub/gnet/blob/master/request.go), You can read it from this link.


##### Request Paramters

You can use these methods get paramters from path param, query, form, and so on.

You don't care about  where the paramters come from.

```go
type RequestParameter interface {
	// Get param inetface ----------------------------------------------------------
	// just get string by key string.
	// q=a
	// return a
	GetString(key string, args ...string) (value string, ret bool)
	// q[]=a&q[]=b
	// return []string{a, b}
	GetArray(key string, args ...[]string) ([]string, bool)
	// q[a]=a&q[b]=b
	// return map[string]string{"a":"a", "b":"b"}
	GetMapString(key string, args ...map[string]string) (map[string]string, bool)
	GetInt64(key string, args ...int64) (int64, bool)
	GetInt(key string, args ...int) (int, bool)

	// set
	SetParam(key, value string)

	// input :// data
	BodyByte() ([]byte, error)
}
```

```go
	// example
	boy.R.GET(url, func(ctx gnet.Contexter){
		s1, ok := ctx.Request().GetString(key, default_value string)
	})
```

##### Request Interface

```go
type IRequest interface {
	GetHttpRequest() *http.Request
	RequestParameter

	// init and reset
	InitWithHttp(*http.Request)
	Reset()

	// header
	GetHeader(key string) string
	ContentType(args ...bool) string
	GetRemoteAddr() string
	//file
	FormFile(key string) (*multipart.FileHeader, error)
}
```

```go
	// example
	boy.R.GET(url, func(ctx gnet.Contexter){
		s1 := ctx.Request().ContentType()
	})
```


#### Response

Rewrite the interface of http.ResponseWriter. 

```go
type IResponse interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier

	// Returns the HTTP response status code of the current request.
	Status() int

	// Returns the number of bytes already written into the response http body.
	// See Written()
	Size() int

	// Writes the string into the response body.
	WriteString(string) (int, error)

	// Returns true if the response body was already written.
	Written() bool

	// Forces to write the http header (status code + headers).
	FlushHeader()

	// get the http.Pusher for server push
	Pusher() http.Pusher

	// init reset
	Reset()
	InitSelf(http.ResponseWriter)

	// update ResponseWriter.Header()
	Headers(key, value string)
}

```

```go
	// example
	boy.R.GET(url, func(ctx gnet.Contexter){
		ctx.Response().Write([]bytes{"hello girls"})
		ctx.Response().WriteString("hello girls")
	})

```

#### Contexter Api

[Source Code](https://github.com/slclub/gnet/blob/master/context.go)

These methods been used in the same way.

```go
    // example
    boy.R.GET(url, func(ctx gnet.Contexter){
        ctx.Xxx(args ...)
    })

```


```go
	Reset()
	// gerror.StackError([]error) support:
	// Push(err error)
	// Pop() error
	// Size() int
	GetStackError() gerror.StackError
	SetSameSite(st http.SameSite)

	ClientIP() string
	//
	GetHandler() HandleFunc
	SetHandler(HandleFunc)

	GetExecute() Executer
	SetExecute(exe Executer)

	//redirect
	Redirect(location string, args ...int)
```

##### Contexter Setter Getter

For custome key-value pairs extension.

```go
type SetterGetter interface {
	// setter
	Set(key string, val interface{})
	// getter
	Get(key string) (interface{}, bool)
	GetString(key string) string
	GetInt(key string) int
	GetInt64(key string) int64
}
```

##### Contexter Get Set Request 

```go
// request from other place.
type IContextRequest interface {
	Request() IRequest
	SetRequest(IRequest) bool

	// cookie
	SetCookie(name, value string, args ...interface{})
	Cookie(string) (string, error)
}
```

##### Contexter Get Set Response

```go
// response to client or other server.
type IContextResponse interface {
	Response() IResponse
	SetResponse(IResponse) bool
}
```

##### Contexter Abort

Execution process jump control.

```go
// interrupt interface.
type Aborter interface {
	// abort current handle .
	Abort()
	AbortStatus(int)
	// Jump out of the whole execution process.
	// break whole work flow.
	Exit()
}
```

#### Save File

Upload file.

```go
f1 = func(ctx gnet.Contexter) {
    f, err := ctx.Request().FormFile("file")
    gnet.SaveUploadFile(f, "/tmp/glog/test")
}
```

### MiddlerWare

[Source Code](https://github.com/slclub/gcore/blob/master/execute/middleware.go). You can use or deny any flow node or url handle middlerware.

No matter where the middleware is used.

- interface

```go
type Middler interface {
	// public excuter interface.
	flow.IExecuteNode
	// middle ware interface
	Use(gnet.HandleFunc)
	Deny(gnet.HandleFunc)

	GetHandle(i int) (gnet.HandleFunc, string)
	Combine(Middler)
	Size() int
}
```
- Use Deny
```go
	example:
	mf1 := func(ctx gnet.Contexter){}
	mf2 := func(ctx gnet.Contexter){}
	mf3 := func(ctx gnet.Contexter){}

	// public before node use or deny middlerware.
	boy.MiddlerBefore.Use(mf1)
	boy.MiddlerBefore.Use(mf2)

	// The url1 handle will only execute the first gnet.HandleFunc. 
	// flow: mf1(ctx), handle
	boy.R.Deny(mf2)
	boy.R.GET(url1, func(ctx gnet.Contexter){})

	// The url2 will execute boths handles.
	// flow:  mf1(ctx), mf2(ctx), mf3(ctx) handle
	boy.R.Use(mf3)
	boy.R.GET(url2, func...)

	// after node use or deny middlerware.
	boy.MiddlerAfter.Use ...

```

- Group Router Use Deny

Router Group use or deny middlerware.

```go
	example:
    mf1 := func(ctx gnet.Contexter){}
    mf2 := func(ctx gnet.Contexter){}
	mf3 := func(ctx gnet.Contexter){}

	boy.MiddlerBefore.Use(f1)
	boy.MiddlerBefore.Use(f2)

	// These routes of group have to same flow way.
	// flow: mf1(ctx), mf3(ctx) , handle, mf5(ctx)
	boy.R.Group(func(group grouter.IGroup){
		group.Deny(mf2)
		group.Use(mf3)
		group.Deny(mf4)
		boy.R.GET(url1, gnet.HandleFunc)
		boy.R.GET(url2, gnet.HandleFunc)
		boy.R.GET(url3, gnet.HandleFunc)
	})

    mf4 := func(ctx gnet.Contexter){}
	mf5 := func(ctx gnet.Contexter){}

	boy.MiddlerAfter.Use(f4)
	boy.MiddlerAfter.Use(f5)


```

### Static

Static service listening

You just need to change the file of etc/go.ini

```ini
# static service setting.
[static_service]
# Static root path.
# Default value is empty
# If you set the value of this field. It should be an absoluted path.
root=

# example: 
# Listening to multiple static folders needs to be separated by backspaces
service=sa  sb  
# param     @1  aliase of static path that is used to url.
#           @2  actual floder path.
#           @3  folder listing. true,on,yes,false; Where can browse directories.
sa=source  source      true
sb=static  assets     true
```

### Custom


- Rewrite router

 Example : Just implement the [github.com/slclub/gouter.Router](https://github.com/slclub/grouter), if you want to rewrite router.

```go
// examples:

func NewRouter() grouter.Router {

    r := &router{}
    r.initself(grouter.NewRouter())
    //r.SetStore(grouter.NewStore())
    //r.SetDecoder(grouter.NewPath())
    //r.code_handles = make(map[int]gnet.HandleFunc)
    //bind code handle
    r.BindCodeHandle(http.StatusNotFound, func(ctx gnet.Contexter) {
        ctx.Response().WriteHeader(404)
        ctx.Response().WriteString("grouter 404 not found")
    })  
    r.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(404)
        w.Write([]byte("grouter 404 not found"))
    }   
    //http.StatusMethodNotAllowed
    //r.BindCodeHandle(http.StatusMethodNotAllowed, grouter.http_405_handle)
    //r.BindCodeHandle(http.StatusInternalServerError, grouter.http_500_handle)
    return r
}

type router struct {
    grouter.Router
    NotFoundHandler func(http.ResponseWriter, *http.Request)
}

func (r *router) initself(rr grouter.Router) {
    r.Router = rr
}



router := NewRouter()
router.SetKey("router")
boy.App.DriverRegister(router)

```
