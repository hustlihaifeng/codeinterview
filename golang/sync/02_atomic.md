# 1. 哪个库提供`CAS`原子操作

1. 标准库`sync/atomic`

# 2. 原子操作可以对哪些数据类型使用

1. 类型包括`int32,int64,uint32,uint64,uintptr,unsafe.Pointer`，共六个
2. int长度与平台相关
3. `uintptr`在32位平台下为4字节,64位平台下为8字节
4. 如：

```go
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
type Value
    func (v *Value) Load() (x interface{})
    func (v *Value) Store(x interface{})
eg:
	var config atomic.Value
	config.Store(xxx)
	config.Store(xxx)
```

 - 为了安全的考虑，Go语言是不允许两个指针类型进行转换的。
 - `unsafe.Pointer`是一种特殊意义的指针，它可以包含任意类型的地址，有点类似于C语言里的void*指针，全能型的。
    - 任何指针都可以转换为`unsafe.Pointer`
   - `unsafe.Pointer`可以转换为任何指针
   - `uintptr`可以转换为`unsafe.Pointer`
   - `unsafe.Pointer`可以转换为`uintptr`
   - 详见： <https://www.flysnow.org/2017/07/06/go-in-action-unsafe-pointer.html>

```go
	i:= 10
	ip:=&i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
```

# 3. 有哪些原子操作

## 3.1 增或减Add

1. 被操作的类型只能是数值类型
2. `int32,int64,uint32,uint64,uintptr`类型可以使用原子增或减操作
3. 第一个参数值必须是一个指针类型的值，以便施加特殊的CPU指令
4. 第二个参数值的类型和第一个被操作值的类型总是相同的。
5. 如：

```go
    var i64 int64
    newI64 := atomic.AddInt64(&i64,-3)
```

## 3.2 比较并交换`CAS`

1. 函数定义：

   ```go
   func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
   ```

2. 调用函数后，会先判断参数addr指向的被操作值与参数old的值是否相等.仅当此判断得到肯定的结果之后，才会用参数new代表的新值替换掉原先的旧值，否则操作就会被忽略。so, 需要用for循环不断进行尝试,直到成功为止

3. 使用CAS操作的做法趋于乐观(**乐观锁**): 总是假设被操作值未曾被改变（即与旧值相等），并一旦确认这个假设的真实性就立即进行值替换。

4. 如：

```go
//不断地尝试原子地更新value的值,直到操作成功为止
func addValue(delta int32){
    //在被操作值被频繁变更的情况下,CAS操作并不那么容易成功
    //so 不得不利用for循环以进行多次尝试
    for {
        v := value
        if atomic.CompareAndSwapInt32(&value, v, (v + delta)){
            //在函数的结果值为true时,退出循环
            break
        }
        //操作失败的缘由总会是value的旧值已不与v的值相等了.
        //CAS操作虽然不会让某个Goroutine阻塞在某条语句上,但是仍可能会使流产的执行暂时停一下,不过时间大都极其短暂.
    }
}
```

## 3.3 载入Load

1. 上面的比较并交换案例总 v:= value为变量v赋值，但… 要注意，在进行读取value的操作的过程中,其他对此值的读写操作是可以被同时进行的,**那么这个读操作很可能会读取到一个只被修改了一半的数据**. so , 我们要使用sync/atomic代码包同样为我们提供了一系列的函数，**以Load为前缀(载入)，来确保这样的糟糕事情发生**。
2. 如：

```go
//不断地尝试原子地更新value的值,直到操作成功为止
func addValue(delta int32){
    //在被操作值被频繁变更的情况下,CAS操作并不那么容易成功
    //so 不得不利用for循环以进行多次尝试
    for {
        //v := value
        //在进行读取value的操作的过程中,其他对此值的读写操作是可以被同时进行的,那么这个读操作很可能会读取到一个只被修改了一半的数据.
        //因此我们要使用载入
        v := atomic.LoadInt32(&value)
        if atomic.CompareAndSwapInt32(&value, v, (v + delta)){
            //在函数的结果值为true时,退出循环
            break
        }
        //操作失败的缘由总会是value的旧值已不与v的值相等了.
        //CAS操作虽然不会让某个Goroutine阻塞在某条语句上,但是仍可能会使流产的执行暂时停一下,不过时间大都极其短暂.
    }
}
```

