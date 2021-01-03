# 122 Best Time to Buy and Sell Stock III

- **Finished Date:** 2021-01-03
- **Tried:** 2
- **LeetCode Link:** [Best Time to Buy and Sell Stock II - LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)
- **Tags:** [`Array`](https://leetcode.com/tag/array/), [`Dynamic Programming`](https://leetcode.com/tag/dynamic-programming/)
- **Difficulty:** Hard
- **Pass Rate:** 39.575%

[Go Version](../Go/123_Best_Time_to_Buy_and_Sell_Stock_III/main.go)

## Problem description

Say you have an array for which the *i<sup>th</sup>* element is the price of a given stock on day *i*.

Design an algorithm to find the maximum profit. You may complete at most *two* transactions.

**Note:** You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

### Constraints

- 1 <= `prices.length` <= 10<sup>5</sup>
- 0 <= `prices[i]` <= 10<sup>5</sup>

## Examples

### Example 1

```
Input: prices = [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
```

### Example 2

```
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.
```

### Example 3

```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```

### Example 4

```
Input: prices = [1]
Output: 0
```
