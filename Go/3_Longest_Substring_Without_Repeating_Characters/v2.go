package main

// Runtime: 8 ms, faster than 68.49% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 2.6 MB, less than 84.25% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstringV2(s string) int {
	result := 0
	left := 0
	for right, c := range []byte(s) {
		for i := left; i < right; i++ {
			if s[i] == c {
				left = i + 1
				break
			}
		}
		length := right - left + 1
		if length > result {
			result = length
		}
	}
	return result
}
