# 454 4Sum II

- **Finished Date:** 2021-01-06
- **Tried:** 2
- **LeetCode Link:** [4Sum II - LeetCode](https://leetcode.com/problems/4sum-ii/)
- **Tags:** [`Hash Table`](https://leetcode.com/tag/hash-table/), [`Binary Search`](https://leetcode.com/tag/binary-search/)
- **Difficulty:** Medium
- **Pass Rate:** 54.5195%

[Go Version](../Go/454_4Sum_II/main.go)

## Problem description

Given four lists A, B, C, D of integer values, compute how many tuples `(i, j, k, l)` there are such that `A[i] + B[j] + C[k] + D[l]` is zero.

### Constraints

To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the range of -2<sup>28</sup> to 2<sup>28</sup> - 1 and the result is guaranteed to be at most 2<sup>31</sup> - 1.

## Examples

```Input:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

Output:
2

Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
```
