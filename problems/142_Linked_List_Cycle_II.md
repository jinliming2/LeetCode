# 142 Linked List Cycle II

- **Finished Date:** 2021-01-30
- **Tried:** 1
- **LeetCode Link:** [Linked List Cycle II - LeetCode](https://leetcode.com/problems/linked-list-cycle-ii/)
- **Tags:** [`Linked List`](https://leetcode.com/tag/linked-list/), [`Two Pointers`](https://leetcode.com/tag/two-pointers/)
- **Difficulty:** Medium
- **Pass Rate:** 39.4419%

[Go Version](../Go/142_Linked_List_Cycle_II/main.go)

## Problem description

Given a linked list, return the node where the cycle begins. If there is no cycle, return `null`.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.

**Notice** that you **should not modify** the linked list.

**Follow up:** Can you solve it using `O(1)` (i.e. constant) memory?

### Constraints

- The number of the nodes in the list is in the range <code>[0, 10<sup>4</sup>]</code>.
- <code>-10<sup>5</sup> <= Node.val <= 10<sup>5</sup></code>
- `pos` is `-1` or a **valid index** in the linked-list.

## Examples

### Example 1

![](./assets/142.Linked_List_Cycle_II_1.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

### Example 2

![](./assets/142.Linked_List_Cycle_II_2.png)

```
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

### Example 3

![](./assets/142.Linked_List_Cycle_II_3.png)

```
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
```
