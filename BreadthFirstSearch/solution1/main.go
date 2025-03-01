package main

import "fmt"

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	q := []*Node{}
	q = append(q, n)
	current := 0

	for {
		if len(q) == current {
			break
		}
		for _, v := range q[current:] {
			array = append(array, v.Name)
			q = append(q, v.Children...)
			current++
		}
	}

	return array
}

func main() {
	array := []string{}
	node := Node{
		Name: "A",
		Children: []*Node{
			{
				Name: "B",
				Children: []*Node{
					{
						Name: "E",
					},
					{
						Name: "F",
						Children: []*Node{
							{
								Name: "I",
							},
							{
								Name: "J",
							},
						},
					},
				},
			},
			{
				Name: "C",
			},
			{
				Name: "D",
				Children: []*Node{
					{
						Name: "G",
						Children: []*Node{
							{
								Name: "K",
							},
						},
					},
					{
						Name: "H",
					},
				},
			},
		},
	}
	fmt.Println(node.BreadthFirstSearch(array))
}
