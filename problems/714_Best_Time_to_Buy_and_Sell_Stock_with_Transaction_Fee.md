# 714 Best Time to Buy and Sell Stock with Transaction Fee

- **Finished Date:** 2021-01-03
- **Tried:** 1
- **LeetCode Link:** [Best Time to Buy and Sell Stock with Transaction Fee - LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)
- **Tags:** [`Array`](https://leetcode.com/tag/array/), [`Dynamic Programming`](https://leetcode.com/tag/dynamic-programming/), [`Greedy`](https://leetcode.com/tag/greedy/)
- **Difficulty:** Medium
- **Pass Rate:** 55.7453%

[Go Version](../Go/714_Best_Time_to_Buy_and_Sell_Stock_with_Transaction_Fee/main.go)

## Problem description

Your are given an array of integers `prices`, for which the `i`-th element is the price of a given stock on day `i`; and a non-negative integer `fee` representing a transaction fee.

You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction. You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)

Return the maximum profit you can make.

### Constraints

- `0 < prices.length <= 50000`.
- `0 < prices[i] < 50000`.
- `0 <= fee < 50000`.

## Examples

### Example 1

```
Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
Output: 8
Explanation: The maximum profit can be achieved by:
Buying at prices[0] = 1
Selling at prices[3] = 8
Buying at prices[4] = 4
Selling at prices[5] = 9
The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
```
