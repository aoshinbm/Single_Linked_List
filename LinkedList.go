package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
}

func (l *LinkedList) Display() {
	if l.Head == nil {
		fmt.Println("Linked List is empty.")
		return
	}
	current := l.Head
	fmt.Println("Linked List values:")
	for current != nil {
		fmt.Print("-->", current.Value)
		current = current.Next
	}
}

func main() {
	list := &LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Display()
}
