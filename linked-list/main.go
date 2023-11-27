/*
Demonstrates a singly linked list
*/
package main

import "fmt"

func main() {
	ll := list{}

	for i := 1; i <= 5; i++ {
		ll.insert(i)
	}

	ll.display()
}

// node represents an item in a linked list
type node struct {
	data int
	next *node
}

// list rpresents the starting point of the linked list
type list struct {
	head *node
}

// insert creates a new node with the data and adds it to the linked list
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

// display iterates through the linked list and displays the value in each node
func (l *list) display() {
	h := l.head

	for h != nil {
		fmt.Print(h.data, " ")
		h = h.next
	}

	fmt.Println("")
}
