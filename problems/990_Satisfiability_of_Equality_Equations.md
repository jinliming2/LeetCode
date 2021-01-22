# 990 Satisfiability of Equality Equations

- **Finished Date:** 2021-01-22
- **Tried:** 1
- **LeetCode Link:** [Satisfiability of Equality Equations - LeetCode](https://leetcode.com/problems/satisfiability-of-equality-equations/)
- **Tags:** [`Union Find`](https://leetcode.com/tag/union-find/), [`Graph`](https://leetcode.com/tag/graph/)
- **Difficulty:** Medium
- **Pass Rate:** 46.552%

[Go Version](../Go/990_Satisfiability_of_Equality_Equations/main.go)

## Problem description

Given an array equations of strings that represent relationships between variables, each string `equations[i]` has length `4` and takes one of two different forms: `"a==b"` or `"a!=b"`.  Here, `a` and `b` are lowercase letters (not necessarily different) that represent one-letter variable names.

Return `true` if and only if it is possible to assign integers to variable names so as to satisfy all the given equations.

### Constraints

- `1 <= equations.length <= 500`
- `equations[i].length == 4`
- `equations[i][0]` and `equations[i][3]` are lowercase letters
- `equations[i][1]` is either `'='` or `'!'`
- `equations[i][2]` is `'='`

## Examples

### Example 1

```
Input: ["a==b","b!=a"]
Output: false
Explanation: If we assign say, a = 1 and b = 1, then the first equation is satisfied, but not the second.  There is no way to assign the variables to satisfy both equations.
```

### Example 2

```
Input: ["b==a","a==b"]
Output: true
Explanation: We could assign a = 1 and b = 1 to satisfy both equations.
```

### Example 3

```
Input: ["a==b","b==c","a==c"]
Output: true
```

### Example 4

```
Input: ["a==b","b!=c","c==a"]
Output: false
```

### Example 5

```
Input: ["c==c","b==d","x!=z"]
Output: true
```
