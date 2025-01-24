package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func (tree *BST) walk(array []int) *BST {
	c := len(array)
	if len(array) == 1 {
		tree.Insert(array[0])
		return tree
	}
	if c%2 == 0 {
		c = c/2 - 1
		if c >= 0 {
			tree.Insert(array[c])
		} else {
			return tree
		}
	} else {
		c = c / 2
		tree.Insert(array[c])
	}

	tree.walk(array[0:c])
	tree.walk(array[c+1:])
	return tree
}

func MinHeightBST(array []int) *BST {
	tree := &BST{}
	c := len(array)

	if c%2 == 0 {
		c = c/2 - 1
		tree.Value = array[c]
	} else {
		c = c / 2
		tree.Value = array[c]
	}
	tree.walk(array[0:c])
	tree.walk(array[c+1:])
	return tree
}

func main() {
	// array := []int{1, 2, 5, 7, 10, 13, 14, 15, 22}
	// array := []int{1, 2, 5, 7}
	array := []int{1}
	tree := MinHeightBST(array)
	fmt.Printf("%v", tree)
}
