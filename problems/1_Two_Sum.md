# 1 Two Sum

- **Finished Date:** 2021-01-05
- **Tried:** 2
- **LeetCode Link:** [Two Sum - LeetCode](https://leetcode.com/problems/two-sum/)
- **Tags:** [`Array`](https://leetcode.com/tag/array/), [`Hash Table`](https://leetcode.com/tag/hash-table/)
- **Difficulty:** Easy
- **Pass Rate:** 46.12127%

[Go Version](../Go/1_Two_Sum/main.go)

## Problem description

Given an array of integers `nums` and an integer `target`, return *indices of the two numbers such that they add up to `target`*.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in any order.

### Constraints

- 2 <= `nums.length` <= 10<sup>3</sup>
- -10<sup>9</sup> <= `nums[i]` <= 10<sup>9</sup>
- -10<sup>9</sup> <= `target` <= 10<sup>9</sup>
- **Only one valid answer exists**.

## Examples

### Example 1

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

### Example 2

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

### Example 3

```
Input: nums = [3,3], target = 6
Output: [0,1]
```
