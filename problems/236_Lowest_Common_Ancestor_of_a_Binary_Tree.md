# 236 Lowest Common Ancestor of a Binary Tree

- **Finished Date:** 2016-09-07
- **Tried:** 5
- **LeetCode Link:** [Lowest Common Ancestor of a Binary Tree | LeetCode OJ](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)
- **Tags:** [`Tree`](https://leetcode.com/tag/tree/)
- **Difficulty:** Medium
- **Pass Rate:** 29.09%

[C# Version](../C#/CSharp/236_LowestCommonAncestorOfABinaryTree.cs)

## Problem description

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes v and w as the lowest node in T that has both v and w as descendants (where we allow **a node to be a descendant of itself**).”

```C#
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int x) { val = x; }
 * }
 */
```

## Examples

```
        _______3______
       /              \
    ___5__          ___1__
   /      \        /      \
   6      _2       0       8
         /  \
         7   4
```

For example, the lowest common ancestor (LCA) of nodes `5` and `1` is `3`. Another example is LCA of nodes `5` and `4` is `5`, since a node can be a descendant of itself according to the LCA definition.
