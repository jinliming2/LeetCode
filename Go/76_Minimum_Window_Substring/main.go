package main

// Runtime: 32 ms, faster than 25.19% of Go online submissions for Minimum Window Substring.
// Memory Usage: 2.9 MB, less than 45.50% of Go online submissions for Minimum Window Substring.
func minWindow(s string, t string) string {
	required := map[byte]int{}
	for _, c := range t {
		required[byte(c)]++
	}
	ms := ""
	left := 0
sLoop:
	for right, c := range s {
		b := byte(c)
		if _, ok := required[b]; !ok {
			continue
		}
		required[b]--
		for _, r := range required {
			if r > 0 {
				continue sLoop
			}
		}
		r := right + 1
		for _, c := range s[left:r] {
			l := left
			left++
			b := byte(c)
			if _, ok := required[b]; !ok {
				continue
			}
			required[b]++
			if required[b] > 0 {
				if ms == "" || right-l+1 < len(ms) {
					ms = s[l:r]
				}
				break
			}
		}
	}
	return ms
}
