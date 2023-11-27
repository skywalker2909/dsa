/*
Demonstrates a singly linked list
*/
package main

import "fmt"

func main() {
	ll := list{}

	ll.insert(1)
	ll.insert(2)
	ll.insert(3)
	ll.insert(4)
	ll.insert(5)

	ll.display()
}

type node struct {
	data int
	next *node
}

type list struct {
	head *node
}

func (l *list) insert(data int) {
	n := &node{data: data, next: nil}

	if l.head == nil {
		l.head = n
		return
	}

	h := l.head
	for h.next != nil {
		h = h.next
	}

	h.next = n
}

func (l *list) display() {
	h := l.head

	for h != nil {
		fmt.Print(h.data, " ")
		h = h.next
	}

	fmt.Println("")
}
