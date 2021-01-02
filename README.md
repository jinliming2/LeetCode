# LeetCode [![Build And Tests][AppVeyorBadge]][AppVeyorProject]
[![DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE][IconLicense]][LinkLicense]
[![GitHub stars][IconStars]][LinkStars]
[![GitHub forks][IconForks]][LinkForks]
[![GitHub issues][IconIssues]][LinkIssues]

[AppVeyorBadge]:https://ci.appveyor.com/api/projects/status/0tnkehhdhxq2qlck?svg=true&retina=true
[AppVeyorProject]:https://ci.appveyor.com/project/LimingJin/leetcode/build/tests
[IconLicense]:https://img.shields.io/badge/license-WTF%20License-blue.svg
[IconStars]:https://img.shields.io/github/stars/jinliming2/LeetCode.svg
[IconForks]:https://img.shields.io/github/forks/jinliming2/LeetCode.svg
[IconIssues]:https://img.shields.io/github/issues/jinliming2/LeetCode.svg
[LinkLicense]:./LICENSE
[LinkStars]:https://github.com/jinliming2/LeetCode/stargazers
[LinkForks]:https://github.com/jinliming2/LeetCode/network
[LinkIssues]:https://github.com/jinliming2/LeetCode/issues

LeetCode OJ

[BackToList]:#list
## List

