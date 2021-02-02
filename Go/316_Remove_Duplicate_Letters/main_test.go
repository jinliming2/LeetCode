package main

import "testing"

func TestRemoveDuplicateLetters(t *testing.T) {
	for input, output := range map[string]string{
		"bcabc":    "abc",
		"cbacdcbc": "acdb",
		"bbcaac":   "bac",
	} {
		if value := removeDuplicateLetters(input); value != output {
			t.Errorf("Expected removeDuplicateLetters(%s) == %s, But got %s", input, output, value)
		}
	}
}
