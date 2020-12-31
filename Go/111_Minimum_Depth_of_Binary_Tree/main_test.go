package main

import (
	"strconv"
	"testing"
)

var testCase = map[*TreeNode]int{
	{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}: 2,
	{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
		},
	}: 5,
	nil: 0,
	{
		Val: 5,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			Right: &TreeNode{
				Val: -6,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: -4,
			Left: &TreeNode{
				Val: -9,
				Right: &TreeNode{
					Val: 0,
					Right: &TreeNode{
						Val: -7,
						Left: &TreeNode{
							Val: -4,
							Left: &TreeNode{
								Val: 1,
								Right: &TreeNode{
									Val: -4,
								},
							},
						},
					},
				},
			},
		},
	}: 4,
	{
		Val: 0,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: -1,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}: 4,
}

func TestMinDepthBFS(t *testing.T) {
	for input, output := range testCase {
		if value := minDepth(input); value != output {
			tree := []string{}
			printTree(input, &tree)
			t.Errorf("Expected minDepth(%v) == %d, But got %d", tree, output, value)
		}
	}
}

func TestMinDepthDFS(t *testing.T) {
	for input, output := range testCase {
		if value := minDepthDFS(input); value != output {
			tree := []string{}
			printTree(input, &tree)
			t.Errorf("Expected minDepthDFS(%v) == %d, But got %d", tree, output, value)
		}
	}
}

func printTree(root *TreeNode, res *[]string) {
	if root == nil {
		*res = append(*res, "null")
		return
	}
	*res = append(*res, strconv.Itoa(root.Val))
	printTree(root.Left, res)
	printTree(root.Right, res)
}
