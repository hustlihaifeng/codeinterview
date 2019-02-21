package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	num := 4
	for i := 1; i <= num; i++ {
		go sleepFunc(i, &wg)
	}
	wg.Wait()
}

func sleepFunc(sec int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	fmt.Printf("goroutine %v start at %v\n", sec, time.Now())
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf("goroutine %v   end at %v\n", sec, time.Now())
}
