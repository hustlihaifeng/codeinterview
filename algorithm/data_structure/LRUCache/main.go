package main

import "container/list"

func main() {

}

type LRUCache struct {
	index     map[int]*list.Element
	lru       *list.List
	maxLength int
	length    int
}
type KV struct {
	k int
	v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		index:     make(map[int]*list.Element),
		lru:       list.New(),
		maxLength: capacity,
		length:    0}
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.index[key]; !ok {
		return -1
	} else {
		this.lru.MoveToFront(e)
		return e.Value.(KV).v
	}
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.index[key]; !ok {
		if this.length < this.maxLength {
			this.lru.PushFront(KV{key, value})
			this.index[key] = this.lru.Front()
			this.length++
		} else {
			e = this.lru.Back()
			delete(this.index, e.Value.(KV).k)
			e.Value = KV{key, value}
			this.lru.MoveToFront(e)
			this.index[key] = e
		}
	} else {
		this.lru.MoveToFront(e)
		this.lru.Front().Value = KV{key, value}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
