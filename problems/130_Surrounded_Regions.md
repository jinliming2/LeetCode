# 130 Surrounded Regions

- **Finished Date:** 2021-01-21
- **Tried:** 3
- **LeetCode Link:** [Surrounded Regions - LeetCode](https://leetcode.com/problems/surrounded-regions/)
- **Tags:** [`Depth-first Search`](https://leetcode.com/tag/depth-first-search/), [`Breadth-first Search`](https://leetcode.com/tag/breadth-first-search/), [`Union Find`](https://leetcode.com/tag/union-find/)
- **Difficulty:** Medium
- **Pass Rate:** 29.185%

[Go Version](../Go/130_Surrounded_Regions/main.go)

## Problem description

Given a 2D board containing `'X'` and `'O'` (**the letter O**), capture all regions surrounded by `'X'`.

A region is captured by flipping all `'O'`s into `'X'`s in that surrounded region.

### Explanation

Surrounded regions shouldn’t be on the border, which means that any `'O'` on the border of the board are not flipped to `'X'`. Any `'O'` that is not on the border and it is not connected to an `'O'` on the border will be flipped to `'X'`. Two cells are connected if they are adjacent cells connected horizontally or vertically.

## Examples

```
X X X X
X O O X
X X O X
X O X X
```

After running your function, the board should be:

```
X X X X
X X X X
X X X X
X O X X
```
