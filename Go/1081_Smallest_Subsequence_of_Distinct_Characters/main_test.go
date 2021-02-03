package main

import "testing"

func TestSmallestSubsequence(t *testing.T) {
	for input, output := range map[string]string{
		"bcabc":    "abc",
		"cbacdcbc": "acdb",
		"bbcaac":   "bac",
	} {
		if value := smallestSubsequence(input); value != output {
			t.Errorf("Expected smallestSubsequence(%s) == %s, But got %s", input, output, value)
		}
	}
}
