package main

import "testing"

var testCase = map[string]int{
	"abcabcbb": 3,
	"bbbbb":    1,
	"pwwkew":   3,
	"":         0,
	" ":        1,
}

func TestLengthOfLongestSubstringV1(t *testing.T) {
	for input, output := range testCase {
		if value := lengthOfLongestSubstringV1(input); value != output {
			t.Errorf("Expected lengthOfLongestSubstringV1(%s) == %d, But got %d", input, output, value)
		}
	}
}

func TestLengthOfLongestSubstringV2(t *testing.T) {
	for input, output := range testCase {
		if value := lengthOfLongestSubstringV2(input); value != output {
			t.Errorf("Expected lengthOfLongestSubstringV2(%s) == %d, But got %d", input, output, value)
		}
	}
}
