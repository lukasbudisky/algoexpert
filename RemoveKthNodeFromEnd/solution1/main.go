package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	f, s := head, head
	i := 0
	c := 1
	for {
		if i == k {
			break
		}
		if s.Next != nil {
			s = s.Next
			c++
		} else {
			break
		}
		i++
	}
	for {
		if s.Next != nil {
			s = s.Next
			f = f.Next
			c++
		} else {
			if c == k {
				*f = *f.Next
			} else {
				f.Next = f.Next.Next
			}
			break
		}
	}
}

func main() {
	ll := &LinkedList{Value: 0}
	ll.Next = &LinkedList{Value: 1}
	ll.Next.Next = &LinkedList{Value: 2}
	ll.Next.Next.Next = &LinkedList{Value: 3}
	ll.Next.Next.Next.Next = &LinkedList{Value: 4}
	ll.Next.Next.Next.Next.Next = &LinkedList{Value: 5}
	ll.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 6}
	ll.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 7}
	ll.Next.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 8}
	ll.Next.Next.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 9}
	k := 10

	RemoveKthNodeFromEnd(ll, k)
	fmt.Println(ll)
}
