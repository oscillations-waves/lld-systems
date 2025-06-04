package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data   int
	KeyPtr *list.Element
}

type LRUCache struct {
	Queue    *list.List
	Items    map[int]*Node
	Capacity int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{Queue: list.New(), Items: make(map[int]*Node), Capacity: capacity}

}

func (l *LRUCache) Get(key int) int {
	if item, found := l.Items[key]; found {
		l.Queue.MoveToFront(item.KeyPtr)
		return item.Data
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if item, ok := l.Items[key]; !ok {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			if back != nil {
				l.Queue.Remove(back)
				delete(l.Items, back.Value.(int))
			}
		}
		l.Items[key] = &Node{Data: value, KeyPtr: l.Queue.PushFront(key)}
	} else {
		item.Data = value
		l.Queue.MoveToFront(item.KeyPtr)
	}

}

func main() {
	cache := NewLRUCache(2)
	cache.Put(1, 10)
	cache.Put(2, 20)
	fmt.Println("Get 1:", cache.Get(1)) // Should print 10
	cache.Put(3, 30)                    // Evicts key 2
	fmt.Println("Get 2:", cache.Get(2)) // Should print -1 (not found)
	cache.Put(4, 40)                    // Evicts key 1
	fmt.Println("Get 1:", cache.Get(1)) // Should print -1 (not found)
	fmt.Println("Get 3:", cache.Get(3)) // Should print 30
	fmt.Println("Get 4:", cache.Get(4)) // Should print 40
}
