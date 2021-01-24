# 295 Find Median from Data Stream

- **Finished Date:** 2021-01-24
- **Tried:** 1
- **LeetCode Link:** [Find Median from Data Stream - LeetCode](https://leetcode.com/problems/find-median-from-data-stream/)
- **Tags:** [`Heap`](https://leetcode.com/tag/heap/), [`Design`](https://leetcode.com/tag/design/)
- **Difficulty:** Hard
- **Pass Rate:** 46.5069%

[Go Version](../Go/295_Find_Median_from_Data_Stream/main.go)

## Problem description

Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

For example,

`[2,3,4]`, the median is `3`

`[2,3]`, the median is `(2 + 3) / 2 = 2.5`

Design a data structure that supports the following two operations:

- void addNum(int num) - Add a integer number from the data stream to the data structure.
- double findMedian() - Return the median of all elements so far.

**Follow up:**

- If all integer numbers from the stream are between 0 and 100, how would you optimize it?
- If 99% of all integer numbers from the stream are between 0 and 100, how would you optimize it?

## Examples

```
addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
```
