# LeetCode [![Build And Tests](https://ci.appveyor.com/api/projects/status/0tnkehhdhxq2qlck?svg=true&retina=true)](https://ci.appveyor.com/project/LimingJin/leetcode/build/tests)
[![DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE](https://img.shields.io/badge/license-WTF%20License-blue.svg)](https://raw.githubusercontent.com/772807886/LeetCode/master/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/772807886/LeetCode.svg)](https://github.com/772807886/LeetCode/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/772807886/LeetCode.svg)](https://github.com/772807886/LeetCode/network)
[![GitHub issues](https://img.shields.io/github/issues/772807886/LeetCode.svg)](https://github.com/772807886/LeetCode/issues)

LeetCode OJ

## List

* [29 Divide Two Integers](#29-divide-two-integers)
* [65 Valid Number](#65-valid-number)
* [331 Verify Preorder Serialization of a Binary Tree](#331-verify-preorder-serialization-of-a-binary-tree)

## 29 Divide Two Integers

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

## 65 Valid Number

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

## 331 Verify Preorder Serialization of a Binary Tree

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
