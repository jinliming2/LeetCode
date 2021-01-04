package main

import "sort"

type intSorter [][]int

func (is intSorter) Len() int {
	return len(is)
}

func (is intSorter) Less(i, j int) bool {
	if is[i][0] != is[j][0] {
		return is[i][0] < is[j][0]
	}
	return is[i][1] > is[j][1]
}

func (is intSorter) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// Runtime: 12 ms, faster than 100.00% of Go online submissions for Remove Covered Intervals.
// Memory Usage: 5.5 MB, less than 100.00% of Go online submissions for Remove Covered Intervals.
func removeCoveredIntervals(intervals [][]int) int {
	sort.Sort(intSorter(intervals))
	res := 0
	left, right := intervals[0][0], intervals[0][1]
	for _, interval := range intervals[1:] {
		if left <= interval[0] && interval[1] <= right {
			res++
		} else {
			left, right = interval[0], interval[1]
		}
	}
	return len(intervals) - res
}
