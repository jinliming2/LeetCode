package main

import "sort"

type intSorter [][]int

func (is intSorter) Len() int {
	return len(is)
}

func (is intSorter) Less(i, j int) bool {
	return is[i][0] < is[j][0]
}

func (is intSorter) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// Runtime: 8 ms, faster than 90.17% of Go online submissions for Merge Intervals.
// Memory Usage: 4.7 MB, less than 89.66% of Go online submissions for Merge Intervals.
func merge(intervals [][]int) [][]int {
	sort.Sort(intSorter(intervals))

	results := make([][]int, 0, len(intervals)) // results := [][]int{}

	left, right := intervals[0][0], intervals[0][1]
	for _, interval := range intervals[1:] {
		if interval[0] > right {
			results = append(results, []int{left, right})
			left, right = interval[0], interval[1]
			continue
		}
		if interval[1] > right {
			right = interval[1]
		}
	}
	results = append(results, []int{left, right})

	return results
}
