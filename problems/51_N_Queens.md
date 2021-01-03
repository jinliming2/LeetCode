# 51 N-Queens

- **Finished Date:** 2020-12-29
- **Tried:** 3
- **LeetCode Link:** [N-Queens - LeetCode](https://leetcode.com/problems/n-queens/)
- **Tags:** [`Backtracking`](https://leetcode.com/tag/backtracking/)
- **Difficulty:** Hard
- **Pass Rate:** 48.738965%

[Go Version](../Go/51_N_Queens/main.go)

## Problem description

The **n-queens** puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other.

Given an integer `n`, return *all distinct solutions to the **n-queens puzzle***.

Each solution contains a distinct board configuration of the n-queens' placement, where `'Q'` and `'.'` both indicate a queen and an empty space, respectively.

### Constraints

- 1 <= n <= 9

## Examples

### Example 1

![4-Queens](./assets/51.N-Queens.jpg)

```
Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
```

### Example 2

```
Input: n = 1
Output: [["Q"]]
```
