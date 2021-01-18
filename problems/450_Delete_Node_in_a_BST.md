# 450 Delete Node in a BST

- **Finished Date:** 2021-01-18
- **Tried:** 1
- **LeetCode Link:** [Delete Node in a BST - LeetCode](https://leetcode.com/problems/delete-node-in-a-bst/)
- **Tags:** [`Tree`](https://leetcode.com/tag/tree/)
- **Difficulty:** Medium
- **Pass Rate:** 45.1494%

[Go Version](../Go/450_Delete_Node_in_a_BST/main.go)

## Problem description

Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

1. Search for a node to remove.
2. If the node is found, delete the node.

**Follow up:** Can you solve it with time complexity `O(height of tree)`?

### Constraints

- The number of nodes in the tree is in the range [0, 10<sup>4</sup>].
- -10<sup>5</sup> <= `Node.val` <= 10<sup>5</sup>
- Each node has a **unique** value.
- `root` is a valid binary search tree.
- -10<sup>5</sup> <= `key` <= 10<sup>5</sup>

## Examples

### Example 1

![](./assets/450.Delete_Node_in_a_BST_1.jpg)

```
Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.
```

![](./assets/450.Delete_Node_in_a_BST_2.jpg)

### Example 2

```
Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.
```

### Example 3

```
Input: root = [], key = 0
Output: []
```
