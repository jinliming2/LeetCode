# 167 Two Sum II - Input array is sorted

- **Finished Date:** 2021-01-30
- **Tried:** 1
- **LeetCode Link:** [Two Sum II - Input array is sorted - LeetCode](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)
- **Tags:** [`Array`](https://leetcode.com/tag/array/), [`Tow Pointers`](https://leetcode.com/tag/two-pointers/), [`Binary Search`](https://leetcode.com/tag/binary-search/)
- **Difficulty:** Easy
- **Pass Rate:** 55.4833%

[Go Version](../Go/167_Two_Sum_II_Input_array_is_sorted/main.go)

## Problem description

Given an array of integers `numbers` that is already ***sorted in ascending order***, find two numbers such that they add up to a specific `target` number.

Return the indices of the two numbers (**1-indexed**) as an integer array `answer` of size `2`, where `1 <= answer[0] < answer[1] <= numbers.length`.

You may assume that each input would have **exactly one solution** and you **may not** use the same element twice.

### Constraints

- <code>2 <= numbers.length <= 3 * 10<sup>4</sup></code>
- `-1000 <= numbers[i] <= 1000`
- `numbers` is sorted in increasing order.
- `-1000 <= target <= 1000`
- **Only one valid answer exists.**

## Examples

### Example 1

```
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
```

### Example 2

```
Input: numbers = [2,3,4], target = 6
Output: [1,3]
```

### Example 3

```
Input: numbers = [-1,0], target = -1
Output: [1,2]
```
