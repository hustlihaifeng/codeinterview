package main

import (
	"golang/rpc/arith"
	"log"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	var err error
	arith := new(arith.Arith)

	err = rpc.Register(arith)
	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()

	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("listen error:", err)
	} else {
		time.Sleep(3600 * time.Second)
	}
}
