package main

// Runtime: 4 ms, faster than 85.58% of Go online submissions for Next Greater Element I.
// Memory Usage: 3.1 MB, less than 85.58% of Go online submissions for Next Greater Element I.
// https://leetcode.com/submissions/detail/448093927/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	index2 := len(nums2) - 1

	cache := make(map[int]int, index2+1)
	stack := make([]int, index2+1)

	stackIndex := 0

	for ; index2 >= 0; index2-- {
		v2 := nums2[index2]
		for stackIndex > 0 && stack[stackIndex-1] <= v2 {
			stackIndex--
		}
		if stackIndex == 0 {
			cache[v2] = -1
		} else {
			cache[v2] = stack[stackIndex-1]
		}
		stack[stackIndex] = v2
		stackIndex++
	}

	for index, num := range nums1 {
		nums1[index] = cache[num]
	}

	return nums1
}
