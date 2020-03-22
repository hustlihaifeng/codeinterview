# 问题
1. 找无序有重复数组的第k大元素

# 思路
1. 先排序，nlogn,然后返回第k个，1:也是范例
```go
func findKthLargest(nums []int, k int) int {
    sort.Ints(nums)
    return nums[len(nums)-k]
}
```
2. 维护一个k大的数组，扫描一遍。kn，平均1/2n^2
3. 类似快排的分而治之，理论上肯定比全部排好序的快排快。
    1. 关键点是：重复元素怎么处理：
        - 将重复元素放到左边，那么在就是重复元素的时候，找到最后选的pivot都是同一个，但是左边的永远达不到要求的k个。
        - 左边小，中间等（只需要计数即可），右边大
    2. 存储空间用什么？快排中，将pivot移动到第一个元素。那么等于的怎么移动？因为找k大，所以把等于的放到大的里面，并对等于的进行计数。找的时候，如果在等于的这个区间，就是pivot值；否则将等于的放到大的里面一并查找。

# 类似快排分治发伪代码
```go
取最左边第一个元素作为pivot，与最后一个元素交换，pivot计数1
从0，次小索引开始找
for lidx<rdix {
    for ;lidx <ridx && nums[lidx]<pivot; lidx++{}
    for ;ridx > lidx && nums[ridx]>=pivit; ridx--{
        if nums[ridx]==pivot{
            // TODO：与后面不是pivot的交换，位置可以由lidx和pivot计数得到，来避免递归是一直取到pivot导致的不能结束问题
            pivot计数加一
        }
    }
    if lidx<ridx{
        swap(lidx,ridx) // 这里比改变idx，以免造成有多种结束条件
    }
}
if nums[lidx]<pivot{
    lidx++ // lidx左边都是小的，[lidx右边是>=
}else{
    // lidx左边都是小的，[lidx右边是>=
    if nums[lidx]==pivot{
        // 加上上面的额TODO
        pivot计数加1
    }
}
sNum := lidx-bgnIdx
beNum := endIdx - bgnIdx+1 - sNum=endIdx-lidx+1
if beNum == k {
    return pivot
}else if beNum < k {
    return Recurse(nums,bgnIdx,lidx-1,k-beNum)
}else{ // beNum > k
    bNum = beNum - pivot数
    if nNum >= k {
        // TODO:与上面的额TODO配合
        return Recurse(nums,lidx,endIdx,k) // 如果下次一直取到pivot。那么会一直递归下去
    }else{
        return pivot
    }
}
```

# 详见
1. [main.go](main.go)