package main

import "testing"

func TestReverseString(t *testing.T) {
	for input, output := range map[*[]byte][]byte{
		{'h', 'e', 'l', 'l', 'o'}:      {'o', 'l', 'l', 'e', 'h'},
		{'H', 'a', 'n', 'n', 'a', 'h'}: {'h', 'a', 'n', 'n', 'a', 'H'},
	} {
		reverseString(*input)
		if len(*input) != len(output) {
			t.Errorf("Expected len(reverseString(%v)) == %d , But got %d", input, len(output), len(*input))
			continue
		}
		for index, item := range output {
			if (*input)[index] != item {
				t.Errorf("Expected reverseString(%v) == %v, But got %v", input, output, input)
				break
			}
		}
	}
}
