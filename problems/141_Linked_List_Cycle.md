# 141 Linked List Cycle

- **Finished Date:** 2021-01-30
- **Tried:** 3
- **LeetCode Link:** [Linked List Cycle - LeetCode](https://leetcode.com/problems/linked-list-cycle/)
- **Tags:** [`Linked List`](https://leetcode.com/tag/linked-list/), [`Two Pointers`](https://leetcode.com/tag/two-pointers/)
- **Difficulty:** Easy
- **Pass Rate:** 42.3333%

[Go Version](../Go/141_Linked_List_Cycle/main.go)

## Problem description

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.

Return `true` if there is a cycle in the linked list. Otherwise, return `false`.

**Follow up:** Can you solve it using `O(1)` (i.e. constant) memory?

### Constraints

- The number of the nodes in the list is in the range <code>[0, 10<sup>4</sup>]</code>.
- <code>-10<sup>5</sup> <= Node.val <= 10<sup>5</sup></code>
- `pos` is `-1` or a **valid index** in the linked-list.

## Examples

### Example 1

![](./assets/141.Linked_List_Cycle_1.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
```

### Example 2

![](./assets/141.Linked_List_Cycle_2.png)

```
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
```

### Example 3

![](./assets/141.Linked_List_Cycle_3.png)

```
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
```
