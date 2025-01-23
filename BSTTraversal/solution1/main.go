package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(array []int) []int {
	if tree.Left != nil {
		array = tree.Left.InOrderTraverse(array)
	}

	if tree.Right != nil {
		array = append(array, tree.Value)
		array = tree.Right.InOrderTraverse(array)
	} else {
		array = append(array, tree.Value)
	}

	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	if tree.Left != nil {
		array = append(array, tree.Value)
		array = tree.Left.PreOrderTraverse(array)
		if tree.Right != nil {
			array = tree.Right.PreOrderTraverse(array)
		}
	} else {
		array = append(array, tree.Value)
		if tree.Right != nil {
			array = tree.Right.PreOrderTraverse(array)
		}
	}
	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	if tree.Left != nil {
		array = tree.Left.PostOrderTraverse(array)
	}

	if tree.Right != nil {
		array = tree.Right.PostOrderTraverse(array)
	}

	array = append(array, tree.Value)

	return array
}

func main() {
	t := &BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left: &BST{
					Value: 1,
				},
			},
			Right: &BST{
				Value: 5,
			},
		},
		Right: &BST{
			Value: 15,
			Right: &BST{
				Value: 22,
			},
		},
	}

	result := []int{}
	fmt.Printf("InOrderTraverse: %v\n", t.InOrderTraverse(result))
	result = []int{}
	fmt.Printf("PreOrderTraverse: %v\n", t.PreOrderTraverse(result))
	result = []int{}
	fmt.Printf("PostOrderTraverse: %v\n", t.PostOrderTraverse(result))
}
