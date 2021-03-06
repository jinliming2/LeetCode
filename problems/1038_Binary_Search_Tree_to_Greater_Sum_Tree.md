# 1038 Binary Search Tree to Greater Sum Tree

- **Finished Date:** 2021-01-15
- **Tried:** 1
- **LeetCode Link:** [Binary Search Tree to Greater Sum Tree - LeetCode](https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/)
- **Tags:** [`Binary Search Tree`](https://leetcode.com/tag/binary-search-tree/)
- **Difficulty:** Medium
- **Pass Rate:** 82.02%

[Go Version](../Go/1038_Binary_Search_Tree_to_Greater_Sum_Tree/main.go)

## Problem description

Given the `root` of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

As a reminder, a *binary search tree* is a tree that satisfies these constraints:

- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

**Note:** This question is the same as 538: [538 Convert BST to Greater Tree](./538_Convert_BST_to_Greater_Tree.md)

### Constraints

- The number of nodes in the tree is in the range `[0, 100]`.
- `0 <= Node.val <= 100`
- All the values in the tree are **unique**.
- `root` is guaranteed to be a valid binary search tree.

## Examples

### Example 1

![](./assets/1038.Binary_Search_Tree_to_Greater_Sum_Tree.png)

```
Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
```

### Example 2

```
Input: root = [0,null,1]
Output: [1,null,1]
```

### Example 3

```
Input: root = [1,0,2]
Output: [3,3,2]
```

### Example 4

```
Input: root = [3,2,4,1]
Output: [7,9,4,10]
```
