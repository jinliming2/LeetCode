# 701 Insert into a Binary Search Tree

- **Finished Date:** 2021-01-17
- **Tried:** 1
- **LeetCode Link:** [Insert into a Binary Search Tree - LeetCode](https://leetcode.com/problems/insert-into-a-binary-search-tree/)
- **Tags:** [`Tree`](https://leetcode.com/tag/tree/)
- **Difficulty:** Medium
- **Pass Rate:** 75.8563%

[Go Version](../Go/701_Insert_into_a_Binary_Search_Tree/main.go)

## Problem description

You are given the root `node` of a binary search tree (BST) and a `value` to insert into the tree. Return *the root node of the BST after the insertion*. It is **guaranteed** that the new value does not exist in the original BST.

**Notice** that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return **any of them**.

### Constraints

- The number of nodes in the tree will be in the range [0, 10<sup>4</sup>].
- -10<sup>8</sup> <= `Node.val` <= 10<sup>8</sup>
- All the values Node.val are unique.
- -10<sup>8</sup> <= `val` <= 10<sup>8</sup>
- It's **guaranteed** that `val` does not exist in the original BST.

## Examples

### Example 1

![](./assets/701.Insert_into_a_Binary_Search_Tree_1.jpg)

```
Input: root = [4,2,7,1,3], val = 5
Output: [4,2,7,1,3,5]
Explanation: Another accepted tree is:
```

![](./assets/701.Insert_into_a_Binary_Search_Tree_2.jpg)

### Example 2

```
Input: root = [40,20,60,10,30,50,70], val = 25
Output: [40,20,60,10,30,50,70,null,null,25]
```

### Example 3

```
Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
Output: [4,2,7,1,3,5]
```
