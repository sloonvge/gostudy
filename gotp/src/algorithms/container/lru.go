package container

import "container/list"

type LRUCache struct {
	cap int
	list *list.List
	m map[int]*list.Element
}

type item struct {
	k int
	v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{ cap:capacity, list: list.New(), m:make(map[int]*list.Element) }
}


func (this *LRUCache) Get(key int) int {
	i, ok := this.m[key]
	if !ok {
		return -1
	}
	this.list.MoveToFront(i)
	return i.Value.(*item).v
}


func (this *LRUCache) Put(key int, value int)  {
	i, ok := this.m[key]
	if ok {
		this.list.MoveToFront(i)
		i.Value = &item{k:key, v:value}
		return
	}
	if this.list.Len() == this.cap {
		old := this.list.Back()
		oldKey := old.Value.(*item).k
		delete(this.m, oldKey)
		this.list.Remove(old)
	}
	this.list.PushFront(&item{k: key, v: value})
	this.m[key] = this.list.Front()
}

