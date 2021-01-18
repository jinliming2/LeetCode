package main

import (
	"strconv"
	"strings"
	"testing"
)

type inputData struct {
	root *TreeNode
	key  int
}

func TestDeleteNode(t *testing.T) {
	for input, output := range map[*inputData]string{
		{&TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 7,
				},
			},
		}, 3}: "2 4 5 6 7",
		{&TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 7,
				},
			},
		}, 0}: "2 3 4 5 6 7",
		{nil, 0}: "",
	} {
		value := deleteNode(input.root, input.key)
		var answer []string
		printTree(value, &answer)
		if strings.Join(answer, " ") != output {
			t.Errorf("Expected deleteNode(%v, %d) == %s, But got %s", input.root, input.key, output, answer)
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
