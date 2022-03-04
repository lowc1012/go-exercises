/**
  HashSet (string)
  - ref: https://go.dev/play/p/_FvECoFvhq
  - generic version (for v1.18+): https://github.com/zyedidia/generic/tree/master/hashset
*/

package main

import "fmt"

type SHashSet struct {
	item map[string]bool
}

func NewSHashSet() *SHashSet {
	return &SHashSet{
		make(map[string]bool),
	}
}

func (sh *SHashSet) Get(s string) bool {
	_, exist := sh.item[s]
	return exist
}

func (sh *SHashSet) Add(s string) bool {
	_, exist := sh.item[s]
	sh.item[s] = true
	return !exist
}

func (sh *SHashSet) Remove(s string) {
	delete(sh.item, s)
}

func main() {
	s := NewSHashSet()
	fmt.Println("Add test1")
	s.Add("test1")
	fmt.Println(s.Get("test1"))
	fmt.Println("Removing test1")
	s.Remove("test1")
	fmt.Println(s.Get("test1"))
}
