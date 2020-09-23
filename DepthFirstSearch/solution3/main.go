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

func (n *Node) walk(array []string) []string {
	array = append(array, n.Name)
	for _, v := range n.Children {
		array = v.walk(array)
	}
	return array
}

func (n *Node) DepthFirstSearch(array []string) []string {
	return n.walk(array)
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
	fmt.Println(n.DepthFirstSearch(nil))
}
