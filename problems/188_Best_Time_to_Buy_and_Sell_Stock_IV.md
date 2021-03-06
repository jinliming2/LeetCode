# 188 Best Time to Buy and Sell Stock IV

- **Finished Date:** 2021-01-03
- **Tried:** 2
- **LeetCode Link:** [Best Time to Buy and Sell Stock IV - LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/)
- **Tags:** [`Dynamic Programming`](https://leetcode.com/tag/dynamic-programming/)
- **Difficulty:** Hard
- **Pass Rate:** 29.1558%

[Go Version](../Go/188_Best_Time_to_Buy_and_Sell_Stock_IV/main.go)

## Problem description

You are given an integer array `prices` where `prices[i]` is the price of a given stock on the ***i<sup>th</sup>*** day.

Design an algorithm to find the maximum profit. You may complete at most `k` transactions.

**Notice** that you may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

### Constraints

- `0 <= k <= 100`
- `0 <= prices.length <= 1000`
* `0 <= prices[i] <= 1000`

## Examples

### Example 1

```
Input: k = 2, prices = [2,4,1]
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
```

### Example 2

```
Input: k = 2, prices = [3,2,6,5,0,3]
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4. Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
```
