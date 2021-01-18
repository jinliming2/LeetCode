package main

import (
	"strconv"
	"strings"
)

// Runtime: 444 ms, faster than 9.55% of Go online submissions for Serialize and Deserialize Binary Tree.
// Memory Usage: 76.3 MB, less than 92.97% of Go online submissions for Serialize and Deserialize Binary Tree.
// https://leetcode.com/submissions/detail/444616075/

// Codec :
type Codec struct {
}

// Constructor :
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (codec *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + " " + codec.serialize(root.Left) + " " + codec.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	list := strings.Split(data, " ")
	index := 0
	length := len(list)

	var traverse func(*[]string, *int) *TreeNode
	traverse = func(list *[]string, index *int) *TreeNode {
		if *index >= length {
			return nil
		}
		char := (*list)[*index]
		*index++
		if char == "#" {
			return nil
		}
		val, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		root := &TreeNode{Val: val}
		root.Left = traverse(list, index)
		root.Right = traverse(list, index)
		return root
	}

	return traverse(&list, &index)
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
