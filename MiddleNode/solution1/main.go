package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (ll *LinkedList) walk(count int) int {
	if ll.Next != nil {
		count++
		count = ll.Next.walk(count)
	}
	return count
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
	count := linkedList.walk(1)
	for i := 1; i <= count/2; i++ {
		linkedList = linkedList.Next
	}
	return linkedList
}

func main() {
	ll := &LinkedList{Value: 2}
	ll.Next = &LinkedList{Value: 7}
	ll.Next.Next = &LinkedList{Value: 3}
	ll.Next.Next.Next = &LinkedList{Value: 5}

	node := MiddleNode(ll)
	fmt.Println(node)
}
