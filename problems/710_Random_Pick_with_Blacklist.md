# 710 Random Pick with Blacklist

- **Finished Date:** 2021-02-01
- **Tried:** 1
- **LeetCode Link:** [Random Pick with Blacklist - LeetCode](https://leetcode.com/problems/random-pick-with-blacklist/)
- **Tags:** [`Hash Table`](https://leetcode.com/tag/hash-table/), [`Binary Search`](https://leetcode.com/tag/binary-search/), [`Sort`](https://leetcode.com/tag/sort/), [`Random`](https://leetcode.com/tag/random/)
- **Difficulty:** Hard
- **Pass Rate:** 32.689%

[Go Version](../Go/710_Random_Pick_with_Blacklist/main.go)

## Problem description

Given a blacklist `B` containing unique integers from `[0, N)`, write a function to return a uniform random integer from `[0, N)` which is **NOT** in `B`.

Optimize it such that it minimizes the call to system’s `Math.random()`.

**Explanation of Input Syntax:**

The input is two lists: the subroutines called and their arguments. `Solution`'s constructor has two arguments, `N` and the blacklist `B`. `pick` has no arguments. Arguments are always wrapped with a list, even if there aren't any.

### Constraints

- `1 <= N <= 1000000000`
- `0 <= B.length < min(100000, N)`
- `[0, N)` does NOT include N. See interval notation.

## Examples

### Example 1

```
Input:
["Solution","pick","pick","pick"]
[[1,[]],[],[],[]]
Output: [null,0,0,0]
```

### Example 2

```
Input:
["Solution","pick","pick","pick"]
[[2,[]],[],[],[]]
Output: [null,1,1,1]
```

### Example 3

```
Input:
["Solution","pick","pick","pick"]
[[3,[1]],[],[],[]]
Output: [null,0,0,2]
```

### Example 4

```
Input:
["Solution","pick","pick","pick"]
[[4,[2]],[],[],[]]
Output: [null,1,3,1]
```
