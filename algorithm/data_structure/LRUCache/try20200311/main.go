package main

import "container/list"

func main() {

}

type AllOne struct {
	Ref    map[string]*list.Element
	LstRef map[int]*list.List
	max    int
	min    int
}

type kcnt struct {
	k   string
	cnt int
}

/*
inc:
初始化LstRef[0]为空lst
if key 在Ref中存在 {
    从Ref中获取Elem，获取长度cnt
    获取长度为cnt的链，将Elem从LstRef[cnt]中删除

} else { // key不存在
    得到Elem{key,0}
    将key，Elem加入Ref
    cnt=0
}
if min==cnt && len(LstRef[cnt])==0 {
    min=cnt+1
}
if cnt+1 > max {
    max = cnt+1
}
cnt=cnt+1,Elem.cnt=cnt
从LstRef中获取长度为cnt的链，没有则创建并加入LstRef
将key，Elem加入长度为cnt的链


dec:
if key在Ref中存在 {
    获取Elem并获取长度cnt
    从LstRef[cnt]中删除Elem
    if max==cnt && len(LstRef[cnt])==0 {
        max=cnt-1
    }
    if cnt==1 && len(LstRef[1])>0 {
    }else{
        if cnt-1<min{
            min=cnt-1
        }
    }
    if cnt != 1 {
        cnt=cnt-1,Elem.cnt=cnt
        获取长度为cnt的链，没有则创建
        将Elem加入LstRef[cnt]
    }else{
        从Ref中删除key
    }
}

getMin：
if min==0 {
    return ""
}
return LstRef[min].Front().Value.(kcnt).k
getMax:
if max==0 {
    return ""
}
return LstRef[max].Front().Value.(kcnt).k

*/

/*
链表的增删改都是O(1)的操作，链表加hash又可以实现按key索引时，链表查为O(1)
*/
/** Initialize your data structure here. */
func Constructor() AllOne {
	rst := AllOne{
		Ref:    make(map[string]*list.Element),
		LstRef: make(map[int]*list.List),
		max:    0,
		min:    0,
	}
	// 初始化LstRef[0]为空lst
	rst.LstRef[0] = list.New()
	return rst
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	cnt := 0
	if elem, ok := this.Ref[key]; ok {
		// 	从Ref中获取Elem，获取长度cnt
		// 	获取长度为cnt的链，将Elem从LstRef[cnt]中删除
		cnt = elem.Value.(*kcnt).cnt
		this.LstRef[cnt].Remove(elem)
	} else {
		// 	cnt=0
		cnt = 0
	}
	if this.min == cnt && this.LstRef[cnt].Len() == 0 {
		this.min = cnt + 1
	}
	if cnt+1 > this.max {
		this.max = cnt + 1
	}
	cnt = cnt + 1
	// 从LstRef中获取长度为cnt的链，没有则创建并加入LstRef
	_, ok := this.LstRef[cnt]
	if !ok {
		this.LstRef[cnt] = list.New()
	}
	// 将key，Elem加入长度为cnt的链
	elem := this.LstRef[cnt].PushBack(&kcnt{k: key, cnt: cnt})
	this.Ref[key] = elem
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if elem, ok := this.Ref[key]; ok {
		// 获取Elem并获取长度cnt
		cnt := elem.Value.(*kcnt).cnt
		// 从LstRef[cnt]中删除Elem
		this.LstRef[cnt].Remove(elem)
		if this.max == cnt && this.LstRef[cnt].Len() == 0 {
			this.max = cnt - 1
		}
		if cnt == 1 && this.LstRef[1].Len() > 0 {
		} else {
			if cnt-1 < this.min {
				this.min = cnt - 1
			}
		}
		if cnt != 1 {
			cnt = cnt - 1

			// 获取长度为cnt的链，没有则创建
			_, ok := this.LstRef[cnt]
			if !ok {
				this.LstRef[cnt] = list.New()
			}
			// 将Elem加入LstRef[cnt]
			elem = this.LstRef[cnt].PushBack(&kcnt{k: key, cnt: cnt})
			this.Ref[key] = elem
		} else {
			// 从Ref中删除key
			delete(this.Ref, key)
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {

}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {

}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
