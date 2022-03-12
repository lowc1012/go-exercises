package main

import "fmt"

type node struct {
	Key  string
	Val  string
	Next *node
}

type HashMap struct {
	Items []*node
	Size  int
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		make([]*node, size),
		size,
	}
}

func (h *HashMap) Get(key string) (string, bool) {
	index := h.getIndex(key)
	if h.Items[index] != nil {
		firstNode := h.Items[index]
		for ; ; firstNode = firstNode.Next {
			if firstNode.Key == key {
				return firstNode.Val, true
			}

			if firstNode.Next == nil {
				break
			}
		}
	}

	return "", false
}

func (h *HashMap) Put(key string, val string) {
	index := h.getIndex(key)
	if h.Items[index] == nil {

		h.Items[index] = &node{Key: key, Val: val}

	} else {

		// collision
		firstNode := h.Items[index]
		for ; firstNode.Next != nil; firstNode = firstNode.Next {
			if firstNode.Key == key {
				// replace
				firstNode.Val = val
				return
			}
		}
		firstNode.Next = &node{Key: key, Val: val}
	}
}

func (h *HashMap) Delete(key string) {
	index := h.getIndex(key)
	if h.Items[index] != nil {
		if h.Items[index].Key == key {
			h.Items[index] = h.Items[index].Next
		} else {
			head := h.Items[index]
			next := head.Next
			for next != nil {
				if next.Key == key {
					head.Next = next.Next
				}
				next = next.Next
			}
		}
	}
	return
}

func (h *HashMap) getIndex(key string) int {
	return int(hash(key)) % h.Size
}

// hash implements Jenkins hash function
func hash(key string) uint32 {
	var hash uint32
	hash = 0
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash
}

func main() {

	m := NewHashMap(3)
	m.Put("Ryan", "Software engineer")
	m.Put("Jessie", "UI/UX designer")
	m.Put("Danny", "Project Manager")
	m.Put("Roy", "Lawyer")
	m.Put("Kevin", "Office worker")
	m.Put("Jessica", "Teacher")

	m.Delete("Jessie")
	m.Delete("Ryan")

	if res, exist := m.Get("Jessie"); exist {
		fmt.Println("Find Jessie!")
		fmt.Printf("She is a %s!\n", res)
	} else {
		fmt.Println("Who is Jessie?")
	}

	if res, exist := m.Get("Ryan"); exist {
		fmt.Println("Find Ryan!")
		fmt.Printf("He is a %s!\n", res)
	} else {
		fmt.Println("Who is Ryan?")
	}

	if res, exist := m.Get("Roy"); exist {
		fmt.Println("Find Roy!")
		fmt.Printf("He is a %s!\n", res)
	} else {
		fmt.Println("Who is Roy?")
	}
}
