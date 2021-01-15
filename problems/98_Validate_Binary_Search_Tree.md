# 98 Validate Binary Search Tree

- **Finished Date:** 2021-01-16
- **Tried:** 4
- **LeetCode Link:** [Validate Binary Search Tree - LeetCode](https://leetcode.com/problems/validate-binary-search-tree/)
- **Tags:** [`Tree`](https://leetcode.com/tag/tree/), [`Depth-first Search`](https://leetcode.com/tag/depth-first-search/), [`Recursion`](https://leetcode.com/tag/recursion/)
- **Difficulty:** Medium
- **Pass Rate:** 28.5779%

[Go Version](../Go/98_Validate_Binary_Search_Tree/main.go)

## Problem description

Given the `root` of a binary tree, *determine if it is a valid binary search tree (BST)*.

A **valid BST** is defined as follows:

- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

### Constraints

- The number of nodes in the tree is in the range [1, 10<sup>4</sup>].
- -2<sup>31</sup> <= Node.val <= 2<sup>31</sup> - 1

## Examples

### Example 1

![](./assets/98.Validate_Binary_Search_Tree_1.jpg)

```
Input: root = [2,1,3]
Output: true
```

### Example 2

![](./assets/98.Validate_Binary_Search_Tree_2.jpg)

```
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
```
