# 986 Interval List Intersections

- **Finished Date:** 2021-01-05
- **Tried:** 2
- **LeetCode Link:** [Interval List Intersections - LeetCode](https://leetcode.com/problems/interval-list-intersections/)
- **Tags:** [`Two Pointers`](https://leetcode.com/tag/two-pointers/)
- **Difficulty:** Medium
- **Pass Rate:** 68.02647%

[Go Version](../Go/986_Interval_List_Intersections/main.go)

## Problem description

Given two lists of **closed** intervals, each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

*(Formally, a closed interval `[a, b]` (with `a <= b`) denotes the set of real numbers `x` with `a <= x <= b`.  The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval.  For example, the intersection of [1, 3] and [2, 4] is [2, 3].)*

### Constraints

- 0 <= `A.length` < 1000
- 0 <= `B.length` < 1000
- 0 <= `A[i].start`, `A[i].end`, `B[i].start`, `B[i].end` < 10<sup>9</sup>

## Examples

### Example 1

![](./assets/986.Interval_List_Intersections.png)

```
Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
```
