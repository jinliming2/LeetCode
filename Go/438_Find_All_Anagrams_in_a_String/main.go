package main

// Runtime: 4 ms, faster than 97.48% of Go online submissions for Find All Anagrams in a String.
// Memory Usage: 5.2 MB, less than 38.36% of Go online submissions for Find All Anagrams in a String.
func findAnagrams(s string, p string) []int {
	required := [26]int{} // Strings consists of lowercase English letters only
	for _, c := range p {
		required[c-'a']++
	}
	length := len(p)
	left := 0
	results := []int{}
	for _, c := range s {
		b := c - 'a'
		required[b]--
		length--
		if required[b] < 0 {
			for {
				l := left
				left++
				required[s[l]-'a']++
				length++
				if s[l] == byte(c) {
					break
				}
			}
			continue
		}
		if length == 0 {
			results = append(results, left)
			required[s[left]-'a']++
			length++
			left++
		}
	}
	return results
}
