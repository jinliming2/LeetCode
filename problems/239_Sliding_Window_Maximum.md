# 239 Sliding Window Maximum

- **Finished Date:** 2021-01-27
- **Tried:** 3
- **LeetCode Link:** [Sliding Window Maximum - LeetCode](https://leetcode.com/problems/sliding-window-maximum/)
- **Tags:** [`Heap`](https://leetcode.com/tag/heap/), [`Sliding Window`](https://leetcode.com/tag/sliding-window/), [`Dequeue`](https://leetcode.com/tag/dequeue/)
- **Difficulty:** Hard
- **Pass Rate:** 44.5576%

[Go Version 1](../Go/239_Sliding_Window_Maximum/v1.go)

[Go Version 2](../Go/239_Sliding_Window_Maximum/v2.go)

## Problem description

You are given an array of integers `nums`, there is a sliding window of size `k` which is moving from the very left of the array to the very right. You can only see the `k` numbers in the window. Each time the sliding window moves right by one position.

Return *the max sliding window*.

### Constraints

- <code>1 <= nums.length <= 10<sup>5</sup></code>
- <code>-10<sup>4</sup> <= nums[i] <= 10<sup>4</sup></code>
- 1 <= k <= nums.length

## Examples

### Example 1

```
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
```

### Example 2

```
Input: nums = [1], k = 1
Output: [1]
```

### Example 3

```
Input: nums = [1,-1], k = 1
Output: [1,-1]
```

### Example 4

```
Input: nums = [9,11], k = 2
Output: [11]
```

### Example 5

```
Input: nums = [4,-2], k = 2
Output: [4]
```
