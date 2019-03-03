package main

import (
	"fmt"
	"golang/rpc/arith"
	"net/rpc"
)

func main() {
	var client, err = rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("连接不到服务器：", err)
	}
	var args = arith.Args{A: 40, B: 3}
	var result int
	fmt.Println("开始调用！")
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		fmt.Println("调用失败！", err)
	}
	fmt.Println("Arith: ", args.A, " * ", args.B, " = ", result)
}
