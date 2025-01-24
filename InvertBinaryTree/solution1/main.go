package main

import "fmt"

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	temp := &BinaryTree{}

	temp = tree.Left
	tree.Left = tree.Right
	tree.Right = temp

	if tree.Left != nil {
		tree.Left.InvertBinaryTree()
	}
	if tree.Right != nil {
		tree.Right.InvertBinaryTree()
	}
}

func main() {
	tree := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
				},
				Right: &BinaryTree{
					Value: 9,
				},
			},
			Right: &BinaryTree{
				Value: 5,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
			},
			Right: &BinaryTree{
				Value: 7,
			},
		},
	}
	// tree := &BinaryTree{
	// 	Value: 1,
	// 	Left: &BinaryTree{
	// 		Value: 2,
	// 	},
	// }
	tree.InvertBinaryTree()
	fmt.Println(tree)
}
