package main

import "testing"

type inputData struct {
	root, p, q *TreeNode
}

func TestLowestCommonAncestor(t *testing.T) {
	for input, output := range map[inputData]int{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			p: &TreeNode{Val: 5},
			q: &TreeNode{Val: 1},
		}: 3,
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			p: &TreeNode{Val: 5},
			q: &TreeNode{Val: 4},
		}: 5,
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			p: &TreeNode{Val: 1},
			q: &TreeNode{Val: 2},
		}: 1,
	} {
		if value := lowestCommonAncestor(input.root, input.p, input.q); value.Val != output {
			t.Errorf("Expected lowestCommonAncestor(%v, %d, %d) == %d, But got %d", input.root, input.p.Val, input.q.Val, output, value.Val)
		}
	}
}
