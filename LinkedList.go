package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
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
func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
}

func (l *LinkedList) DeleteElement(delete int) {
	if l.Head == nil {
		fmt.Println("LinkedList is empty")
		return
	}
	if l.Head.Value == delete {
		//deleting the head node
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == delete {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) AddLast(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func main() {
	list := &LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)
	list.Display()

	fmt.Println(" ")
	list.DeleteElement(50)
	list.Display()

	fmt.Println(" ")
	list.AddLast(90)
	list.AddLast(100)
	list.AddLast(110)
	list.Display()
}
