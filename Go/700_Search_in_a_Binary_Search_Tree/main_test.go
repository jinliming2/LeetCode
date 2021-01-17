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

func TestSearchBST(t *testing.T) {
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
		}, 2}: "2 1 null null 3 null null",
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
		}, 5}: "null",
	} {
		value := searchBST(input.root, input.val)
		var answer []string
		printTree(value, &answer)
		if strings.Join(answer, " ") != output {
			t.Errorf("Expected searchBST(%v, %d) == %s, But got %s", input.root, input.val, output, answer)
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
