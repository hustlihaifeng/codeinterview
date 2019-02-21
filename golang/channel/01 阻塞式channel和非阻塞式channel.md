# 目录
- [关键点](#关键点)
- [例子](#例子)
	- [只在一个goroutine中读写没有缓冲区的channel会产生死锁](#只在一个goroutine中读写没有缓冲区的channel会产生死锁)
	- [非阻塞的使用channel](#非阻塞的使用channel)
	- [给channel读写设置超时时间](#给channel读写设置超时时间)
- [测试](#测试)
- [参考资料](#参考资料)

# 关键点
 1. go的channel是阻塞式的。
 2. **channel默认缓冲区大小为0(而非1)**，也即往一个没有设置缓冲区大小的channel里面写数据，只有在其他goroutine在对这个channel读的时候，才能写入成功。极端情况是，如果一个没有设置缓冲区大小的channel只在一个goroutine里面使用，对该channel读写数据会产生死锁。
 3. 要想非阻塞的使用channel，需要在select中使用channel，并加一个default分支。如果channel中当前不能读或者写，就走default分支，以此来实现非阻塞式channel。
> select case must be receive, send or assign recv

select的case语句中可以执行读、写、读后赋值这三种操作。
4. 可以给channel读写操作设置一个超时时间，实现同样借助select，关键点是加一个`case <-time.After(time.Second):`这样的分支。
# 例子
## 只在一个goroutine中读写没有缓冲区的channel会产生死锁
```go
func deadlockChan() {
	var c1 chan string = make(chan string)
	c1 <- "haha"
	msg := <-c1
	fmt.Println(msg)
}
```
执行结果：
```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.deadlockChan()
```
上面这个死锁可以通过给channel设置一个非0的缓冲区或者将读写放在不同的goroutine解决：
- 通过非0缓冲区解决死锁
```go
func deadlockChanDealtByBuffer() {
	var c1 chan string = make(chan string, 1)
	c1 <- "haha"
	msg := <-c1
	fmt.Println(msg)
}
```
- 通过不同的goroutine读写channel解决死锁
```go
func deadlockChanDealtByGoroutine() {
	var c1 chan string = make(chan string, 1)
	go func() {
		time.Sleep(time.Second)
		c1 <- "haha"
	}()
	msg := <-c1
	fmt.Println(msg)
}
```
## 非阻塞的使用channel
- 关键点：select里面操作channel + default分支
```go
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
```
## 给channel读写设置超时时间
- 关键点：select里面操作channel + `case <-time.After(time.Second):`分支
```go
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
```
# 测试
代码见：[block_nonblock/main.go](block_nonblock/main.go)
结果：
```
haha

haha

channel if full or no goroutine is reading on the channel
channel is empty or no goroutine has written to the channel

channel is blocking, 1s passed, abort sending
channel is blocking, 1s passed, abort receiving
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.deadlockChan()
        D:/code/codeinterview/golang/channel/block_nonblock/main.go:26 +0x60
main.main()
        D:/code/codeinterview/golang/channel/block_nonblock/main.go:21 +0x95
```
# 参考资料
- [golang channel阻塞与非阻塞用法](https://zhuanlan.zhihu.com/p/22620172)
- [https://gobyexample.com/non-blocking-channel-operations](https://gobyexample.com/non-blocking-channel-operations)