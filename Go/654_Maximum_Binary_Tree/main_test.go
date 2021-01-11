package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	for input, output := range map[*[]int]string{
		{3, 2, 1, 6, 0, 5}: "6 3 null 2 null 1 null null 5 0 null null null",
		{3, 2, 1}:          "3 null 2 null 1 null null",
	} {
		value := constructMaximumBinaryTree(*input)
		var answer []string
		printTree(value, &answer)
		if value := strings.Join(answer, " "); value != output {
			t.Errorf("Expected constructMaximumBinaryTree(%v) == %s, But got %s", input, output, value)
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
