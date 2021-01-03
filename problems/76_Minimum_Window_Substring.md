# 76 Minimum Window Substring

- **Finished Date:** 2021-01-02
- **Tried:** 1
- **LeetCode Link:** [Minimum Window Substring - LeetCode](https://leetcode.com/problems/minimum-window-substring/)
- **Tags:** [`Hash Table`](https://leetcode.com/tag/hash-table/), [`Two Pointers`](https://leetcode.com/tag/two-pointers/), [`String`](https://leetcode.com/tag/string/), [`Sliding Window`](https://leetcode.com/tag/sliding-window/)
- **Difficulty:** Hard
- **Pass Rate:** 35.595%

[Go Version](../Go/76_Minimum_Window_Substring/main.go)

## Problem description

Given two strings `s` and `t`, return *the minimum window in `s` which will contain all the characters in `t`*. If there is no such window in `s` that covers all characters in `t`, return *the empty string `""`*.

**Note** that If there is such a window, it is guaranteed that there will always be only one unique minimum window in `s`.

**Follow up:** Could you find an algorithm that runs in `O(n)` time?

### Constraints

- 1 <= s.length, t.length <= 10<sup>5</sup>
- `s` and `t` consist of English letters.

## Examples

### Example 1

```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
```

### Example 2

```
Input: s = "a", t = "a"
Output: "a"
```
