package main

// Runtime: 476 ms, faster than 5.05% of Go online submissions for Open the Lock.
// Memory Usage: 6.6 MB, less than 73.74% of Go online submissions for Open the Lock.
func openLockDoubleBFS(deadends []string, target string) int {
	dp := append([]string{}, deadends...)

	pc := []string{"0000"}
	pt := []string{target}

	step := 0

	for len(pc) > 0 {
		new := make([]string, 0, len(pc)*8)
	pcLoop:
		for _, item := range pc {
			for _, d := range dp {
				if d == item {
					continue pcLoop
				}
			}
			for _, t := range pt {
				if t == item {
					return step
				}
			}
			dp = append(dp, item)
			for i := 0; i < 4; i++ {
				b := []byte(item)
				b[i]--
				if b[i] < '0' {
					b[i] = '9'
				}
				new = append(new, string(b))

				b = []byte(item)
				b[i]++
				if b[i] > '9' {
					b[i] = '0'
				}
				new = append(new, string(b))
			}
		}

		pc = pt
		pt = new
		step++
	}

	return -1
}
