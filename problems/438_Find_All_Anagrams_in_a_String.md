# 438 Find All Anagrams in a String

- **Finished Date:** 2021-01-02
- **Tried:** 1
- **LeetCode Link:** [Find All Anagrams in a String - LeetCode](https://leetcode.com/problems/find-all-anagrams-in-a-string/)
- **Tags:** [`Hash Table`](https://leetcode.com/tag/hash-table/)
- **Difficulty:** Medium
- **Pass Rate:** 44.54787%

[Go Version](../Go/438_Find_All_Anagrams_in_a_String/main.go)

## Problem description

Given a string **s** and a **non-empty** string **p**, find all the start indices of **p**'s anagrams in **s**.

Strings consists of lowercase English letters only and the length of both strings **s** and **p** will not be larger than 20,100.

The order of output does not matter.

## Examples

### Example 1

```
Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
```

### Example 2

```
Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
```
