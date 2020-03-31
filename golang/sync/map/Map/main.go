package main

import (
	"fmt"
	"sync"
)

func main() {
	var cm sync.Map
	cm.Store(1, 1)
	rst, ok := cm.Load(1)
	fmt.Println(cm.Load(1))
	fmt.Printf("%#v %#v\n", rst, ok)
	fmt.Printf("%T %T\n", rst, ok)
}
