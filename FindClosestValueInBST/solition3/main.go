package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func navigate(tree *BST, target int, result *int) *BST {
	switch {
	case tree.Left != nil && tree.Value > target:
		if absolute(*result-target) > absolute(tree.Left.Value-target) {
			*result = tree.Left.Value
		}
		return tree.Left
	case tree.Right != nil && tree.Value < target:
		if absolute(*result-target) > absolute(tree.Right.Value-target) {
			*result = tree.Right.Value
		}
		return tree.Right
	default:
		return nil
	}
}

func (tree *BST) FindClosestValue(target int) int {
	var result int = nil
	for {
		tree = navigate(tree, target, &result)
		if tree == nil {
			break
		}
	}
	return result
}

func main() {
	tree := BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left: &BST{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &BST{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BST{
			Value: 15,
			Left: &BST{
				Value: 13,
				// Left:  nil,
				// Left: &BST{Value: 12},
				Left: &BST{Value: 8},
				Right: &BST{
					Value: 14,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BST{Value: 22},
		},
	}

	fmt.Println(tree.FindClosestValue(12))
}
