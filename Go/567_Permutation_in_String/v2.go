package main

// Runtime: 16 ms, faster than 25.81% of Go online submissions for Permutation in String.
// Memory Usage: 2.9 MB, less than 24.73% of Go online submissions for Permutation in String.
func checkInclusionV2(s1 string, s2 string) bool {
	required := make(map[byte]int, 26)
	for _, c := range s1 {
		required[byte(c)]++
	}
	length := len(s1)
	left := 0
	for _, c := range s2 {
		b := byte(c)
		required[b]--
		length--
		if required[b] < 0 {
			for {
				l := left
				left++
				required[s2[l]]++
				length++
				if s2[l] == b {
					break
				}
			}
			continue
		}
		if length == 0 {
			return true
		}
	}
	return false
}
