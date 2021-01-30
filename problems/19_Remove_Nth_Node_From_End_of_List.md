# 19 Remove Nth Node From End of List

- **Finished Date:** 2021-01-30
- **Tried:** 1
- **LeetCode Link:** [Remove Nth Node From End of List - LeetCode](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)
- **Tags:** [`Linked List`](https://leetcode.com/tag/linked-list/), [`Two Pointers`](https://leetcode.com/tag/two-pointers/)
- **Difficulty:** Medium
- **Pass Rate:** 35.6647%

[Go Version](../Go/19_Remove_Nth_Node_From_End_of_List/main.go)

## Problem description

Given the `head` of a linked list, remove the <code>n<sup>th</sup></code> node from the end of the list and return its head.

**Follow up:** Could you do this in one pass?

### Constraints

- The number of nodes in the list is `sz`.
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

## Examples

### Example 1

![](./assets/19.Remove_Nth_Node_From_End_of_List.jpg)

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

### Example 2

```
Input: head = [1], n = 1
Output: []
```

### Example 3

```
Input: head = [1,2], n = 1
Output: [1]
```
