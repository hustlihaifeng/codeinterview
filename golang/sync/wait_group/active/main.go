package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	num := 5
	finish := make(chan bool, num)

	for i := 1; i <= num; i++ {
		go activeFinish(i, finish, &wg)
	}

	time.Sleep(time.Millisecond * time.Duration(500))

	for i := 1; i <= num; i++ {
		finish <- true
	}

	wg.Wait()
	fmt.Println("all goroutine finished")
}

func activeFinish(id int, finish <-chan bool, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for true {
		select {
		case <-finish:
			fmt.Printf("goroutine %v finished at %v\n", id, time.Now())
			return
		default:
		}

		time.Sleep(time.Second * time.Duration(id))
	}
}
