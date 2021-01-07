package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestConnect(t *testing.T) {
	for input, output := range map[*Node]string{
		{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
				},
				Right: &Node{
					Val: 5,
				},
			},
			Right: &Node{
				Val: 3,
				Left: &Node{
					Val: 6,
				},
				Right: &Node{
					Val: 7,
				},
			},
		}: "null 3 5 6 null 7 null",
	} {
		connect(input)
		var value []string
		if printTreeNext(input, &value); strings.Join(value, " ") != output {
			t.Errorf("connect got wrong answer %v", value)
		}
	}
}

func printTreeNext(root *Node, res *[]string) {
	if root == nil {
		return
	}
	if root.Next == nil {
		*res = append(*res, "null")
	} else {
		*res = append(*res, strconv.Itoa(root.Next.Val))
	}
	printTreeNext(root.Left, res)
	printTreeNext(root.Right, res)
}
