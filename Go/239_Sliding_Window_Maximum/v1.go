package main

// Runtime: 500 ms, faster than 37.72% of Go online submissions for Sliding Window Maximum.
// Memory Usage: 8.6 MB, less than 98.26% of Go online submissions for Sliding Window Maximum.
// https://leetcode.com/submissions/detail/448547718/

func maxSlidingWindowV1(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	q := make([]int, k)
	qHead, qIndex := 0, 0
	for index, num := range nums {
		for qIndex-qHead > 0 && q[(qHead%k+qIndex-qHead-1)%k] < num {
			qIndex--
		}
		head := qHead % k
		q[(head+qIndex-qHead)%k] = num
		qIndex++
		if index < k-1 {
			continue
		}
		res = append(res, q[head])
		if q[head] == nums[index-k+1] {
			qHead++
		}
	}
	return res
}
