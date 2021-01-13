# 652 Find Duplicate Subtrees

- **Finished Date:** 2021-01-13
- **Tried:** 2
- **LeetCode Link:** [Find Duplicate Subtrees - LeetCode](https://leetcode.com/problems/find-duplicate-subtrees/)
- **Tags:** [`Tree`](https://leetcode.com/tag/tree/)
- **Difficulty:** Medium
- **Pass Rate:** 51.961%

[Go Version](../Go/652_Find_Duplicate_Subtrees/main.go)

## Problem description

Given the `root` of a binary tree, return all **duplicate subtrees**.

For each kind of duplicate subtrees, you only need to return the root node of any **one** of them.

Two trees are **duplicate** if they have the **same structure** with the **same node values**.

### Constraints

- The number of the nodes in the tree will be in the range `[1, 10^4]`
- `-200 <= Node.val <= 200`

## Examples

### Example 1

![](./assets/652.Find_Duplicate_Subtrees_1.jpg)

```
Input: root = [1,2,3,4,null,2,4,null,null,4]
Output: [[2,4],[4]]
```

### Example 2

![](./assets/652.Find_Duplicate_Subtrees_2.jpg)

```
Input: root = [2,1,1]
Output: [[1]]
```

### Example 3

![](./assets/652.Find_Duplicate_Subtrees_3.jpg)

```
Input: root = [2,2,2,3,null,3,null]
Output: [[2,3],[3]]
```
