package main

import "testing"

type inputData struct {
	deadends []string
	target   string
}

var testCase = map[*inputData]int{
	{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202"}:                         6,
	{[]string{"8888"}, "0009"}:                                                         1,
	{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"}: -1,
	{[]string{"0000"}, "8888"}:                                                         -1,
}

func TestOpenLockBFS(t *testing.T) {
	for input, output := range testCase {
		if value := openLockBFS(input.deadends, input.target); value != output {
			t.Errorf("Expected openLockBFS(%v, %s) == %d, But got %d", input.deadends, input.target, output, value)
		}
	}
}

func TestOpenLockDoubleBFS(t *testing.T) {
	for input, output := range testCase {
		if value := openLockDoubleBFS(input.deadends, input.target); value != output {
			t.Errorf("Expected openLockDoubleBFS(%v, %s) == %d, But got %d", input.deadends, input.target, output, value)
		}
	}
}
