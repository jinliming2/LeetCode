# 341 Flatten Nested List Iterator

- **Finished Date:** 2021-01-19
- **Tried:** 2
- **LeetCode Link:** [Flatten Nested List Iterator - LeetCode](https://leetcode.com/problems/flatten-nested-list-iterator/)
- **Tags:** [`Stack`](https://leetcode.com/tag/stack/), [`Design`](https://leetcode.com/tag/design/)
- **Difficulty:** Medium
- **Pass Rate:** 54.2455%

[Go Version](../Go/341_Flatten_Nested_List_Iterator/main.go)

## Problem description

Given a nested list of integers, implement an iterator to flatten it.

Each element is either an integer, or a list -- whose elements may also be integers or other lists.

## Examples

### Example 1

```
Input: [[1,1],2,[1,1]]
Output: [1,1,2,1,1]
Explanation: By calling next repeatedly until hasNext returns false,
             the order of elements returned by next should be: [1,1,2,1,1].
```

### Example 2

```
Input: [1,[4,[6]]]
Output: [1,4,6]
Explanation: By calling next repeatedly until hasNext returns false,
             the order of elements returned by next should be: [1,4,6].
```
