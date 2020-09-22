package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func walk(root *BinaryTree, position int) int {
	sum := 0
	position++
	if root.Left != nil {
		sum += position + walk(root.Left, position)
	}
	if root.Right != nil {
		sum += position + walk(root.Right, position)
	}
	return sum
}

func NodeDepths(root *BinaryTree) int {
	sum := walk(root, 0)
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
