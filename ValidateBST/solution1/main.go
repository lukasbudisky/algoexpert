package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) ValidateBst() bool {
	result := true
	if tree.Left != nil {
		if tree.Left.Left != nil {
			if !(tree.Left.Left.Value < tree.Value && tree.Left.Value < tree.Value) {
				return false
			}
		}
		if tree.Left.Right != nil {
			if !(tree.Left.Right.Value < tree.Value && tree.Left.Value < tree.Value) {
				return false
			}
			if tree.Left.Right.Right != nil {
				if !(tree.Left.Right.Right.Value < tree.Value) {
					return false
				}
			}
		}
		if tree.Left.Value < tree.Value {
			result = tree.Left.ValidateBst()
		} else {
			return false
		}
	}

	if !result {
		return false
	}

	if tree.Right != nil {
		if tree.Right.Left != nil {
			if !(tree.Right.Left.Value >= tree.Value && tree.Right.Left.Value < tree.Right.Value) {
				return false
			}
		}
		if tree.Right.Right != nil {
			if !(tree.Right.Right.Value >= tree.Right.Value && tree.Right.Right.Value >= tree.Value) {
				return false
			}
		}
		if tree.Right.Value >= tree.Value {
			result = tree.Right.ValidateBst()
		} else {
			return false
		}
	}
	return result
}

func main() {
	// true
	t := &BST{
		Value: 5000,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Right: &BST{
					Value: 3,
				},
				Left: &BST{
					Value: 1,
					Right: &BST{
						Value: 1,
						Right: &BST{
							Value: 1,
							Right: &BST{
								Value: 1,
								Right: &BST{
									Value: 1,
									Right: &BST{
										Value: 1,
									},
								},
							},
						},
					},
				},
			},
			Right: &BST{
				Value: 15,
				Left: &BST{
					Value: 5,
				},
				Right: &BST{
					Value: 22,
					Right: &BST{
						Value: 502,
						Left: &BST{
							Value: 204,
							Left: &BST{
								Value: 203,
							},
							Right: &BST{
								Value: 205,
								Right: &BST{
									Value: 207,
									Left: &BST{
										Value: 206,
									},
									Right: &BST{
										Value: 208,
									},
								},
							},
						},
					},
				},
			},
		},
		Right: &BST{
			Value: 55000,
		},
	}

	// true
	// t := &BST{
	// 	Value: 10,
	// 	Left: &BST{
	// 		Value: 5,
	// 		Left: &BST{
	// 			Value: 2,
	// 			Left: &BST{
	// 				Value: 1,
	// 			},
	// 		},
	// 		Right: &BST{
	// 			Value: 5,
	// 		},
	// 	},
	// 	Right: &BST{
	// 		Value: 15,
	// 		Left: &BST{
	// 			Value: 13,
	// 			Right: &BST{
	// 				Value: 14,
	// 			},
	// 		},
	// 		Right: &BST{
	// 			Value: 22,
	// 		},
	// 	},
	// }

	// false
	// t := &BST{
	// 	Value: 10,
	// 	Left: &BST{
	// 		Value: 5,
	// 		Left: &BST{
	// 			Value: 2,
	// 			Left: &BST{
	// 				Value: 1,
	// 			},
	// 		},
	// 		Right: &BST{
	// 			Value: 5,
	// 			Right: &BST{
	// 				Value: 11,
	// 			},
	// 		},
	// 	},
	// 	Right: &BST{
	// 		Value: 15,
	// 		Right: &BST{
	// 			Value: 22,
	// 		},
	// 	},
	// }

	// false
	// t := &BST{
	// 	Value: 10,
	// 	Left: &BST{
	// 		Value: 5,
	// 		Left: &BST{
	// 			Value: 2,
	// 			Left: &BST{
	// 				Value: 1,
	// 			},
	// 		},
	// 		Right: &BST{
	// 			Value: 5,
	// 		},
	// 	},
	// 	Right: &BST{
	// 		Value: 15,
	// 		Left: &BST{
	// 			Value: 13,
	// 			Right: &BST{
	// 				Value: 16,
	// 			},
	// 		},
	// 		Right: &BST{
	// 			Value: 22,
	// 		},
	// 	},
	// }

	// false
	// t := &BST{
	// 	Value: 10,
	// 	Left: &BST{
	// 		Value: 5,
	// 		Right: &BST{
	// 			Value: 10,
	// 		},
	// 	},
	// 	Right: &BST{
	// 		Value: 15,
	// 	},
	// }

	fmt.Println(t.ValidateBst())

}
