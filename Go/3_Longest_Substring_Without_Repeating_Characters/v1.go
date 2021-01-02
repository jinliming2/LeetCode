package main

// Runtime: 16 ms, faster than 31.44% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 3.3 MB, less than 37.17% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstringV1(s string) int {
	t := map[byte]int{}
	result := 0
	length := 0
	left := 0
	for _, c := range s {
		b := byte(c)
		t[b]++
		length++
		if t[b] > 1 {
			if result < length-1 {
				result = length - 1
			}
			for {
				l := left
				left++
				t[s[l]]--
				length--
				if s[l] == b {
					break
				}
			}
			continue
		}
	}
	if result < length {
		result = length
	}
	return result
}
