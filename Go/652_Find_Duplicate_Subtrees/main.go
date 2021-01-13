package main

import "strconv"

// Runtime: 12 ms, faster than 87.39% of Go online submissions for Find Duplicate Subtrees.
// Memory Usage: 10.6 MB, less than 42.34% of Go online submissions for Find Duplicate Subtrees.
// https://leetcode.com/submissions/detail/442492818/
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	trees := map[string]int{}
	list := []*TreeNode{}
	traverse(root.Left, trees, &list)
	traverse(root.Right, trees, &list)
	return list
}

func traverse(root *TreeNode, trees map[string]int, list *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	tree := strconv.Itoa(root.Val) + " " + traverse(root.Left, trees, list) + " " + traverse(root.Right, trees, list)
	trees[tree]++
	if trees[tree] == 2 {
		*list = append(*list, root)
	}
	return tree
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
