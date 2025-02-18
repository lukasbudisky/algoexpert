package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func findDepth(node *BST) int {
	result := 0
	if node.Left != nil {
		result++
		result += findDepth(node.Left)
	}
	if node.Right != nil {
		result++
		result += findDepth(node.Right)
	}
	return result
}

func ValidateThreeNodes(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	i := 0
	var values []*BST
	var root *BST

	one := findDepth(nodeOne)
	two := findDepth(nodeTwo)
	three := findDepth(nodeThree)

	if one > two && one > three {
		root = nodeOne
		values = append(values, nodeOne, nodeTwo, nodeThree)
	} else if two > three && two > one {
		root = nodeTwo
		values = append(values, nodeTwo, nodeOne, nodeThree)
	} else if three > two && three > one {
		root = nodeThree
		values = append(values, nodeThree, nodeTwo, nodeOne)
	}

	for {
		if i <= len(values)-1 && root.Value == values[i].Value {
			i++
			if i > len(values)-1 {
				return true
			}
		} else if i <= len(values)-1 && root.Value > values[i].Value {
			if root.Left != nil {
				root = root.Left
			} else {
				return false
			}
		} else if i <= len(values)-1 && root.Value <= values[i].Value {
			if root.Right != nil {
				root = root.Right
			} else {
				return false
			}
		} else {
			return false
		}
	}
}

func main() {
	// nodeOne := &BST{
	// 	Value: 5,
	// 	Left: &BST{
	// 		Value: 2,
	// 		Left: &BST{
	// 			Value: 1,
	// 			Left: &BST{
	// 				Value: 0,
	// 			},
	// 		},
	// 		Right: &BST{
	// 			Value: 4,
	// 			Left: &BST{
	// 				Value: 3,
	// 			},
	// 		},
	// 	},
	// 	Right: &BST{
	// 		Value: 7,
	// 		Left: &BST{
	// 			Value: 6,
	// 		},
	// 		Right: &BST{
	// 			Value: 8,
	// 		},
	// 	},
	// }

	// nodeTwo := &BST{
	// 	Value: 2,
	// 	Left: &BST{
	// 		Value: 1,
	// 		Left: &BST{
	// 			Value: 0,
	// 		},
	// 	},
	// 	Right: &BST{
	// 		Value: 4,
	// 		Left: &BST{
	// 			Value: 3,
	// 		},
	// 	},
	// }

	// nodeThree := &BST{
	// 	Value: 3,
	// }

	// nodeOne := &BST{
	// 	Value: 1,
	// }

	// nodeTwo := &BST{
	// 	Value: 2,
	// 	Left: &BST{
	// 		Value: 1,
	// 	},
	// 	Right: &BST{
	// 		Value: 3,
	// 		Right: &BST{
	// 			Value: 4,
	// 		},
	// 	},
	// }

	// nodeThree := &BST{
	// 	Value: 3,
	// 	Right: &BST{
	// 		Value: 4,
	// 	},
	// }

	// true
	// nodeOne := &BST{
	// 	Value: 1,
	// }

	// nodeTwo := &BST{
	// 	Value: 2,
	// 	Left: &BST{
	// 		Value: 1,
	// 	},
	// }

	// nodeThree := &BST{
	// 	Value: 3,
	// 	Left: &BST{
	// 		Value: 2,
	// 		Left: &BST{
	// 			Value: 1,
	// 		},
	// 	},
	// }

	// false
	// nodeTwo := &BST{
	// 	Value: 1,
	// }

	// nodeOne := &BST{
	// 	Value: 2,
	// 	Left: &BST{
	// 		Value: 1,
	// 	},
	// }

	// nodeThree := &BST{
	// 	Value: 3,
	// 	Left: &BST{
	// 		Value: 2,
	// 		Left: &BST{
	// 			Value: 1,
	// 		},
	// 	},
	// }

	// root := &BST{Value: 5}
	// root.Left = &BST{Value: 2}
	// root.Right = &BST{Value: 7}
	// root.Left.Left = &BST{Value: 1}
	// root.Left.Right = &BST{Value: 4}
	// root.Right.Left = &BST{Value: 6}
	// root.Right.Right = &BST{Value: 8}
	// root.Left.Left.Left = &BST{Value: 0}
	// root.Left.Right.Left = &BST{Value: 3}

	// nodeOne := root
	// nodeTwo := root.Left
	// nodeThree := root.Left.Right.Left

	root := &BST{Value: 6}
	root.Left = &BST{Value: 4}
	root.Left.Left = &BST{Value: 2}
	root.Left.Right = &BST{Value: 5}
	root.Left.Left.Left = &BST{Value: 1}
	root.Left.Left.Right = &BST{Value: 3}
	root.Right = &BST{Value: 8}
	root.Right.Left = &BST{Value: 7}
	root.Right.Right = &BST{Value: 9}

	nodeOne := root.Right.Right
	nodeTwo := root.Right
	nodeThree := root

	fmt.Println(ValidateThreeNodes(nodeOne, nodeTwo, nodeThree))
}
