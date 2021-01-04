# 56 Merge Intervals

- **Finished Date:** 2021-01-04
- **Tried:** 1
- **LeetCode Link:** [Merge Intervals - LeetCode](https://leetcode.com/problems/merge-intervals/)
- **Tags:** [`Array`](https://leetcode.com/tag/array/), [`Sort`](https://leetcode.com/tag/sort/)
- **Difficulty:** Medium
- **Pass Rate:** 40.52944%

[Go Version](../Go/56_Merge_Intervals/main.go)

## Problem description

Given an array of `intervals` where <code>intervals[i] = [start<sub>i</sub>, end<sub>i</sub>]</code>, merge all overlapping intervals, and return *an array of the non-overlapping intervals that cover all the intervals in the input*.

### Constraints

- 1 <= `intervals.length` <= 10<sup>4</sup>
- `intervals[i].length` == 2
- 0 <= <code>start<sub>i</sub></code> <= <code>end<sub>i</sub></code> <= 10<sup>4</sup>

## Examples

### Example 1

```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
```

### Example 2

```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```
