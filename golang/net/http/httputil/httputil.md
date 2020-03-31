# 重新编译go源码

## windows

1. 不用重新编译go所有的包， `go build -a` 即可。-a表示所有import的包都从源码编译，而不是引用已经编译好的静态库。
2. 需要删掉 `$GOROOT/pkg/` 下面对应package的.a文件，进入`$GOROOT/src/`对应目录，执行`go install`