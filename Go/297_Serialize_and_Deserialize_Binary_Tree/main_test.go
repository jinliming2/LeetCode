package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestSerializeDeserialize(t *testing.T) {
	for root, output := range map[*TreeNode]string{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		}: "1 2 null null 3 4 null null 5 null null",
		nil: "null",
		{
			Val: 1,
		}: "1 null null",
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}: "1 2 null null null",
	} {

		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(root)
		ans := deser.deserialize(data)

		var answer []string
		printTree(ans, &answer)
		if strings.Join(answer, " ") != output {
			t.Errorf("Expected SerializeDeserialize(%v) == %s, But got %s", root, output, answer)
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
