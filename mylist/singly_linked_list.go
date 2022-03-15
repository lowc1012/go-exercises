/**
 * Definition for singly-linked list.
 * type Node struct {
 *     Value string
 *     Next *Node
 * }
 */

package mylist

import "fmt"

type SinglyLinkedList struct {
	head *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (l *SinglyLinkedList) PushFront(val string) {
	if l.head == nil {
		l.head = &Node{val, nil}
	} else {
		first := l.head
		l.head = &Node{val, first}
		first.Next = nil
	}
}

func (l *SinglyLinkedList) PushBack(val string) {
	n := new(Node)
	n.Value = val

	if l.head == nil {
		l.head = n
	} else {
		curr := l.head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = n
	}
}

func (l *SinglyLinkedList) Delete(val string) {
	if l.head != nil {
		curr := l.head
		if curr.Value == val {
			curr = curr.Next
		}
		for curr.Next != nil {
			if curr.Next.Value == val {
				curr.Next = curr.Next.Next
			}
		}
	}
}

func (l *SinglyLinkedList) Clear() {
	if l.head != nil {
		curr := l.head
		nextPtr := curr.Next
		for ; nextPtr != nil; curr = curr.Next {
			curr.Next = curr.Next.Next
			nextPtr = nil
		}
		l.head = nil
	}
}

func (l *SinglyLinkedList) Reverse() {
	curr := l.head
	var prev *Node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *SinglyLinkedList) PrintList() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.Value + " ")
		curr = curr.Next
	}
}
