package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if tree.Value > value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else if tree.Value <= value {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func (tree *BST) Contains(value int) bool {
	result := false
	if value == tree.Value {
		return true
	} else if value >= tree.Value && tree.Right != nil {
		result = tree.Right.Contains(value)
	} else if value < tree.Value && tree.Left != nil {
		result = tree.Left.Contains(value)
	}
	return result
}

func (tree *BST) Search() int {
	result := tree.Value
	if tree.Left != nil {
		if tree.Left.Left == nil && tree.Left.Right == nil {
			result = tree.Left.Value
			tree.Left = nil
			return result
		} else {
			result = tree.Left.Search()
		}
	}
	return result
}

func (tree *BST) Remove(value int) *BST {
	if value == tree.Value && (tree.Right != nil) {
		if tree.Right.Left == nil {
			tree.Value = tree.Right.Value
			tree.Right = tree.Right.Right
		} else {
			tree.Value = tree.Right.Search()
		}
	} else if value == tree.Value && tree.Right == nil && tree.Left != nil {
		tree.Value = tree.Left.Value
		tree.Left = tree.Left.Left
	} else if value >= tree.Value && tree.Right != nil {
		if tree.Right.Left == nil && tree.Right.Right == nil {
			tree.Right = nil
		} else {
			tree.Right.Remove(value)
		}
	} else if value < tree.Value && tree.Left != nil {
		if tree.Left.Left == nil && tree.Left.Right == nil {
			tree.Left = nil
		} else {
			tree.Left.Remove(value)
		}
	}
	return tree
}

func main() {
	// result := &BST{
	// 	Value: 10,
	// 	Right: &BST{
	// 		Value: 15,
	// 		Right: &BST{
	// 			Value: 22,
	// 		},
	// 		Left: &BST{
	// 			Value: 13,
	// 			Right: &BST{
	// 				Value: 14,
	// 			},
	// 		},
	// 	},
	// 	Left: &BST{
	// 		Value: 5,
	// 		Right: &BST{
	// 			Value: 5,
	// 		},
	// 		Left: &BST{
	// 			Value: 2,
	// 			Left: &BST{
	// 				Value: 1,
	// 			},
	// 		},
	// 	},
	// }

	// result := &BST{
	// 	Value: 10,
	// 	Left: &BST{
	// 		Value: 5,
	// 	},
	// 	// Right: &BST{
	// 	// 	Value: 15,
	// 	// },
	// }

	// result := &BST{
	// 	Value: 1,
	// 	Right: &BST{
	// 		Value: 2,
	// 	},
	// 	// Right: &BST{
	// 	// 	Value: 15,
	// 	// },
	// }

	result := &BST{
		Value: 1,
		Left: &BST{
			Value: -2,
		},
	}

	result.Insert(-3)
	result.Insert(-4)
	result.Remove(1)
	fmt.Println("end")
}

func NewBST(value int) *BST {
	return &BST{Value: value}
}
