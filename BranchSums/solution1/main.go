package main

import "fmt"

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func walk(root *BinaryTree, result *[]int, sum int) int {
	sum += root.Value
	if root.Left != nil {
		sum += walk(root.Left, result, sum)
	}
	if root.Right != nil {
		sum += walk(root.Right, result, sum)
	}
	if root.Left == nil && root.Right == nil {
		*result = append(*result, sum)
	}
	sum = 0
	return sum
}

func BranchSums(root *BinaryTree) []int {
	result := []int{}
	sum := 0
	walk(root, &result, sum)
	return result
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
				Left: &BinaryTree{
					Value: 10,
					Left:  nil,
					Right: nil,
				},
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

	fmt.Println(BranchSums(&tree))
}
