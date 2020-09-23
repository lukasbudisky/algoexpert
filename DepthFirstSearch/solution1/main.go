package main

import "fmt"

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func walk(node *Node, array *[]string) {
	for _, v := range node.Children {
		*array = append(*array, v.Name)
		walk(v, array)
	}
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// return walk(n, &array)
	array = append(array, n.Name)
	walk(n, &array)
	return array
}

func main() {
	n := Node{
		Name: "A",
		Children: []*Node{
			&Node{
				Name: "B",
				Children: []*Node{
					&Node{
						Name:     "E",
						Children: nil,
					},
					&Node{
						Name: "F",
						Children: []*Node{
							&Node{
								Name:     "I",
								Children: nil,
							},
							&Node{
								Name:     "J",
								Children: nil,
							},
						},
					},
				},
			},
			&Node{
				Name:     "C",
				Children: nil,
			},
			&Node{
				Name: "D",
				Children: []*Node{
					&Node{
						Name: "G",
						Children: []*Node{
							&Node{
								Name:     "K",
								Children: nil,
							},
						},
					},
					&Node{
						Name:     "H",
						Children: nil,
					},
				},
			},
		},
	}
	a := []string{}
	fmt.Println(n.DepthFirstSearch(a))
}
