# 297 Serialize and Deserialize Binary Tree

- **Finished Date:** 2021-01-19
- **Tried:** 2
- **LeetCode Link:** [Serialize and Deserialize Binary Tree - LeetCode](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)
- **Tags:** [`Tree`](https://leetcode.com/tag/tree/), [`Design`](https://leetcode.com/tag/design/)
- **Difficulty:** Hard
- **Pass Rate:** 49.3844%

[Go Version](../Go/297_Serialize_and_Deserialize_Binary_Tree/main.go)

[Go Bug Version](../Go/297_Serialize_and_Deserialize_Binary_Tree/bug.go)

## Problem description

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

**Clarification:** The input/output format is the same as [how LeetCode serializes a binary tree]. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

[how LeetCode serializes a binary tree]: https://leetcode.com/faq/#binary-tree

### Constraints

- The number of nodes in the tree is in the range [0, 10<sup>4</sup>].
- `-1000 <= Node.val <= 1000`

## Examples

### Example 1

![](./assets/297.Serialize_and_Deserialize_Binary_Tree.jpg)

```
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
```

### Example 2

```
Input: root = []
Output: []
```

### Example 3

```
Input: root = [1]
Output: [1]
```

### Example 4

```
Input: root = [1,2]
Output: [1,2]
```