- atomic.LoadInt32接受一个*int32类型的指针值
- 返回该指针指向的那个值

## 3.4 存储Store

1. 与读取操作相对应的是写入操作。 而sync/atomic包也提供了与原子的载入函数相对应的原子的值存储函数。 以Store为前缀
2. **在原子地存储某个值的过程中，任何CPU都不会进行针对同一个值的读或写操作**。
3. **原子的值存储操作总会成功，因为它并不会关心被操作值的旧值是什么**

### 3.4.1 交换Swap

1. 直接设置新值,返回被操作值的旧值. 此类操作比`CAS`操作的约束更少，同时又比原子载入操作的功能更强

# 4. 无锁map实现

## 4.1 map的实现

1. 参考 <https://tiancaiamao.gitbooks.io/go-internals/content/zh/02.3.html>
2. Go中的map在底层是用哈希表实现的，你可以在 `$GOROOT/src/pkg/runtime/hashmap.goc` 找到它的实现。

### 4.1.1 map数据结构

1. map数据结构：

```go
struct Hmap
{
    uint8   B;    // 可以容纳2^B个项
    uint16  bucketsize;   // 每个桶的大小

    byte    *buckets;     // 2^B个Buckets的数组
    byte    *oldbuckets;  // 前一个buckets，只有当正在扩容时才不为空
};
```

- hash值mod当前hash表大小决定某一个值属于哪个桶，而hash表大小是2的指数，即上面结构体中的2^B
- 每次扩容，会增大到上次大小的两倍。结构体中有一个buckets和一个`oldbuckets`是用来实现增量扩容的。正常情况下直接使用buckets，而`oldbuckets`为空。如果当前哈希表正在扩容中，则`oldbuckets`不为空，并且buckets大小是`oldbuckets`大小的两倍。

### 4.1.2 map如何变化

## 4.2 无锁map的实现

1. 思路：使用`cas`，操作前读取，然后使用`CompareAndSwap`来对一个指针进行替换。操作对象是一个指针，如果对整个map用一个指针来操作的话，那么每次操作都需要对整个map进行替换，消耗太大。如果map里面的每个value都是一个原子操作的指针，似乎可以实现，也就是需要传给map的都是key和指针，这样才能用原子操作。那么问题的关键点变为如何将用户传入的value转换为一个指针？问题的复杂性在于，map并不能获取到存储value的这个变量的指针，map返回的直接是value，也就不能对该地址进行`CAS`操作。

