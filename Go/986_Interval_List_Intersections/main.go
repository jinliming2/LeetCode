package main

// Runtime: 20 ms, faster than 80.83% of Go online submissions for Interval List Intersections.
// Memory Usage: 6.3 MB, less than 10.83% of Go online submissions for Interval List Intersections.
// https://leetcode.com/submissions/detail/438485014/
func intervalIntersection(A [][]int, B [][]int) [][]int {
	results := [][]int{}

	lengthA, lengthB := len(A), len(B)

	for i, j := 0, 0; i < lengthA && j < lengthB; {
		A0, A1 := A[i][0], A[i][1]
		B0, B1 := B[j][0], B[j][1]
		if A0 > B1 {
		} else if A0 >= B0 {
			if A1 > B1 {
				results = append(results, []int{A0, B1})
			} else {
				results = append(results, []int{A0, A1})
			}
		} else if A1 >= B0 {
			if A1 > B1 {
				results = append(results, []int{B0, B1})
			} else {
				results = append(results, []int{B0, A1})
			}
		}
		if A1 < B1 {
			i++
		} else {
			j++
		}
	}

	return results
}
