# Boy framework

[English](https://github.com/slclub/boy/blob/master/README.md) [中文](https://github.com/slclub/boy/blob/master/README.ch.md)

## 概述

一个轻型的go web 框架。80% 的代码都是用接口实现的。支持自定义执行节点 以及中间件。

这个Boy 框架 像是是一个demo，仅仅是定义一些别名简写方便使用，定义了一些执行节点。

您也可以迅速的用它重写一个属于自己的框架。可重可轻，不过需要知道相关的几个包都是

做什么的，以及如何使用他们。



通过ini配置文件我们可以 同时监听http 和 https 甚至是websocket, 以及静态服务配置等

比较工程化。


通过这些年的工作，一个可伸缩的框架才是我理想中的框架。所以我们此框架最终主要为了

以下几点：

1 伸缩性：框架具有伸缩性，可重，可轻，满足不同的需求。

2 可变：框架可以根据不同配置参数适当的变化运行方式，不需要再去改框架代码。减少依赖和修改源码。

3 简洁:更为简洁统一的写法，很多时候不必关系内部处理。

4 最小重写：当重写时尽可能的少写代码，仅改需要改的部分即可。

## 安装

你需要先安装go 环境。暂时还没有做docker 支持

Go 版本 1.14+

```ssh
go get -u github.com/slclub/boy
```

如果您使用go mod 管理您的项目。那么仅仅需要import此包即可

## 快速开始

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
