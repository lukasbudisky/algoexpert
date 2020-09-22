package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	stack := []*BinaryTree{root}
	sum := 0
	pos := 0
	for {
		ts := stack
		pos++
		if len(ts) == 0 {
			break
		}
		stack = []*BinaryTree{}
		for i := range ts {
			t := ts[i]
			if t.Left != nil {
				sum += pos
				stack = append(stack, t.Left)
			}
			if t.Right != nil {
				sum += pos
				stack = append(stack, t.Right)
			}
		}
	}
	return sum
}

func main() {
	tree := BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
					Left:  nil,
					Right: nil,
				},
				Right: &BinaryTree{
					Value: 9,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BinaryTree{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryTree{
				Value: 7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(NodeDepths(&tree))
}
