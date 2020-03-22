package main

import "container/list"

func main() {

}

type AllOne struct {
	Ref    map[string]*list.Element // 一级索引，用来对元素进行定位，其value是一个kcnt结构
	LstRef map[int]*list.Element    // 二级索引，用来对不同长度的链表进行定位，其value是一个链表。
	Lst    *list.List               // 二级链表，和二级索引LstRef相配对
}

type kcnt struct {
	k   string
	cnt int
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	rst := AllOne{
		Ref:    make(map[string]*list.Element),
		LstRef: make(map[int]*list.Element),
		Lst:    list.New(),
	}
	elem := rst.Lst.PushBack(list.New())
	rst.LstRef[0] = elem
	return rst
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	var cnt int
	if elem, ok := this.Ref[key]; ok {
		// 获取计数cnt
		cnt = elem.Value.(*kcnt).cnt
		// 从old链中删除
		this.LstRef[cnt].Value.(*list.List).Remove(elem)
	} else {
		// 计数为0
		cnt = 0
	}
	// 如果cnt链没有，则新建
	if _, ok := this.LstRef[cnt+1]; !ok {
		elem := this.Lst.InsertAfter(list.New(), this.LstRef[cnt]) // 这里需要初始化一个cnt为0的一个空链表
		this.LstRef[cnt+1] = elem
	}
	// 应该先加后删除
	if this.LstRef[cnt].Value.(*list.List).Len() == 0 && cnt > 0 {
		this.Lst.Remove(this.LstRef[cnt])
		delete(this.LstRef, cnt)
	}
	// cnt链中添加元素，返回*list.Element
	elem := this.LstRef[cnt+1].Value.(*list.List).PushBack(&kcnt{k: key, cnt: cnt + 1})
	// 更新Ref[key]
	this.Ref[key] = elem
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	var cnt int
	if elem, ok := this.Ref[key]; !ok {
		return
	} else {
		cnt = elem.Value.(*kcnt).cnt
		// 从老链中删除
		this.LstRef[cnt].Value.(*list.List).Remove(elem)
		delete(this.Ref, key)
	}

	// 如果新链没有，则创建
	if _, ok := this.LstRef[cnt-1]; !ok {
		elem := this.Lst.InsertBefore(list.New(), this.LstRef[cnt])
		this.LstRef[cnt-1] = elem
	}
	if this.LstRef[cnt].Value.(*list.List).Len() == 0 {
		this.Lst.Remove(this.LstRef[cnt])
		delete(this.LstRef, cnt)
	}
	// 将key，cnt加入新链
	if cnt == 1 {
		return
	}
	elem := this.LstRef[cnt-1].Value.(*list.List).PushBack(&kcnt{k: key, cnt: cnt - 1})
	// 更新Ref[key]
	this.Ref[key] = elem
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMinKey() string {
	if this.Lst.Len() < 2 {
		return ""
	}
	lst := this.Lst.Front().Next().Value.(*list.List)
	return lst.Front().Value.(*kcnt).k
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMaxKey() string {
	if this.Lst.Len() < 2 {
		return ""
	}
	lst := this.Lst.Back().Value.(*list.List)
	return lst.Front().Value.(*kcnt).k
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
