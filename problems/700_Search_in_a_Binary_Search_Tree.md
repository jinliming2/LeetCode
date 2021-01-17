# 700 Search in a Binary Search Tree

- **Finished Date:** 2021-01-17
- **Tried:** 1
- **LeetCode Link:** [Search in a Binary Search Tree - LeetCode](https://leetcode.com/problems/search-in-a-binary-search-tree/)
- **Tags:** [`Tree`](https://leetcode.com/tag/tree/)
- **Difficulty:** Easy
- **Pass Rate:** 73.3865%

[Go Version](../Go/700_Search_in_a_Binary_Search_Tree/main.go)

## Problem description

You are given the `root` of a binary search tree (BST) and an integer `val`.

Find the node in the BST that the node's value equals `val` and return the subtree rooted with that node. If such a node does not exist, return `null`.

### Constraints

- The number of nodes in the tree is in the range `[1, 5000]`.
- 1 <= `Node.val` <= 10<sup>7</sup>
- root is a binary search tree.
- 1 <= `val` <= 10<sup>7</sup>

## Examples

### Example 1

![](./assets/700.Search_in_a_Binary_Search_Tree_1.jpg)

```
Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]
```

### Example 2

![](./assets/700.Search_in_a_Binary_Search_Tree_2.jpg)

```
Input: root = [4,2,7,1,3], val = 5
Output: []
```
