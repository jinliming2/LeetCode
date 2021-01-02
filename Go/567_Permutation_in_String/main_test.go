package main

import "testing"

var testCase = map[[2]string]bool{
	{"ab", "eidbaooo"}: true,
	{"ab", "eidboaoo"}: false,
}

func TestCheckInclusionV1(t *testing.T) {
	for input, output := range testCase {
		if value := checkInclusionV1(input[0], input[1]); value != output {
			t.Errorf("Expected checkInclusionV1(%s, %s) == %t, But got %t", input[0], input[1], output, value)
		}
	}
}

func TestCheckInclusionV2(t *testing.T) {
	for input, output := range testCase {
		if value := checkInclusionV2(input[0], input[1]); value != output {
			t.Errorf("Expected checkInclusionV2(%s, %s) == %t, But got %t", input[0], input[1], output, value)
		}
	}
}
