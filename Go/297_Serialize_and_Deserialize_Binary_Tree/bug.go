// +build ignore

package main

// Runtime: 412 ms, faster than 15.58% of Go online submissions for Serialize and Deserialize Binary Tree.
// Memory Usage: 99.6 MB, less than 80.40% of Go online submissions for Serialize and Deserialize Binary Tree.
// https://leetcode.com/submissions/detail/444603104/

// Codec :
type Codec struct {
}

// Constructor :
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (codec *Codec) serialize(root *TreeNode) *TreeNode {
	return root
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data *TreeNode) *TreeNode {
	return data
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
