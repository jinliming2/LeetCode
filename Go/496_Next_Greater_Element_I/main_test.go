package main

import "testing"

type inputData struct {
	nums1, nums2 []int
}

func TestNextGreaterElement(t *testing.T) {
	for input, output := range map[*inputData][]int{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}}: {-1, 3, -1},
		{[]int{2, 4}, []int{1, 2, 3, 4}}:    {3, -1},
	} {
		value := nextGreaterElement(input.nums1, input.nums2)
		if len(value) != len(output) {
			t.Errorf("Expected len(nextGreaterElement(%v, %v)) == %d , But got %d", input.nums1, input.nums2, len(output), len(value))
			continue
		}
		for index, item := range output {
			if value[index] != item {
				t.Errorf("Expected nextGreaterElement(%v, %v) == %v, But got %v", input.nums1, input.nums2, output, value)
				break
			}
		}
	}
}
