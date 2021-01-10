package main

import "testing"

var testCase = map[*ListNode]bool{
	{1, &ListNode{2, nil}}:                             false,
	{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}: true,
}

func TestIsPalindromeV1(t *testing.T) {
	for input, output := range testCase {
		if value := isPalindromeV1(input); value != output {
			var arr []int
			for input != nil {
				arr = append(arr, input.Val)
				input = input.Next
			}
			t.Errorf("Expected isPalindromeV1(%v) == %t, But got %t", arr, output, value)
			break
		}
	}
}

func TestIsPalindromeV2(t *testing.T) {
	for input, output := range testCase {
		if value := isPalindromeV2(input); value != output {
			var arr []int
			for input != nil {
				arr = append(arr, input.Val)
				input = input.Next
			}
			t.Errorf("Expected isPalindromeV2(%v) == %t, But got %t", arr, output, value)
			break
		}
	}
}
