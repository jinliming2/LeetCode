package main

// Runtime: 8 ms, faster than 40.00% of Go online submissions for Remove Duplicate Letters.
// Memory Usage: 2.3 MB, less than 8.89% of Go online submissions for Remove Duplicate Letters.
// https://leetcode.com/submissions/detail/451059357/

func removeDuplicateLetters(s string) string {
	stack, index := make([]byte, 26), 0
	inStack := make(map[byte]bool, 26)
	count := make(map[byte]int, 26)
	for _, c := range []byte(s) {
		count[c]++
	}
	for _, c := range []byte(s) {
		count[c]--
		if inStack[c] {
			continue
		}
		for index > 0 {
			if stack[index-1] > c && count[stack[index-1]] > 0 {
				inStack[stack[index-1]] = false
				index--
			} else {
				break
			}
		}
		stack[index] = c
		index++
		inStack[c] = true
	}
	return string(stack[:index])
}
