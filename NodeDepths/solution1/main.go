package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func walk(root *BinaryTree, sum int, position int) int {
	fmt.Println(root.Value)
	position++
	if root.Left != nil {
		sum += position
		sum = walk(root.Left, sum, position)
	}
	if root.Right != nil {
		sum += position
		sum = walk(root.Right, sum, position)
	}
	return sum
}

func NodeDepths(root *BinaryTree) int {
	sum := walk(root, 0, 0)
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
