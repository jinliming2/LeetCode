# 111 Minimum Depth of Binary Tree

- **Finished Date:** 2020-12-30
- **Tried:** 2(BFS) + 3(DFS)
- **LeetCode Link:** [Minimum Depth of Binary Tree - LeetCode](https://leetcode.com/problems/minimum-depth-of-binary-tree/)
- **Tags:** [`Tree`](https://leetcode.com/tag/tree/), [`Depth-first Search`](https://leetcode.com/tag/depth-first-search/), [`Breadth-first Search`](https://leetcode.com/tag/breadth-first-search/)
- **Difficulty:** Easy
- **Pass Rate:** 39.07%

[Go BFS Version](../Go/111_Minimum_Depth_of_Binary_Tree/bfs.go)

[Go DFS Version](../Go/111_Minimum_Depth_of_Binary_Tree/dfs.go)

## Problem description

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

**Note:** A leaf is a node with no children.

```Go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
```

### Constraints

- The number of nodes in the tree is in the range [0, 10<sup>5</sup>].
- -1000 <= Node.val <= 1000

## Examples

### Example 1

![tree](./assets/111.Minimum_Depth_of_Binary_Tree.jpg)

```
   ____3___
  /        \
 9       __20__
        /      \
       15       7

Input: root = [3,9,20,null,null,15,7]
Output: 2
```

### Example 2

```
Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5
```
