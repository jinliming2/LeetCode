# 322 Coin Change

- **Finished Date:** 2020-12-26
- **Tried:** 3
- **LeetCode Link:** [Coin Change - LeetCode](https://leetcode.com/problems/coin-change/)
- **Tags:** [`Dynamic Programming`](https://leetcode.com/tag/dynamic-programming/)
- **Difficulty:** Medium
- **Pass Rate:** 36.72237%

[Go Version 0](../Go/322_Coin_Change/version0.go)

[Go Version](../Go/322_Coin_Change/main.go)

## Problem description

You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

You may assume that you have an infinite number of each kind of coin.

### Constraints

- 1 <= coins.length <= 12
- 1 <= coins[i] <= 2<sup>31</sup> - 1
- 0 <= amount <= 10<sup>4</sup>

## Examples

### Example 1

```
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
```

### Example 2

```
Input: coins = [2], amount = 3
Output: -1
```

### Example 3

```
Input: coins = [1], amount = 0
Output: 0
```

### Example 4

```
Input: coins = [1], amount = 1
Output: 1
```

### Example 5

```
Input: coins = [1], amount = 2
Output: 2
```
