# 337 House Robber III

- **Finished Date:** 2021-01-03
- **Tried:** 3
- **LeetCode Link:** [House Robber III - LeetCode](https://leetcode.com/problems/house-robber-iii/)
- **Tags:** [`Dynamic Programming`](https://leetcode.com/tag/dynamic-programming/), [`Tree`](https://leetcode.com/tag/tree/), [`Depth-first Search`](https://leetcode.com/tag/depth-first-search/)
- **Difficulty:** Medium
- **Pass Rate:** 51.64995%

[Go Version](../Go/337_House_Robber_III/main.go)

## Problem description

The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

## Examples

### Example 1

```
Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
```

### Example 2

```
Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
```
