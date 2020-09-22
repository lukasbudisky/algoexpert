package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// FindClosestValue function
// Time:
// n - is number of nodes in tree
// Avg: O(log(n))
// Worst: O(n)
//
// Space:
// O(1) space
func (tree *BST) FindClosestValue(target int) int {
	value := tree.Value
	for {
		switch {
		case tree.Right != nil && tree.Value < target:
			tree = tree.Right
			fmt.Println(tree.Value)
			if absolute(value-target) > absolute(tree.Value-target) {
				value = tree.Value
			}
		case tree.Left != nil && tree.Value > target:
			tree = tree.Left
			fmt.Println(tree.Value)
			if absolute(value-target) > absolute(tree.Value-target) {
				value = tree.Value
			}
		default:
			return value
		}
	}
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
