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
* [46 Permutations](#46-permutations-back-to-list)
* [51 N-Queens](#51-n-queens-back-to-list)
* [65 Valid Number](#65-valid-number-back-to-list)
* [111 Minimum Depth of Binary Tree](#111-minimum-depth-of-binary-tree-back-to-list)
* [144 Binary Tree Preorder Traversal](#144-binary-tree-preorder-traversal-back-to-list)
* [236 Lowest Common Ancestor of a Binary Tree](#236-lowest-common-ancestor-of-a-binary-tree-back-to-list)
* [322 Coin Change](#322-coin-change-back-to-list)
* [331 Verify Preorder Serialization of a Binary Tree](#331-verify-preorder-serialization-of-a-binary-tree-back-to-list)
* [509 Fibonacci Number](#509-fibonacci-number-back-to-list)

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

## 111 Minimum Depth of Binary Tree *[Back To List][BackToList]*

> **Finished Date:** 2020-12-30
>
> **Tried:** 2
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
* [Go Version](./Go/111_Minimum_Depth_of_Binary_Tree/main.go)

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
