package main

import "testing"

func TestFib(t *testing.T) {
	for input, output := range map[int]int{
		2: 1,
		3: 2,
		4: 3,
		5: 5,
		6: 8,
		7: 13,
		8: 21,
	} {
		if value := fib(input); value != output {
			t.Errorf("Expected fib(%d) == %d, But got %d", input, output, value)
		}
	}
}
