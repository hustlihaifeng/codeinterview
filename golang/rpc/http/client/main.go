package main

import (
	"fmt"
	"golang/rpc/arith"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &arith.Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("Arith: ", args.A, " * ", args.B, " = ", reply)
	time.Sleep(10)
}
