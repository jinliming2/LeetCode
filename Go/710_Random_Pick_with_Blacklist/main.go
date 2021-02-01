package main

import "math/rand"

// Runtime: 168 ms, faster than 50.00% of Go online submissions for Random Pick with Blacklist.
// Memory Usage: 13.6 MB, less than 75.00% of Go online submissions for Random Pick with Blacklist.
// https://leetcode.com/submissions/detail/450618716/

// Solution :
type Solution struct {
	n        int
	blackMap map[int]int
}

// Constructor :
func Constructor(N int, blacklist []int) Solution {
	blackMap := make(map[int]int, len(blacklist))
	size, last := N-len(blacklist), N-1
	for _, i := range blacklist {
		if i > size {
			blackMap[i] = -1
		}
	}
	for _, i := range blacklist {
		if i >= size {
			continue
		}
		for _, ok := blackMap[last]; ok; _, ok = blackMap[last] {
			last--
		}
		blackMap[i] = last
		last--
	}
	return Solution{size, blackMap}
}

// Pick :
func (s *Solution) Pick() int {
	v := rand.Intn(s.n)
	if n, ok := s.blackMap[v]; ok {
		return n
	}
	return v
}
