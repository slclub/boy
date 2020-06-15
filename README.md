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
- [Router](#router)
- [gnet.Context](https://github.com/slclub/gnet)
	- Request
	- Response

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

### Custome

- Rewrite router

 Example : Just implement the [github.com/slclub/gouter.Router](https://github.com/slclub/grouter), if you want to rewrite router.
