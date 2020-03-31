# go module简介

1. 一个module(模块)是一个 go的包集合，这个集合存储在一个`go.mod`为root的文件树中。go.mod文件定义了这个模块的路径（该模块的import路径）和该模块的依赖。每一个依赖由一个模型路径和一个版本号组成。

> A module is a collection of [Go packages](https://golang.org/ref/spec#Packages) stored in a file tree with a `go.mod` file at its root. 

2. go1.11和go1.12提供了gomodule的初步版本，在go1.11中，如果当前目录或者父目录中有一个`go.mod`文件，go命令启用go module（这里目录需要在`$GOPATH/src`之外。如果在`$GOPATH/src`之内，为了兼容性，依然会使用老的GOPATH模式，即使有`go.mod`文件）。从go1.13开始，module 是默认模式。golang截止20190823的最新稳定版是go1.12.9，在2019/08/15发布。上一个稳定版本是go1.11.13，go1.11在2018/08/24发布，go1.11.13在2019/08/13发布（可见每个主版本后面都跟着其他的bug修复版本）。这里选择使用go1.11.13。
3. 

# go module常见操作

## 启用go module

1. `export GO111MODULE=on`: go 命令会使用go module,不会去`$GOPATH/src`下面查找。另外还有两个值：
   - `export GO111MODULE=off`: go 命令不使用go module功能
   - `export GO111MODULE=auto`: 默认值。当前目录在`GOPATH/src`之外且该目录包含go.mod文件；或者当前文件在包含go.mod文件的目录下面，会使用go module。否则使用原来的`GOPATH/src`模式。
2. 当modules 功能启用时，依赖包的存放位置变更为`$GOPATH/pkg`，允许同一个package多个版本并存，且多个项目可以共享缓存的 module。
   1. 同一个major版本，是否能有多个minor版本呢？

## 新建一个module

1. `go mod init example.com/hello`会生成一个go.mod文件，里面第一行是`module example.com/hello`
2. `go.mod`文件只存在于module的跟目录，如果我们一个子目录world，那么该package的引用路径是：`example.com/hello/world`
3. go mod默认使用go.mod文件中列举的版本，如果go.mod中没有列举，go自动找该模块并加入到go.mod中, 使用最新稳定版本:

> using the latest version. (“Latest” is defined as the latest tagged stable (non-[prerelease](https://semver.org/#spec-item-9)) version, or else the latest tagged prerelease version, or else the latest untagged version.)

4. go mod下载的包会在`$GOPATH/pkg/mod`目录。

## 显示依赖

4. `go list -m all`可以查看当前模块的所有依赖。go会自动
5. 出了`go.mod`文件，go还维护了以`go.sum`文件，记录了每一个依赖包的哈希。`go.mod`和`go.sum`都要被加入版本管理系统。
6. 可以通过：`go list -m -versions rsc.io/sampler`查看`rsc.io/sampler`所有的有tag的版本。直接go get会升级到一个最新的tag版本，可以通过`go get rsc.io/sampler@v1.3.1`来指定版本升级。

## 添加依赖

1. `go build`, `go test`这些编译命令会自动往`go.mod`中添加依赖。

## 升级依赖

1. go module的版本号有三个部分: major.minor.patch。
2. 对于minor版本，升级只需要go get对应的包即可，会自动升级到`the latest tagged version`。go.mod中的`// indirect`表示该某块不是直接被当前模块引用。
3. 对于major版本，每一个主版本使用的 module path不同，**从v2开始**，这样我们可以使用一个包的不同版本,如

```go
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
```

相应的使用的时候：

```go
import (
    "rsc.io/quote"
    quoteV3 "rsc.io/quote/v3"
)
```



这样假定同major版本的是api兼容的，不同major版本的是api不兼容的，所以使用不同的module path。go对同一个module path只允许一个版本，即以每个module的major版本区分。这样版本升级的时候可以实现可控增量升级。

## 移除依赖

1. `go mod tidy`可以移除go.mod中没有被使用的依赖。

# 常见问题

## cannot find module for path github.com/astaxie/beego

原因：beego 官方估计还没有使用go mod。

解决办法：`go get github.com/astaxie/beego`

## 通过公共代理下载包

1. go mod包下载代理：

- [go mod代理和小技巧](https://www.cnblogs.com/xdao/p/go_mod.html)
  - Linux和windows git bash

```shell
export GOPROXY="https://athens.azurefd.net"

```

```
- Windows 还需要按照上面的设置GOPROXY和GO111MODULE环境变量
```

- <https://stackoverflow.com/a/10385612> 如何设置代理

## 修改了公共库

1. 修改了公共库之后，如果不想提交到官方代码，不想自己搭建一个go mod原，那么需要使用replace语句，将改库定向到本地，并将修改了的库，加入本仓库。如：

```go
golang.org/x/net/http/httputil => ./golang.org/x/net/http/httputil
```

## 公共库代码在哪里

1. go mod的代码cache在`$GOPATH/pkg/mod`, windows上升级go后，GOPATH默认是`~/go`

## 公共库代码依赖被强了怎么办

1. 找到`$GOPATH/pkg/mod`下的公共库, 修改其go.mod文件。指向github.com

## go get git.apache.org/thrift.git/lib/go/thrift@v0.9.3: unexpected status (https://goproxy.io/git.apache.org/thrift.git/lib/go/thrift/@v/v0.9.3.info): 410 Gone

1. `v0.9.3`这个tag找不到了，直接用commit找发现能找到。

## go get: inconsistent versions

1. 在`~/go/src/github.com/olivere/elastic`下get导致的，到非项目目录下get.

# 参考文章

1. <https://blog.golang.org/using-go-modules>