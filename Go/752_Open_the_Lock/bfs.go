package main

// Runtime: 296 ms, faster than 14.43% of Go online submissions for Open the Lock.
// Memory Usage: 7.5 MB, less than 21.65% of Go online submissions for Open the Lock.
func openLockBFS(deadends []string, target string) int {
	dp := map[string]bool{}
	current := []string{"0000"}

	l0, r9 := byte('0'-1), byte('9'+1)

	steps := 0
	for len(current) > 0 {
		new := make([]string, 0, len(current)*8)
	currentLoop:
		for _, item := range current {
			for _, d := range deadends {
				if d == item {
					continue currentLoop
				}
			}
			if target == item {
				return steps
			}
			for i := 0; i < 4; i++ {
				b := []byte(item)
				b[i]--
				if b[i] == l0 {
					b[i] = '9'
				}
				s := string(b)
				if _, ok := dp[s]; !ok {
					new = append(new, s)
					dp[s] = true
				}

				b[i]++
				if b[i] == r9 {
					b[i] = '0'
				}
				b[i]++
				if b[i] == r9 {
					b[i] = '0'
				}
				s = string(b)
				if _, ok := dp[s]; !ok {
					new = append(new, s)
					dp[s] = true
				}
			}
		}
		current = new
		steps++
	}

	return -1
}
