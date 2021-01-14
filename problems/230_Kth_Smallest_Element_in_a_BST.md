# 230 Kth Smallest Element in a BST

- **Finished Date:** 2021-01-14
- **Tried:** 2
- **LeetCode Link:** [Kth Smallest Element in a BST - LeetCode](https://leetcode.com/problems/kth-smallest-element-in-a-bst/)
- **Tags:** [`Binary Search`](https://leetcode.com/tag/binary-search/), [`Tree`](https://leetcode.com/tag/tree/)
- **Difficulty:** Medium
- **Pass Rate:** 62.1264%

[Go Version](../Go/230_Kth_Smallest_Element_in_a_BST/main.go)

## Problem description

Given a binary search tree, write a function `kthSmallest` to find the **k**th smallest element in it.

**Follow up:**

What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

### Constraints

- The number of elements of the BST is between `1` to `10^4`.
- You may assume `k` is always valid, `1 ≤ k ≤ BST's total elements`.

## Examples

### Example 1

```
Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
```

### Example 2

```
Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
```
