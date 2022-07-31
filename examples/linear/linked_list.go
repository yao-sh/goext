package main

import (
	linear "github.com/yao-sh/goext/pkg/linear"
)

func main() {
	ll := linear.LinkedList{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	ll.Display()
	ll.Remove(3)
	ll.Display()
}
