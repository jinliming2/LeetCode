# 503 Next Greater Element II

- **Finished Date:** 2021-01-26
- **Tried:** 1
- **LeetCode Link:** [Next Greater Element II - LeetCode](https://leetcode.com/problems/next-greater-element-ii/)
- **Tags:** [`Stack`](https://leetcode.com/tag/stack/)
- **Difficulty:** Medium
- **Pass Rate:** 58.217%

[Go Version](../Go/503_Next_Greater_Element_II/main.go)

## Problem description

Given a circular array (the next element of the last element is the first element of the array), print the Next Greater Number for every element. The Next Greater Number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, output -1 for this number.

### Constraints

Note: The length of given array won't exceed 10000.

## Examples

### Example 1

```
Input: [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number;
The second 1's next greater number needs to search circularly, which is also 2.
```
