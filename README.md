# LeetCode [![Build And Tests](https://ci.appveyor.com/api/projects/status/0tnkehhdhxq2qlck?svg=true&retina=true)](https://ci.appveyor.com/project/LimingJin/leetcode/build/tests)
[![DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE](https://img.shields.io/badge/license-WTF%20License-blue.svg)](https://raw.githubusercontent.com/772807886/LeetCode/master/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/772807886/LeetCode.svg)](https://github.com/772807886/LeetCode/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/772807886/LeetCode.svg)](https://github.com/772807886/LeetCode/network)
[![GitHub issues](https://img.shields.io/github/issues/772807886/LeetCode.svg)](https://github.com/772807886/LeetCode/issues)

LeetCode OJ

[GoToList]:#list "A Link Go To List"
## List

* [3 Longest Substring Without Repeating Characters](#3-longest-substring-without-repeating-characters-go-to-list)
* [29 Divide Two Integers](#29-divide-two-integers-go-to-list)
* [65 Valid Number](#65-valid-number-go-to-list)
* [144 Binary Tree Preorder Traversal](#144-binary-tree-preorder-traversal-go-to-list)
* [236 Lowest Common Ancestor of a Binary Tree](#236-lowest-common-ancestor-of-a-binary-tree-go-to-list)
* [331 Verify Preorder Serialization of a Binary Tree](#331-verify-preorder-serialization-of-a-binary-tree-go-to-list)

## 3 Longest Substring Without Repeating Characters *[Go To List][GoToList]*

> **Finished Date:** 2016-09-06

> **Tried:** 1

> **LeetCode Link:** [Longest Substring Without Repeating Characters | LeetCode OJ](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

> **Problem description:**

>> Given a string, find the length of the **longest substring** without repeating characters.

>> Examples:

>>> Given `"abcabcbb"`, the answer is `"abc"`, which the length is 3.
    
>>> Given `"bbbbb"`, the answer is `"b"`, with the length of 1.
    
>>> Given `"pwwkew"`, the answer is `"wke"`, with the length of 3. Note that the answer must be a **substring**, `"pwke"` is a subsequence and not a substring.

> **Tags:** `Hash Table`, `Two Pointers`, `String`

> **Difficulty:** Medium

> **Pass Rate:** 23.0575%

* [C# Version](https://github.com/772807886/LeetCode/blob/master/CSharp/003_LongestSubstringWithoutRepeatingCharacters.cs)

## 29 Divide Two Integers *[Go To List][GoToList]*

> **Finished Date:** 2016-09-03

> **Tried:** 5

> **LeetCode Link:** [Divide Two Integers | LeetCode OJ](https://leetcode.com/problems/divide-two-integers/)

> **Problem description:**

>> Divide two integers without using multiplication, division and mod operator.

>> If it is overflow, return MAX_INT.

> **Tags:** `Math`, `Binary Search`

> **Difficulty:** Medium

> **Pass Rate:** 15.927%

* [C# Version](https://github.com/772807886/LeetCode/blob/master/CSharp/029_DivideTwoIntegers.cs)

## 65 Valid Number *[Go To List][GoToList]*

> **Finished Date:** 2016-09-03

> **Tried:** 14

> **LeetCode Link:** [Valid Number | LeetCode OJ](https://leetcode.com/problems/valid-number/)

> **Problem description:**

>> Validate if a given string is numeric.

>> Some examples:

>>> `"0"` => `true`

>>> `" 0.1 "` => `true`

>>> `"abc"` => `false`

>>> `"1 a"` => `false`

>>> `"2e10"` => `true`

>> **Note:** It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one.

> **Tags:** `Math`, `String`

> **Difficulty:** Hard

> **Pass Rate:** 12.4065%

* [C# Version](https://github.com/772807886/LeetCode/blob/master/CSharp/065_ValidNumber.cs)

## 144 Binary Tree Preorder Traversal *[Go To List][GoToList]*

> **Finished Date:** 2016-09-05

> **Tried:** 1

> **LeetCode Link:** [Binary Tree Preorder Traversal | LeetCode OJ](https://leetcode.com/problems/binary-tree-preorder-traversal/)

> **Problem description:**

>> Given a binary tree, return the preorder traversal of its nodes' values.

>> For example:

>>> Given binary tree `{1,#,2,3}`,
```
   1
    \
     2
    /
   3
```
>> return `[1,2,3]`.

>> **Note:** Recursive solution is trivial, could you do it iteratively?

> **Tags:** `Tree`, `Stack`

> **Difficulty:** Medium

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
* [C# Version](https://github.com/772807886/LeetCode/blob/master/CSharp/144_BinaryTreePreorderTraversal.cs)

## 236 Lowest Common Ancestor of a Binary Tree *[Go To List][GoToList]*

> **Finished Date:** 2016-09-07

> **Tried:** 5

> **LeetCode Link:** [Lowest Common Ancestor of a Binary Tree | LeetCode OJ](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)

> **Problem description:**

>> Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

>> According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes v and w as the lowest node in T that has both v and w as descendants (where we allow **a node to be a descendant of itself**).”
```
        _______3______
       /              \
    ___5__          ___1__
   /      \        /      \
   6      _2       0       8
         /  \
         7   4
```
>> For example, the lowest common ancestor (LCA) of nodes `5` and `1` is `3`. Another example is LCA of nodes `5` and `4` is `5`, since a node can be a descendant of itself according to the LCA definition.

> **Tags:** `Tree`

> **Difficulty:** Medium

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
* [C# Version](https://github.com/772807886/LeetCode/blob/master/CSharp/236_LowestCommonAncestorOfABinaryTree.cs)

## 331 Verify Preorder Serialization of a Binary Tree *[Go To List][GoToList]*

> **Finished Date:** 2016-09-05

> **Tried:** 3

> **LeetCode Link:** [Verify Preorder Serialization of a Binary Tree | LeetCode OJ](https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/)

> **Problem description:**

>> One way to serialize a binary tree is to use pre-order traversal. When we encounter a non-null node, we record the node's value. If it is a null node, we record using a sentinel value such as `#`.
```
     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
```
>> For example, the above binary tree can be serialized to the string `"9,3,4,#,#,1,#,#,2,#,6,#,#"`, where `#` represents a null node.

>> Given a string of comma separated values, verify whether it is a correct preorder traversal serialization of a binary tree. Find an algorithm without reconstructing the tree.

>> Each comma separated value in the string must be either an integer or a character `'#'` representing `null` pointer.

>> You may assume that the input format is always valid, for example it could never contain two consecutive commas such as `"1,,3"`.

>> Example 1:

>>> `"9,3,4,#,#,1,#,#,2,#,6,#,#"`

>>> Return `true`

>> Example 2:

>>> `"1,#"`

>>> Return `false`

>> Example 3:

>>> `"9,#,#,1"`

>>> Return `false`

> **Tags:** `Stack`

> **Difficulty:** Medium

> **Pass Rate:** 33.52669%

* [C# Version](https://github.com/772807886/LeetCode/blob/master/CSharp/331_VerifyPreorderSerializationOfABinaryTree.cs)
