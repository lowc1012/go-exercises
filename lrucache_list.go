// Use a doubly-linked list (container/list) to implement LRU cache
package main

import (
	"container/list"
	"fmt"
)

type Pair struct {
	Key int
	Val int
}

type LRUCache struct {
	l        *list.List
	m        map[int]*list.Element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list.New(),
		make(map[int]*list.Element, capacity),
		capacity,
	}
}

func (c *LRUCache) Get(key int) int {
	if node, exist := c.m[key]; exist {
		// move the node to front
		c.l.MoveToFront(node)
		return node.Value.(*list.Element).Value.(Pair).Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// if cache hit
	if node, exist := c.m[key]; exist {
		// update the value of c node
		node.Value.(*list.Element).Value = Pair{
			key,
			value,
		}
		// move the node to front
		c.l.MoveToFront(node)

		// if cache miss
	} else {

		// push the new node into list
		ptr := c.l.PushFront(&list.Element{
			Value: Pair{
				Key: key,
				Val: value,
			},
		})

		// add the new item into map
		c.m[key] = ptr

		// if cache is full
		if c.l.Len() > c.capacity {
			// find the node we want to delete
			lastKey := c.l.Back().Value.(*list.Element).Value.(Pair).Key
			// delete the item in map
			delete(c.m, lastKey)

			// delete the node
			c.l.Remove(c.l.Back())
		}
	}
}

func main() {
	obj := Constructor(2)   // nil
	obj.Put(1, 10)          // nil, linked list: [1:10]
	obj.Put(2, 20)          // nil, linked list: [2:20, 1:10]
	fmt.Println(obj.Get(1)) // 10, linked list: [1:10, 2:20]
	obj.Put(3, 30)          // nil, linked list: [3:30, 1:10]
	fmt.Println(obj.Get(2)) // -1, linked list: [3:30, 1:10]
	obj.Put(4, 40)          // nil, linked list: [4:40, 3:30]
	fmt.Println(obj.Get(1)) // -1, linked list: [4:40, 3:30]
	fmt.Println(obj.Get(3)) // 30, linked list: [3:30, 4:40]
}
