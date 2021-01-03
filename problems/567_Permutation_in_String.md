# 567 Permutation in String

- **Finished Date:** 2021-01-02
- **Tried:** 2
- **LeetCode Link:** [Permutation in String - LeetCode](https://leetcode.com/problems/permutation-in-string/)
- **Tags:** [`Two Pointers`](https://leetcode.com/tag/two-pointers/), [`Sliding Window`](https://leetcode.com/tag/sliding-window/)
- **Difficulty:** Medium
- **Pass Rate:** 44.5186%

[Go Version 1](../Go/567_Permutation_in_String/v1.go)

[Go Version 2](../Go/567_Permutation_in_String/v2.go)

## Problem description

Given two strings **s1** and **s2**, write a function to return true if **s2** contains the permutation of **s1**. In other words, one of the first string's permutations is the **substring** of the second string.

### Constraints

- The input strings only contain lower case letters.
- The length of both given strings is in range [1, 10,000].

## Examples

### Example 1

```
Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
```

### Example 2

```
Input:s1= "ab" s2 = "eidboaoo"
Output: False
```
