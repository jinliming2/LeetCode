# 25 Reverse Nodes in k-Group

- **Finished Date:** 2021-01-10
- **Tried:** 1
- **LeetCode Link:** [Reverse Nodes in k-Group - LeetCode](https://leetcode.com/problems/reverse-nodes-in-k-group/)
- **Tags:** [`Linked List`](https://leetcode.com/tag/linked-list/)
- **Difficulty:** Hard
- **Pass Rate:** 44.1548%

[Go Version](../Go/25_Reverse_Nodes_in_k_Group/main.go)

## Problem description

Given a linked list, reverse the nodes of a linked list *k* at a time and return its modified list.

*k* is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of *k* then left-out nodes, in the end, should remain as it is.

**Follow up:**

Could you solve the problem in `O(1)` extra memory space?
You may not alter the values in the list's nodes, only nodes itself may be changed.

### Constraints

- The number of nodes in the list is in the range `sz`.
- `1 <= sz <= 5000`
- `0 <= Node.val <= 1000`
- `1 <= k <= sz`

## Examples

### Example 1

![](./assets/25.Reverse_Nodes_in_k-Group_1.jpg)

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
```

### Example 2

![](./assets/25.Reverse_Nodes_in_k-Group_2.jpg)

```
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
```

### Example 3

```
Input: head = [1,2,3,4,5], k = 1
Output: [1,2,3,4,5]
```

### Example 4

```
Input: head = [1], k = 1
Output: [1]
```
