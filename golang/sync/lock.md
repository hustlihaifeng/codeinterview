[TOC]

# golang锁分类

## 互斥锁

```go
var mutex sync.Mutex // 定义互斥锁变量 mutex
func write(){
    mutex.Lock( )
    defer mutex.Unlock( )
}
```

1. 加锁后其他对同一个锁加锁的地方就需要等待。

## 读写锁



### 读锁

```go
func (*RWMutex)RLock()
func (*RWMutex)RUnlock()
```

1. 读锁之间不互斥，读锁和写锁之间互斥。

### 写锁

```go
func (*RWMutex)Lock()
func (*RWMutex)Unlock()
```

2. 只要有写锁，就互斥。可以等同于互斥锁。

# 参考资料

1. [golang中关于读写锁、互斥锁的理解](<https://blog.csdn.net/wade3015/article/details/90692965>)