2. `sync.Map`的解决办法是，底层的value，直接是指针：`map[interface{}]*entry`，那么我们可以从map中得到指针后，在对该指针进行`CAS`操作。问题在于从map中得到指针的这个过程中发生并发问题怎么办？`sync.Map`的解决办法是：读写分离，底层有两个map；读map read的key、value只有读（也即key和右边的指针变量的地址都不变）；写map dirty的key和value有读写，读写时都加锁；先在读map中不加锁的找，没找到再在写map中加锁找；找到通过atomic操作改变指针所指向的值（注意此时read map中的key和value都是不变的）；当读map中找不到达到一定次数时，加锁将写map升级为读map，进行同步；这样大部分的读都是无锁的，写是加锁的。几个问题：

   1. 读map中的查和update好说，删除怎么办？将指针的值设置为nil？那么如果value值本身是nil怎么办？将指针值指向一个变量，该变量的值是nil。那么说明指针的值本身也是一个指针(`unsafe.Pointer`)，也即指针的指针.

   2. 读map和写map什么时候会不一致？

       - 在第一个问题的基础上，load需要取得指针所指向的值，在对该值进行间访操作。

       ```go
       func (e *entry) load() (value interface{}, ok bool) {
        p := atomic.LoadPointer(&e.p)
        if p == nil || p == expunged {
            return nil, false
        }
        return *(*interface{})(p), true
       }
       ```

       - store需要把指针的值设置为输入变量的地址

       ```go
       func (e *entry) tryStore(i *interface{}) bool {
        for {
            p := atomic.LoadPointer(&e.p)
            if p == expunged {
                return false
            }
            if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
                return true
            }
        }
       }
       // Store sets the value for a key.
       func (m *Map) Store(key, value interface{}) {
        read, _ := m.read.Load().(readOnly)
        if e, ok := read.m[key]; ok && e.tryStore(&value) {
            return
        }
       ```

       - 删除如1所述
       - 那么此时update、delete、select都可以在只读map中通过atomic操作来完成，如果只读map中有对应的key的话。add需要在读写map中完成。也即，读写map中只有比read中多的key。

   3. 读写map如何进行同步？

       1. 简单的，全局加锁，给读map加key？那么此时在读map中读的时候也需要加锁，才能保证读map访问没有并发问题，与前面的相矛盾。

       2. `sync.Map`使用的是直接把指向读map的变量指向写map:`m.read.Store(readOnly{m: m.dirty})`。那么这就要求：

           - 就对该变量操作时使用atomic操作。

             ```go
             type Map struct {
             	read atomic.Value // readOnly
             }
             func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
             	read, _ := m.read.Load().(readOnly)
             	e, ok := read.m[key]
             ```

             

           - 写map中要有所有读map中的k v

             - 写更新为读之后，写map保留
             - 由于原来的update，delete只更新读map，那么要求读写map的value指针是同一个变量。也就是读写map 有两份key，一份值。

       3. 上面这样做，功能上是没问题的，但是由于删除时没有清理key，会造成空间上的浪费，且使用时一直累加。`sync.Map`中考虑到了这个问题，新增加了一个`expunged`状态来处理这种情况：

       ```go
       // expunged is an arbitrary pointer that marks entries which have been deleted
       // from the dirty map.
       var expunged = unsafe.Pointer(new(interface{}))
       
       // An entry is a slot in the map corresponding to a particular key.
       type entry struct {
       	// p points to the interface{} value stored for the entry.
       	//
       	// If p == nil, the entry has been deleted and m.dirty == nil.
       	//
       	// If p == expunged, the entry has been deleted, m.dirty != nil, and the entry
       	// is missing from m.dirty.
       	//
       	// Otherwise, the entry is valid and recorded in m.read.m[key] and, if m.dirty
       	// != nil, in m.dirty[key].
       	//
       	// An entry can be deleted by atomic replacement with nil: when m.dirty is
       	// next created, it will atomically replace nil with expunged and leave
       	// m.dirty[key] unset.
       	//
       	// An entry's associated value can be updated by atomic replacement, provided
       	// p != expunged. If p == expunged, an entry's associated value can be updated
       	// only after first setting m.dirty[key] = e so that lookups using the dirty
       	// map find the entry.
       	p unsafe.Pointer // *interface{}
       }
       ```

       - p指向真正的value变量。
       - p等于nil和`expunged`都表示该`kv`对已经被删除
         - p为nil时，`m.dirty`也为nil
         - p为`expunged`时，`m.dirty`不为nil，但是由于后面写map转换为读map时，是直接替换读map，所以写map `m.dirty`中不能有被删除的key. 通过在m.dirty被创建时将读map中的p改为`expunged`来实现（见源码`tryExpungeLocked`函数）(并且要将读map中不是删除状态的key拷贝到`m.dirty`中)（见源码`dirtyLocked函数`）。
         - 对已经删除了的key重新设置值时，如果`m.dirty`为空，那么近期不会发生读写map替换，直接更新读map的值（后面`m.dirty`创建的时候会把这个值给赋值过来）。`m.dirty`不为空时，需要直接更改`m.dirty`里面的，如果此时更改读map，那么`m.dirty`中将缺失该key。

   4. 注意在`range`时，如果写map中有多的，需要先将写map转化为读map。注意源码中的，在加锁后的二次检查写map是否多；如果没有的话，那么加锁过程中写map并发转化为读map（写map一并清零），如果此时直接用写map替换读map，那么会导致所有数据丢失。**加锁后二次检验，防止加锁过程中前述判断条件被并发修改了**

   5. <https://colobu.com/2017/07/11/dive-into-sync-Map/>这边文章有对`sync.Map`两点的总结。

