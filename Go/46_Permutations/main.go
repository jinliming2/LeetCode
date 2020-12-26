package main

var results [][]int

// Runtime: 4 ms, faster than 11.24% of Go online submissions for Permutations.
// Memory Usage: 2.6 MB, less than 63.31% of Go online submissions for Permutations.
func permute(nums []int) [][]int {
	length := len(nums)
	count := 1
	for i := 1; i <= length; i++ {
		count *= i
	}
	results = make([][]int, 0, count)
	sub(nums, make([]int, len(nums)), 0)
	return results
}

func sub(s []int, path []int, index int) {
	if len(s) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		results = append(results, p)
		return
	}

	s1 := make([]int, len(s)-1)
	for i, v := range s {
		path[index] = v
		copy(s1, s[:i])
		copy(s1[i:], s[i+1:])
		sub(s1, path, index+1)
	}
}
