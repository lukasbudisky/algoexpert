package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func createArray(tree *BST) []int {
	array := []int{}
	if tree.Left != nil {
		array = append(array, createArray(tree.Left)...)
	}
	array = append(array, tree.Value)
	if tree.Right != nil {
		array = append(array, createArray(tree.Right)...)
	}
	return array
}

func findNode(tree *BST, v int) *BST {
	var result *BST
	if tree != nil && v == tree.Value {
		return tree
	}
	if tree.Left != nil {
		result = findNode(tree.Left, v)
		if result != nil {
			return result
		}
	}
	if tree.Right != nil {
		result = findNode(tree.Right, v)
		if result != nil {
			return result
		}
	}
	return result
}

func RepairBst(tree *BST) *BST {
	array := createArray(tree)
	result := []int{}

	for i, v := range array {
		if i+1 <= len(array)-1 {
			if v > array[i+1] {
				result = append(result, v)
				break
			}
		}
	}
	for i := len(array) - 1; i > 0; i-- {
		if array[i] < array[i-1] {
			result = append(result, array[i])
			break
		}
	}

	fmt.Println(result)

	if len(result) == 2 {
		treeFirst := findNode(tree, result[0])
		treeSecond := findNode(tree, result[1])
		treeFirst.Value, treeSecond.Value = treeSecond.Value, treeFirst.Value
	}

	return tree
}

func main() {
	// tree := &BST{Value: 10}
	// tree.Left = &BST{Value: 7}
	// tree.Left.Left = &BST{Value: 3}
	// tree.Left.Right = &BST{Value: 12}
	// tree.Left.Left.Left = &BST{Value: 2}
	// tree.Right = &BST{Value: 20}
	// tree.Right.Left = &BST{Value: 8}
	// tree.Right.Right = &BST{Value: 22}
	// tree.Right.Left.Right = &BST{Value: 14}

	tree := &BST{Value: 1}
	tree.Right = &BST{Value: 3}
	tree.Right.Right = &BST{Value: 2}
	tree = RepairBst(tree)
	fmt.Println(tree)
}
