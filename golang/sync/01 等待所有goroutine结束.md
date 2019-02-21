# 目录
- [分析](#分析)
- [被动等待其他goroutine结束：sync.WaitGroup](#被动等待其他goroutine结束syncWaitGroup)
	- [sync.WaitGroup 介绍](#syncWaitGroup%20介绍)
	- [例子：被动等待其他goroutine结束](#例子被动等待其他goroutine结束)
- [主动通知其他goroutine结束：sync.WaitGroup+finish channel](#主动通知其他goroutine结束syncWaitGroupfinish%20channel)
	- [非阻塞式channel](#非阻塞式channel)
	- [例子：主动通知其他goroutine结束](#例子主动通知其他goroutine结束)
- [参考资料](#参考资料)

# 分析
1. go里面，main所在的goroutine退出后，资源被系统回收，其他的goroutine也会跟着退出。此时如果没有一些同步操作，会导致其他goroutine里面的程序执行到一半，产生不可知的后果。
2. 如果只是被动的等待其他goroutine退出，可以使用go标准库里面的`sync.WaitGroup`，见[被动等待：sync.WaitGroup](#被动等待其他goroutine结束：sync.WaitGroup).
3. 如果需要主动通知其他goroutine结束，只需要上面被动等待结束的基础上，加一个主动通知其他goroutine结束的功能即可。可以将一个finish channel当做参数传递给其他goroutine，需要主动结束时，main往finish里面写goroutine总数个值；其他goroutine每次执行完一个任务可以结束时，非阻塞式的从finish channel读数据，如果finish channel里面读取到了数据，则结束本goroutine。见[主动通知其他goroutine结束：sync.WaitGroup+finish channel](#主动通知其他goroutine结束：sync.WaitGroup+finish channel)

# 被动等待其他goroutine结束：sync.WaitGroup
## sync.WaitGroup 介绍

`sync.WaitGroup`用来等待所有的goroutine结束，如[文档](https://golang.org/pkg/sync/#WaitGroup)所言：

> A WaitGroup waits for a collection of goroutines to finish

`sync.WaitGroup`有三个函数Add、Done、Wait
```go
type WaitGroup
    func (wg *WaitGroup) Add(delta int)
    func (wg *WaitGroup) Done()
    func (wg *WaitGroup) Wait()
```
Add用来并发安全的加一个delta值到goroutine计数器上，delta可为负值。Done用来给goroutine计数器并发安全的减1。goroutine为0时，所有调用Wait阻塞的goroutine会解除阻塞开始继续执行。如果goroutine为负数，Add会panic。

## 例子：被动等待其他goroutine结束
- [源代码](wait_group/passive/main.go)：
```go
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
```
- 结果：
```
goroutine 4 start at 2019-02-20 21:42:08.109619 +0800 CST m=+0.012000701
goroutine 1 start at 2019-02-20 21:42:08.109619 +0800 CST m=+0.012000701
goroutine 2 start at 2019-02-20 21:42:08.109619 +0800 CST m=+0.012000701
goroutine 3 start at 2019-02-20 21:42:08.109619 +0800 CST m=+0.012000701
goroutine 1   end at 2019-02-20 21:42:09.1296774 +0800 CST m=+1.032059101
goroutine 2   end at 2019-02-20 21:42:10.1297346 +0800 CST m=+2.032116301
goroutine 3   end at 2019-02-20 21:42:11.1297918 +0800 CST m=+3.032173501
goroutine 4   end at 2019-02-20 21:42:12.129849 +0800 CST m=+4.032230701
```

# 主动通知其他goroutine结束：sync.WaitGroup+finish channel
## 非阻塞式channel
关键是在select 里面加一个default分支，详见[非阻塞式channel]()

## 例子：主动通知其他goroutine结束
- [源代码](wait_group/active/main.go)
```go
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

```
- 结果：
```
goroutine 1 finished at 2019-02-21 11:46:21.0209375 +0800 CST m=+1.005057501
goroutine 2 finished at 2019-02-21 11:46:22.0209947 +0800 CST m=+2.005114701
goroutine 3 finished at 2019-02-21 11:46:23.0210519 +0800 CST m=+3.005171901
goroutine 4 finished at 2019-02-21 11:46:24.0211091 +0800 CST m=+4.005229101
goroutine 5 finished at 2019-02-21 11:46:25.0211663 +0800 CST m=+5.005286301
all goroutine finished
```

# 参考资料
- [非阻塞式channel](https://gobyexample.com/non-blocking-channel-operations)
