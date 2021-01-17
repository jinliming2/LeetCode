package main

import (
	"strconv"
	"strings"
	"testing"
)

type inputData struct {
	root *TreeNode
	val  int
}

func TestInsertIntoBST(t *testing.T) {
	for input, output := range map[*inputData]string{
		{&TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 7,
			},
		}, 5}: "1 2 3 4 5 7",
		{&TreeNode{
			Val: 40,
			Left: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 10,
				},
				Right: &TreeNode{
					Val: 30,
				},
			},
			Right: &TreeNode{
				Val: 60,
				Left: &TreeNode{
					Val: 50,
				},
				Right: &TreeNode{
					Val: 70,
				},
			},
		}, 25}: "10 20 25 30 40 50 60 70",
		{&TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 7,
			},
		}, 5}: "1 2 3 4 5 7",
	} {
		value := insertIntoBST(input.root, input.val)
		var answer []string
		printTree(value, &answer)
		if strings.Join(answer, " ") != output {
			t.Errorf("Expected insertIntoBST(%v, %d) == %s, But got %s", input.root, input.val, output, answer)
		}
	}
}

func printTree(root *TreeNode, res *[]string) {
	if root == nil {
		return
	}
	printTree(root.Left, res)
	*res = append(*res, strconv.Itoa(root.Val))
	printTree(root.Right, res)
}
