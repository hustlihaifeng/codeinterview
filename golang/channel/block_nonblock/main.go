package main

import (
	"fmt"
	"time"
)

func main() {
	deadlockChanDealtByBuffer()
	fmt.Println()

	deadlockChanDealtByGoroutine()
	fmt.Println()

	deadlockChanDealtByNonblock()
	fmt.Println()

	deadlockChanDealtByTimeoutBlock()
	fmt.Println()

	deadlockChan()
}

func deadlockChan() {
	var c1 chan string = make(chan string)
	c1 <- "haha"
	msg := <-c1
	fmt.Println(msg)
}

func deadlockChanDealtByBuffer() {
	var c1 chan string = make(chan string, 1)
	c1 <- "haha"
	msg := <-c1
	fmt.Println(msg)
}

func deadlockChanDealtByGoroutine() {
	var c1 chan string = make(chan string, 1)
	go func() {
		time.Sleep(time.Second)
		c1 <- "haha"
	}()
	msg := <-c1
	fmt.Println(msg)
}

func deadlockChanDealtByNonblock() {
	var c1 chan string = make(chan string)
	select {
	case c1 <- "result 1":
		fmt.Println("channel is not full or some goroutine is reading on the channel")
	default:
		fmt.Println("channel if full or no goroutine is reading on the channel")
	}

	select {
	case msg := <-c1:
		fmt.Printf("received %v from channel\n", msg)
	default:
		fmt.Printf("channel is empty or no goroutine has written to the channel\n")
	}
}

func deadlockChanDealtByTimeoutBlock() {
	var c1 chan string = make(chan string)
	select {
	case c1 <- "result 1":
		fmt.Println("channel is not full or some goroutine is reading on the channel")
	case <-time.After(time.Second):
		fmt.Println("channel is blocking, 1s passed, abort sending")
	}

	select {
	case msg := <-c1:
		fmt.Printf("received %v from channel\n", msg)
	case <-time.After(time.Second):
		fmt.Printf("channel is blocking, 1s passed, abort receiving")
	}
}
