package main

// Runtime: 8 ms, faster than 85.09% of Go online submissions for Find First and Last Position of Element in Sorted Array.
// Memory Usage: 4.1 MB, less than 100.00% of Go online submissions for Find First and Last Position of Element in Sorted Array.
func searchRange(nums []int, target int) []int {
	l1, r1 := 0, len(nums)-1
	if r1 < 0 {
		return []int{-1, -1}
	}
	l2, r2 := l1, r1
	for {
		b1, b2 := l1 <= r1, l2 <= r2
		if !(b1 || b2) {
			break
		}
		if b1 {
			mid := (l1 + r1) / 2 // 0 <= nums.length <= 10**5
			if nums[mid] >= target {
				r1 = mid - 1
			} else if nums[mid] < target {
				l1 = mid + 1
			}
		}
		if b2 {
			mid := (l2 + r2) / 2 // 0 <= nums.length <= 10**5
			if nums[mid] <= target {
				l2 = mid + 1
			} else {
				r2 = mid - 1
			}
		}
	}
	if l1 >= len(nums) || nums[l1] != target {
		return []int{-1, -1}
	}
	return []int{l1, r2}
}
