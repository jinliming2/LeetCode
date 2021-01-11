package main

import (
	"strconv"
	"strings"
	"testing"
)

type inputData struct {
	preorder []int
	inorder  []int
}

func TestBuildTree(t *testing.T) {
	for input, output := range map[*inputData]string{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}}: "3 9 null null 20 15 null null 7 null null",
		{[]int{1, 2, 3}, []int{2, 3, 1}}:                 "1 2 null 3 null null null",
	} {
		value := buildTree(input.preorder, input.inorder)
		var answer []string
		printTree(value, &answer)
		if value := strings.Join(answer, " "); value != output {
			t.Errorf("Expected buildTree(%v, %v) == %s, But got %s", input.preorder, input.inorder, output, value)
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
