package main

import (
	"container/list"
	"fmt"
)

type LFUNode struct {
	key   int
	value int
	freq  int
	elem  *list.Element
}

type LFUCache struct {
	capacity   int
	size       int
	minFreq    int
	keyToNode  map[int]*LFUNode
	freqToList map[int]*list.List
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity:   capacity,
		keyToNode:  make(map[int]*LFUNode),
		freqToList: make(map[int]*list.List),
	}
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.keyToNode[key]
	if !ok {
		return -1
	}
	c.increaseFreq(node)
	return node.value
}

func (c *LFUCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}
	if node, ok := c.keyToNode[key]; ok {
		node.value = value
		c.increaseFreq(node)
		return
	}
	if c.size == c.capacity {
		lfuList := c.freqToList[c.minFreq]
		lfuElem := lfuList.Back()
		lfuNode := lfuElem.Value.(*LFUNode)
		lfuList.Remove(lfuElem)
		delete(c.keyToNode, lfuNode.key)
		c.size--
	}
	newNode := &LFUNode{key: key, value: value, freq: 1}
	if c.freqToList[1] == nil {
		c.freqToList[1] = list.New()
	}
	newElem := c.freqToList[1].PushFront(newNode)
	newNode.elem = newElem
	c.keyToNode[key] = newNode
	c.minFreq = 1
	c.size++
}

func (c *LFUCache) increaseFreq(node *LFUNode) {
	freq := node.freq
	lst := c.freqToList[freq]
	lst.Remove(node.elem)
	if lst.Len() == 0 {
		delete(c.freqToList, freq)
		if c.minFreq == freq {
			c.minFreq++
		}
	}
	node.freq++
	if c.freqToList[node.freq] == nil {
		c.freqToList[node.freq] = list.New()
	}
	node.elem = c.freqToList[node.freq].PushFront(node)
}

func main() {
	cache := NewLFUCache(2)
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
