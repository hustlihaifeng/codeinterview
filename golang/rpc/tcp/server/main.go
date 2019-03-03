package main

import (
	"fmt"
	"golang/rpc/arith"
	"net"
	"net/rpc"
	"time"
)

func main() {
	var ms = new(arith.Arith)
	rpc.Register(ms)
	var address, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		fmt.Println("启动失败！", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("接收到一个调用请求...")
		rpc.ServeConn(conn)
	}
	time.Sleep(3600 * time.Second)
}
