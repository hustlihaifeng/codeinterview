# 目录
- [经验](#经验)
- [分析](#分析)
- [stack和FIFO queue的实现](#stack和FIFO-queue的实现)
	- [stack](#stack)
	- [FIFO queue](#FIFO-queue)
- [二叉树深度优先遍历](#二叉树深度优先遍历)
	- [二叉树中序遍历](#二叉树中序遍历)
	- [二叉树后序遍历](#二叉树后序遍历)
	- [二叉树先序遍历](#二叉树先序遍历)
- [二叉树广度优先遍历](#二叉树广度优先遍历)
# 经验
1. 将异常处理、初值处理、边界异常处理这些都通过函数预留，最开始只考虑大部分场景的逻辑（即该函数的核心逻辑）。（概要设计时）可以都用函数代替详细逻辑。
2. 面试时把自己的设备带着，这样用自己的设备，顺手一些。
3. 代码不超过25行（14寸屏幕中vscode的一个屏幕），否则一定是划分不清晰。一个良好的代码，如下面的这些，应该只有一个分支判断。

# 分析
1. 先序遍历、中序遍历、后序遍历都是深度优先搜索。

2. 深度优先搜索遇到时不一定访问，此时需要用栈来保持当前节点，先进后出。

3. 广度优先遍历遇到时就访问，此时需要用队列来实现，先进先出。

# stack和FIFO queue的实现

stack和队列都是线性数据结构，都可以用go的标准容器中的list来实现：
## stack
```go
	func StackPush(lst *list.List, elem int)error{
		if StackPushInputInvalid(lst, elem){
             return errors.New("input invalid")
		}
         
         lst.PushBack(elem)
         return nil
	}
	func StackPop(lst *list.List) int,error {
         if StackPopInputInvalid(lst) {
             return 0,errors.New("input invalid")
         }
         
         return lst.Remove(lst.Back()).(int),nil
	}
```
## FIFO queue
```go
	func QueuePush(lst *list.List, elem int)error{
		if QueuePushInputInvalid(lst, elem){
             return errors.New("input invalid")
		}
         
         lst.PushBack(elem)
         return nil
	}
	func QueuePop(lst *list.List) int,error {
         if QueuePopInputInvalid(lst) {
             return 0,errors.New("input invalid")
         }
         
         return lst.Remove(lst.Front()).(int),nil
	}
```
# 二叉树深度优先遍历
深度优先搜索包括先序、中序、后续遍历。先序实现最简单，中序符合我们一般的认知。
## 二叉树中序遍历
1. 伪代码
```go
for 栈非空 {
    pop得到当前节点
    if 左子树为空或者已经被访问 {
        访问当前节点
		将非空右子树压栈
	}else{
		当前节点压栈
		非空左子树压栈
	}
}
```
2. 代码见[transfer.go](transfer.go)里的`TransferLeftMiddleRight`，`TransferLeftMiddleRightBad`，测试代码见`test/main.go`
## 二叉树后序遍历
1. 伪代码
```go
for 栈非空 {
	pop 得到当前节点
	if 左子树为空或者被访问 && 右子树为空或者被访问 {
        访问当前节点
	}else{
        当前节点压栈
        非空右子树压栈
        非空左子树压栈
	}
}
```
2. 代码见[transfer.go](transfer.go)里的`TransferLeftRightMiddle`，`TransferLeftRightMiddleBad`，测试代码见`test/main.go`
## 二叉树先序遍历
1. 伪代码
```go
for 栈非空 {
    访问当前节点
    非空右子树压栈
    非空左子树压栈
}
```
2. 代码见[transfer.go](transfer.go)里的`TransferMiddleLeftRight`，`TransferMiddleLeftRightBad`，测试代码见`test/main.go`
# 二叉树广度优先遍历
1. 伪代码
```go
for 队列非空 {
    pop得到最老节点
    访问当前节点
    for 每一个子节点{
        if 节点非空且没有被访问过 {
            加入队列尾部
    	}
    }
}
```
2. 代码见[transfer.go](transfer.go)里的`TransferWidthFirst`，测试代码见`test/main.go`