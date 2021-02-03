package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Smallest Subsequence of Distinct Characters.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Smallest Subsequence of Distinct Characters.
// https://leetcode.com/submissions/detail/451521732/

func smallestSubsequence(s string) string {
	stack, seen, count, index := make([]byte, 26), make([]bool, 26), make([]int, 26), 0
	for _, b := range []byte(s) {
		count[b-'a']++
	}
	for _, b := range []byte(s) {
		count[b-'a']--
		if seen[b-'a'] {
			continue
		}
		for index > 0 && stack[index-1] > b && count[stack[index-1]-'a'] > 0 {
			index--
			seen[stack[index]-'a'] = false
		}
		seen[b-'a'] = true
		stack[index] = b
		index++
	}
	return string(stack[:index])
}
