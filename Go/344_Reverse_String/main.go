package main

// Runtime: 36 ms, faster than 95.48% of Go online submissions for Reverse String.
// Memory Usage: 6.3 MB, less than 87.34% of Go online submissions for Reverse String.
// https://leetcode.com/submissions/detail/449633371/

func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}
