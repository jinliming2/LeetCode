package main

// Runtime: 16 ms, faster than 25.81% of Go online submissions for Permutation in String.
// Memory Usage: 2.8 MB, less than 47.85% of Go online submissions for Permutation in String.
func checkInclusionV1(s1 string, s2 string) bool {
	required := map[byte]int{}
	for _, c := range s1 {
		required[byte(c)]++
	}
	length := len(s1)
	left := 0
	for right, c := range s2 {
		b := byte(c)
		if _, ok := required[b]; !ok {
			for _, c := range s2[left:right] {
				required[byte(c)]++
				length++
				left++
			}
			left++
			continue
		}
		required[b]--
		length--
		if length == 0 && required[b] == 0 {
			return true
		}
		if length > 0 && required[b] >= 0 {
			continue
		}
	rLoop:
		for _, c := range s2[left : right+1] {
			b := byte(c)
			required[b]++
			length++
			left++
			if length <= 0 {
				continue
			}
			for _, r := range required {
				if r < 0 {
					continue rLoop
				}
			}
			break
		}
	}
	return false
}
