package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	base := linkedList
	if linkedList.Next != nil && linkedList.Value == linkedList.Next.Value {
		if linkedList.Next.Next != nil {
			linkedList.Next = linkedList.Next.Next
			RemoveDuplicatesFromLinkedList(linkedList)
		} else {
			linkedList.Next = nil
		}
	} else if linkedList.Next != nil {
		RemoveDuplicatesFromLinkedList(linkedList.Next)
	}

	return base
}

func main() {
	ll := &LinkedList{Value: 1}
	ll.Next = &LinkedList{Value: 1}
	ll.Next.Next = &LinkedList{Value: 3}
	ll.Next.Next.Next = &LinkedList{Value: 4}
	ll.Next.Next.Next.Next = &LinkedList{Value: 4}
	ll.Next.Next.Next.Next.Next = &LinkedList{Value: 4}
	ll.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 5}
	ll.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 6}
	ll.Next.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 6}

	result := RemoveDuplicatesFromLinkedList(ll)
	fmt.Println(result)
}
