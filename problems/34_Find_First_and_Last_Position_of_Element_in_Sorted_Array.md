# 34 Find First and Last Position of Element in Sorted Array

- **Finished Date:** 2021-01-02
- **Tried:** 2
- **LeetCode Link:** [Find First and Last Position of Element in Sorted Array - LeetCode](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)
- **Tags:** [`Array`](https://leetcode.com/tag/array/), [`Binary Search`](https://leetcode.com/tag/binary-search/)
- **Difficulty:** Medium
- **Pass Rate:** 36.994%

[Go Version](../Go/34_Find_First_and_Last_Position_of_Element_in_Sorted_Array/main.go)

## Problem description

Given an array of integers `nums` sorted in ascending order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

**Follow up:** Could you write an algorithm with `O(log n)` runtime complexity?

### Constraints

- 0 <= `nums.length` <= 10<sup>5</sup>
- -10<sup>9</sup> <= `nums[i]` <= 10<sup>9</sup>
- `nums` is a non-decreasing array.
- -10<sup>9</sup> <= `target` <= 10<sup>9</sup>

## Examples

### Example 1

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

### Example 2

```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

### Example 3

```
Input: nums = [], target = 0
Output: [-1,-1]
```
