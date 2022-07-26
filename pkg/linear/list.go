package linear

import "fmt"

type LinkedListNode struct {
	Value interface{}
	Next  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
}

func (l *LinkedList) Add(value interface{}) {
	node := &LinkedListNode{Value: value}
	if l.Head == nil {
		l.Head = node
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (l *LinkedList) Remove(value interface{}) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) Length() int {
	current := l.Head
	count := 0
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (l *LinkedList) Display() {
	current := l.Head
	for current != nil {
		if current.Next != nil {
			fmt.Printf("%v => ", current.Value)
		} else {
			fmt.Printf("%v", current.Value)
		}
		current = current.Next
	}
	println()
}