* [3 Longest Substring Without Repeating Characters](#3-longest-substring-without-repeating-characters-back-to-list)
* [29 Divide Two Integers](#29-divide-two-integers-back-to-list)
* [34 Find First and Last Position of Element in Sorted Array](#34-find-first-and-last-position-of-element-in-sorted-array-back-to-list)
* [46 Permutations](#46-permutations-back-to-list)
* [51 N-Queens](#51-n-queens-back-to-list)
* [65 Valid Number](#65-valid-number-back-to-list)
* [76 Minimum Window Substring](#76-minimum-window-substring-back-to-list)
* [111 Minimum Depth of Binary Tree](#111-minimum-depth-of-binary-tree-back-to-list)
* [144 Binary Tree Preorder Traversal](#144-binary-tree-preorder-traversal-back-to-list)
* [236 Lowest Common Ancestor of a Binary Tree](#236-lowest-common-ancestor-of-a-binary-tree-back-to-list)
* [322 Coin Change](#322-coin-change-back-to-list)
* [331 Verify Preorder Serialization of a Binary Tree](#331-verify-preorder-serialization-of-a-binary-tree-back-to-list)
* [438 Find All Anagrams in a String](#438-find-all-anagrams-in-a-string-back-to-list)
* [509 Fibonacci Number](#509-fibonacci-number-back-to-list)
* [567 Permutation in String](#567-permutation-in-string-back-to-list)
* [704 Binary Search](#704-binary-search-back-to-list)
* [752 Open the Lock](#752-open-the-lock-back-to-list)

## 3 Longest Substring Without Repeating Characters *[Back To List][BackToList]*

> **Finished Date:** 2016-09-06
>
> **Tried:** 1
>
> **LeetCode Link:** [Longest Substring Without Repeating Characters | LeetCode OJ](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
>
> **Problem description:**
>> Given a string, find the length of the **longest substring** without repeating characters.
>>
>> Examples:
>>> Given `"abcabcbb"`, the answer is `"abc"`, which the length is 3.
>>>
>>> Given `"bbbbb"`, the answer is `"b"`, with the length of 1.
>>>
>>> Given `"pwwkew"`, the answer is `"wke"`, with the length of 3. Note that the answer must be a **substring**, `"pwke"` is a subsequence and not a substring.
>
> **Tags:** `Hash Table`, `Two Pointers`, `String`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 23.0575%

* [C# Version](./C#/CSharp/003_LongestSubstringWithoutRepeatingCharacters.cs)
* [Go Version v1](./Go/3_Longest_Substring_Without_Repeating_Characters/v1.go)
* [Go Version v2](./Go/3_Longest_Substring_Without_Repeating_Characters/v2.go)

## 29 Divide Two Integers *[Back To List][BackToList]*

> **Finished Date:** 2016-09-03
>
> **Tried:** 5
>
> **LeetCode Link:** [Divide Two Integers | LeetCode OJ](https://leetcode.com/problems/divide-two-integers/)
>
> **Problem description:**
>> Divide two integers without using multiplication, division and mod operator.
>>
>> If it is overflow, return MAX_INT.
>
> **Tags:** `Math`, `Binary Search`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 15.927%

* [C# Version](./C#/CSharp/029_DivideTwoIntegers.cs)

## 34 Find First and Last Position of Element in Sorted Array *[Back To List][BackToList]*

> **Finished Date:** 2021-01-02
>
> **Tried:** 2
>
> **LeetCode Link:** [Find First and Last Position of Element in Sorted Array - LeetCode](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)
>
> **Problem description:**
>> Given an array of integers `nums` sorted in ascending order, find the starting and ending position of a given `target` value.
>>
>> If `target` is not found in the array, return `[-1, -1]`.
>>
>> **Follow up:** Could you write an algorithm with `O(log n)` runtime complexity?
>>
>> Example 1:
>> ```
>> Input: nums = [5,7,7,8,8,10], target = 8
>> Output: [3,4]
>> ```
>> Example 2:
>> ```
>> Input: nums = [5,7,7,8,8,10], target = 6
>> Output: [-1,-1]
>> ```
>> Example 3:
>> ```
>> Input: nums = [], target = 0
>> Output: [-1,-1]
>> ```
>> Constraints:
>> - 0 <= nums.length <= 10<sup>5</sup>
>> - -10<sup>9</sup> <= nums[i] <= 10<sup>9</sup>
>> - `nums` is a non-decreasing array.
>> - -10<sup>9</sup> <= target <= 10<sup>9</sup>
>
> **Tags:** `Array`, `Binary Search`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 36.994%

* [Go Version](./Go/34_Find_First_and_Last_Position_of_Element_in_Sorted_Array/main.go)

## 46 Permutations *[Back To List][BackToList]*

> **Finished Date:** 2020-12-26
>
> **Tried:** 1
>
> **LeetCode Link:** [Permutations - LeetCode](https://leetcode.com/problems/permutations/)
>
> **Problem description:**
>> Given an array `nums` of distinct integers, return *all the possible permutations*. You can return the answer in **any order**.
>>
>> Example 1:
>> ```
>> Input: nums = [1,2,3]
>> Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
>> ```
>> Example 2:
>> ```
>> Input: nums = [0,1]
>> Output: [[0,1],[1,0]]
>> ```
>> Example 3:
>> ```
>> Input: nums = [1]
>> Output: [[1]]
>> ```
>> Constraints:
>> - 1 <= nums.length <= 6
>> - -10 <= nums[i] <= 10
>> - All the integers of `nums` are unique.
>
> **Tags:** `Backtracking`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 65.6961%

* [Go Version](./Go/46_Permutations/main.go)

## 51 N-Queens *[Back To List][BackToList]*

> **Finished Date:** 2020-12-29
>
> **Tried:** 3
>
> **LeetCode Link:** [N-Queens - LeetCode](https://leetcode.com/problems/n-queens/)
>
> **Problem description:**
>> The **n-queens** puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other.
>>
>> Given an integer `n`, return *all distinct solutions to the **n-queens puzzle***.
>>
>> Each solution contains a distinct board configuration of the n-queens' placement, where `'Q'` and `'.'` both indicate a queen and an empty space, respectively.
>>
>> Example 1:
>>
>> ![4-Queens](./assets/51.N-Queens.jpg)
>> ```
>> Input: n = 4
>> Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
>> Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
>> ```
>> Example 2:
>> ```
>> Input: n = 1
>> Output: [["Q"]]
>> ```
>> Constraints:
>> - 1 <= n <= 9
>
> **Tags:** `Backtracking`
>
> **Difficulty:** Hard
>
> **Pass Rate:** 48.738965%

* [Go Version](./Go/51_N_Queens/main.go)

## 65 Valid Number *[Back To List][BackToList]*

> **Finished Date:** 2016-09-03
>
> **Tried:** 14
>
> **LeetCode Link:** [Valid Number | LeetCode OJ](https://leetcode.com/problems/valid-number/)
>
> **Problem description:**
>> Validate if a given string is numeric.
>>
>> Some examples:
>>> `"0"` => `true`
>>>
>>> `" 0.1 "` => `true`
>>>
>>> `"abc"` => `false`
>>>
>>> `"1 a"` => `false`
>>>
>>> `"2e10"` => `true`
>>
>> **Note:** It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one.
>
> **Tags:** `Math`, `String`
>
> **Difficulty:** Hard
>
> **Pass Rate:** 12.4065%

* [C# Version](./C#/CSharp/065_ValidNumber.cs)

## 76 Minimum Window Substring *[Back To List][BackToList]*

> **Finished Date:** 2021-01-02
>
> **Tried:** 1
>
> **LeetCode Link:** [Minimum Window Substring - LeetCode](https://leetcode.com/problems/minimum-window-substring/)
>
> **Problem description:**
>> Given two strings `s` and `t`, return *the minimum window in `s` which will contain all the characters in `t`*. If there is no such window in `s` that covers all characters in `t`, return *the empty string `""`*.
>>
>> **Note** that If there is such a window, it is guaranteed that there will always be only one unique minimum window in `s`.
>>
>> Example 1:
>> ```
>> Input: s = "ADOBECODEBANC", t = "ABC"
>> Output: "BANC"
>> ```
>> Example 2:
>> ```
>> Input: s = "a", t = "a"
>> Output: "a"
>> ```
>> Constraints:
>> - 1 <= s.length, t.length <= 10<sup>5</sup>
>> - `s` and `t` consist of English letters.
>>
>> **Follow up:** Could you find an algorithm that runs in `O(n)` time?
>
> **Tags:** `Hash Table`, `Two Pointers`, `String`, `Sliding Window`
>
> **Difficulty:** Hard
>
> **Pass Rate:** 35.595%

* [Go Version](./Go/76_Minimum_Window_Substring/main.go)

## 111 Minimum Depth of Binary Tree *[Back To List][BackToList]*

> **Finished Date:** 2020-12-30
>
> **Tried:** 2(BFS) + 3(DFS)
>
> **LeetCode Link:** [Minimum Depth of Binary Tree - LeetCode](https://leetcode.com/problems/minimum-depth-of-binary-tree/)
>
> **Problem description:**
>> Given a binary tree, find its minimum depth.
>>
>> The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
>>
>> **Note:** A leaf is a node with no children.
>>
>> Example 1:
>> ```
>>    ____3___
>>   /        \
>>  9       __20__
>>         /      \
>>        15       7
>>
>> Input: root = [3,9,20,null,null,15,7]
>> Output: 2
>> ```
>> Example 2:
>> ```
>> Input: root = [2,null,3,null,4,null,5,null,6]
>> Output: 5
>> ```
>> Constraints:
>> - The number of nodes in the tree is in the range [0, 10<sup>5</sup>].
>> - -1000 <= Node.val <= 1000
>
> **Tags:** `Tree`, `Depth-first Search`, `Breadth-first Search`
>
> **Difficulty:** Easy
>
> **Pass Rate:** 39.07%

```Go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
```
* [Go BFS Version](./Go/111_Minimum_Depth_of_Binary_Tree/bfs.go)
* [Go DFS Version](./Go/111_Minimum_Depth_of_Binary_Tree/dfs.go)

## 144 Binary Tree Preorder Traversal *[Back To List][BackToList]*

> **Finished Date:** 2016-09-05
>
> **Tried:** 1
>
> **LeetCode Link:** [Binary Tree Preorder Traversal | LeetCode OJ](https://leetcode.com/problems/binary-tree-preorder-traversal/)
>
> **Problem description:**
>> Given a binary tree, return the preorder traversal of its nodes' values.
>>
>> For example:
>>> Given binary tree `{1,#,2,3}`,
>>> ```
>>>    1
>>>     \
>>>      2
>>>     /
>>>    3
>>> ```
>>
>> return `[1,2,3]`.
>>
>> **Note:** Recursive solution is trivial, could you do it iteratively?
>
> **Tags:** `Tree`, `Stack`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 41.58%

```C#
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int x) { val = x; }
 * }
 */
```
* [C# Version](./C#/CSharp/144_BinaryTreePreorderTraversal.cs)

## 236 Lowest Common Ancestor of a Binary Tree *[Back To List][BackToList]*

> **Finished Date:** 2016-09-07
>
> **Tried:** 5
>
> **LeetCode Link:** [Lowest Common Ancestor of a Binary Tree | LeetCode OJ](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)
>
> **Problem description:**
>> Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
>>
>> According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes v and w as the lowest node in T that has both v and w as descendants (where we allow **a node to be a descendant of itself**).”
>> ```
>>         _______3______
>>        /              \
>>     ___5__          ___1__
>>    /      \        /      \
>>    6      _2       0       8
>>          /  \
>>          7   4
>> ```
>> For example, the lowest common ancestor (LCA) of nodes `5` and `1` is `3`. Another example is LCA of nodes `5` and `4` is `5`, since a node can be a descendant of itself according to the LCA definition.
>
> **Tags:** `Tree`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 29.09%

```C#
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int x) { val = x; }
 * }
 */
```
* [C# Version](./C#/CSharp/236_LowestCommonAncestorOfABinaryTree.cs)

## 322 Coin Change *[Back To List][BackToList]*

> **Finished Date:** 2020-12-26
>
> **Tried:** 3
>
> **LeetCode Link:** [Coin Change - LeetCode](https://leetcode.com/problems/coin-change/)
>
> **Problem description:**
>> You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.
>>
>> You may assume that you have an infinite number of each kind of coin.
>>
>> Example 1:
>> ```
>> Input: coins = [1,2,5], amount = 11
>> Output: 3
>> Explanation: 11 = 5 + 5 + 1
>> ```
>> Example 2:
>> ```
>> Input: coins = [2], amount = 3
>> Output: -1
>> ```
>> Example 3:
>> ```
>> Input: coins = [1], amount = 0
>> Output: 0
>> ```
>> Example 4:
>> ```
>> Input: coins = [1], amount = 1
>> Output: 1
>> ```
>> Example 5:
>> ```
>> Input: coins = [1], amount = 2
>> Output: 2
>> ```
>> Constraints:
>> - 1 <= coins.length <= 12
>> - 1 <= coins[i] <= 2<sup>31</sup> - 1
>> - 0 <= amount <= 10<sup>4</sup>
>
> **Tags:** `Dynamic Programming`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 36.72237%

* [Go Version 0](./Go/322_Coin_Change/version0.go)
* [Go Version](./Go/322_Coin_Change/main.go)

## 331 Verify Preorder Serialization of a Binary Tree *[Back To List][BackToList]*

> **Finished Date:** 2016-09-05
>
> **Tried:** 3
>
> **LeetCode Link:** [Verify Preorder Serialization of a Binary Tree | LeetCode OJ](https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/)
>
> **Problem description:**
>> One way to serialize a binary tree is to use pre-order traversal. When we encounter a non-null node, we record the node's value. If it is a null node, we record using a sentinel value such as `#`.
>> ```
>>      _9_
>>     /   \
>>    3     2
>>   / \   / \
>>  4   1  #  6
>> / \ / \   / \
>> # # # #   # #
>> ```
>> For example, the above binary tree can be serialized to the string `"9,3,4,#,#,1,#,#,2,#,6,#,#"`, where `#` represents a null node.
>>
>> Given a string of comma separated values, verify whether it is a correct preorder traversal serialization of a binary tree. Find an algorithm without reconstructing the tree.
>>
>> Each comma separated value in the string must be either an integer or a character `'#'` representing `null` pointer.
>>
>> You may assume that the input format is always valid, for example it could never contain two consecutive commas such as `"1,,3"`.
>>
>> Example 1:
>>> `"9,3,4,#,#,1,#,#,2,#,6,#,#"`
>>>
>>> Return `true`
>>
>> Example 2:
>>> `"1,#"`
>>>
>>> Return `false`
>>
>> Example 3:
>>> `"9,#,#,1"`
>>>
>>> Return `false`
>
> **Tags:** `Stack`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 33.52669%

* [C# Version](./C#/CSharp/331_VerifyPreorderSerializationOfABinaryTree.cs)

## 438 Find All Anagrams in a String *[Back To List][BackToList]*

> **Finished Date:** 2021-01-02
>
> **Tried:** 1
>
> **LeetCode Link:** [Find All Anagrams in a String - LeetCode](https://leetcode.com/problems/find-all-anagrams-in-a-string/)
>
> **Problem description:**
>> Given a string **s** and a **non-empty** string **p**, find all the start indices of **p**'s anagrams in **s**.
>>
>> Strings consists of lowercase English letters only and the length of both strings **s** and **p** will not be larger than 20,100.
>>
>> The order of output does not matter.
>>
>> Example 1:
>> ```
>> Input:
>> s: "cbaebabacd" p: "abc"
>>
>> Output:
>> [0, 6]
>>
>> Explanation:
>> The substring with start index = 0 is "cba", which is an anagram of "abc".
>> The substring with start index = 6 is "bac", which is an anagram of "abc".
>> ```
>> Example 2:
>> ```
>> Input:
>> s: "abab" p: "ab"
>>
>> Output:
>> [0, 1, 2]
>>
>> Explanation:
>> The substring with start index = 0 is "ab", which is an anagram of "ab".
>> The substring with start index = 1 is "ba", which is an anagram of "ab".
>> The substring with start index = 2 is "ab", which is an anagram of "ab".
>> ```
>
> **Tags:** `Hash Table`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 44.54787%

* [Go Version](./Go/438_Find_All_Anagrams_in_a_String/main.go)

## 509 Fibonacci Number *[Back To List][BackToList]*

> **Finished Date:** 2020-12-24
>
> **Tried:** 1
>
> **LeetCode Link:** [Fibonacci Number - LeetCode](https://leetcode.com/problems/fibonacci-number/submissions/)
>
> **Problem description:**
>> The **Fibonacci numbers**, commonly denoted `F(n)` form a sequence, called the **Fibonacci sequence**, such that each number is the sum of the two preceding ones, starting from `0` and `1`. That is,
>> ```
>> F(0) = 0, F(1) = 1
>> F(n) = F(n - 1) + F(n - 2), for n > 1.
>> ```
>> Given `n`, calculate `F(n)`.
>>
>> Example 1:
>> ```
>> Input: n = 2
>> Output: 1
>> Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
>> ```
>> Example 2:
>> ```
>> Input: n = 3
>> Output: 2
>> Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
>> ```
>> Example 3:
>> ```
>> Input: n = 4
>> Output: 3
>> Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
>> ```
>> Constraints:
>> - `0 <= n <= 30`
>
> **Tags:** `Array`
>
> **Difficulty:** Easy
>
> **Pass Rate:** 67.17115%

* [Go Version](./Go/509_Fibonacci_Number/main.go)

## 567 Permutation in String *[Back To List][BackToList]*

> **Finished Date:** 2021-01-02
>
> **Tried:** 2
>
> **LeetCode Link:** [Permutation in String - LeetCode](https://leetcode.com/problems/permutation-in-string/)
>
> **Problem description:**
>> Given two strings **s1** and **s2**, write a function to return true if **s2** contains the permutation of **s1**. In other words, one of the first string's permutations is the **substring** of the second string.
>>
>> Example 1:
>> ```
>> Input: s1 = "ab" s2 = "eidbaooo"
>> Output: True
>> Explanation: s2 contains one permutation of s1 ("ba").
>> ```
>> Example 2:
>> ```
>> Input:s1= "ab" s2 = "eidboaoo"
>> Output: False
>> ```
>> Constraints:
>> - The input strings only contain lower case letters.
>> - The length of both given strings is in range [1, 10,000].
>
> **Tags:** `Two Pointers`, `Sliding Window`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 44.5186%

* [Go Version 1](./Go/567_Permutation_in_String/v1.go)
* [Go Version 2](./Go/567_Permutation_in_String/v2.go)

## 704 Binary Search *[Back To List][BackToList]*

> **Finished Date:** 2021-01-02
>
> **Tried:** 1
>
> **LeetCode Link:** [Binary Search - LeetCode](https://leetcode.com/problems/binary-search/)
>
> **Problem description:**
>> Given a **sorted** (in ascending order) integer array `nums` of `n` elements and a `target` value, write a function to search `target` in `nums`. If `target` exists, then return its index, otherwise return `-1`.
>>
>> Example 1:
>> ```
>> Input: nums = [-1,0,3,5,9,12], target = 9
>> Output: 4
>> Explanation: 9 exists in nums and its index is 4
>> ```
>> Example 2:
>> ```
>> Input: nums = [-1,0,3,5,9,12], target = 2
>> Output: -1
>> Explanation: 2 does not exist in nums so return -1
>> ```
>> Note:
>> - You may assume that all elements in `nums` are unique.
>> - `n` will be in the range `[1, 10000]`.
>> - The value of each element in `nums` will be in the range `[-9999, 9999]`.
>
> **Tags:** `Binary Search`
>
> **Difficulty:** Easy
>
> **Pass Rate:** 53.9026%

* [Go Version](./Go/704_Binary_Search/main.go)

## 752 Open the Lock *[Back To List][BackToList]*

> **Finished Date:** 2020-12-30
>
> **Tried:** 3
>
> **LeetCode Link:** [Open the Lock - LeetCode](https://leetcode.com/problems/open-the-lock/)
>
> **Problem description:**
>> You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: `'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'`. The wheels can rotate freely and wrap around: for example we can turn `'9'` to be `'0'`, or `'0'` to be `'9'`. Each move consists of turning one wheel one slot.
>>
>> The lock initially starts at `'0000'`, a string representing the state of the 4 wheels.
>>
>> You are given a list of `deadends` dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.
>>
>> Given a `target` representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.
>>
>> Example 1:
>> ```
>> Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
>> Output: 6
>> Explanation:
>> A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
>> Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
>> because the wheels of the lock become stuck after the display becomes the dead end "0102".
>> ```
>> Example 2:
>> ```
>> Input: deadends = ["8888"], target = "0009"
>> Output: 1
>> Explanation:
>> We can turn the last wheel in reverse to move from "0000" -> "0009".
>> ```
>> Example 3:
>> ```
>> Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
>> Output: -1
>> Explanation:
>> We can't reach the target without getting stuck.
>> ```
>> Example 4:
>> ```
>> Input: deadends = ["0000"], target = "8888"
>> Output: -1
>> ```
>> Constraints:
>> - `1 <= deadends.length <= 500`
>> - `deadends[i].length == 4`
>> - `target.length == 4`
>> - target **will not be** in the list `deadends`.
>> - `target` and `deadends[i]` consist of digits only.
>
> **Tags:** `Breadth-first Search`
>
> **Difficulty:** Medium
>
> **Pass Rate:** 52.42%

* [Go BFS Version](./Go/752_Open_the_Lock/bfs.go)
* [Go Double BFS Version](./Go/752_Open_the_Lock/doubleBFS.go)
