package main

import "testing"

func TestFindAnagrams(t *testing.T) {
	for input, output := range map[[2]string][]int{
		{"cbaebabacd", "abc"}: {0, 6},
		{"abab", "ab"}:        {0, 1, 2},
	} {
		value := findAnagrams(input[0], input[1])
		if len(value) != len(output) {
			t.Errorf("Expected len(findAnagrams(%s, %s)) == %d , But got %d", input[0], input[1], len(output), len(value))
			continue
		}
	aLoop:
		for _, n := range output {
			for _, s := range value {
				if n == s {
					continue aLoop
				}
			}
			t.Errorf("Expected findAnagrams(%s, %s) == %v, But got %v", input[0], input[1], output, value)
			break
		}
	}
}
