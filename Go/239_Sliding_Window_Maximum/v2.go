package main

// Runtime: 352 ms, faster than 41.69% of Go online submissions for Sliding Window Maximum.
// Memory Usage: 9.7 MB, less than 56.58% of Go online submissions for Sliding Window Maximum.
// https://leetcode.com/submissions/detail/448552852/

func maxSlidingWindowV2(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	q := []int{}
	qIndex := 0
	for index, num := range nums {
		for qIndex > 0 && q[qIndex-1] < num {
			qIndex--
			q = q[:qIndex]
		}
		q = append(q, num)
		qIndex++
		if index < k-1 {
			continue
		}
		res = append(res, q[0])
		if q[0] == nums[index-k+1] {
			q = q[1:]
			qIndex--
		}
	}
	return res
}
