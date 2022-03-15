package mylist

import (
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	l := NewSinglyLinkedList()
	l.head = &Node{"test", nil}
	if l.head.Value != "test" {
		t.Error("error!")
	}
}

func TestSinglyLinkedList_PushFront(t *testing.T) {
	l := NewSinglyLinkedList()
	l.PushFront("first")
	if l.head.Value != "first" {
		t.Error("error!")
	}
	l.PushFront("second")
	if l.head.Value != "second" {
		t.Error("error!")
	}
}

func TestSinglyLinkedList_PushBack(t *testing.T) {
	l := NewSinglyLinkedList()
	l.PushBack("last")
	if l.head.Value != "last" {
		t.Error("error!")
	}
	l.PushBack("fist")
	if l.head.Next.Value != "fist" {
		t.Error("error!")
	}
}

func TestSinglyLinkedList_Clear(t *testing.T) {
	l := NewSinglyLinkedList()
	l.PushFront("hello")
	l.PushFront("first")
	l.PushBack("hi")
	l.PushBack("last")
	l.Clear()
	if l.head != nil {
		t.Error("error!")
	}
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	l := NewSinglyLinkedList()
	l.PushBack("first")
	l.PushBack("second")
	l.PushBack("third")
	if l.head.Value != "first" {
		t.Error("error!")
	}
	l.Reverse()
	if l.head.Value != "third" {
		t.Error("error!")
	}
}
