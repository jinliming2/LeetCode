package main

import "testing"

func TestEquationsPossible(t *testing.T) {
	for input, output := range map[*[]string]bool{
		{"a==b", "b!=a"}:         false,
		{"b==a", "a==b"}:         true,
		{"a==b", "b==c", "a==c"}: true,
		{"a==b", "b!=c", "c==a"}: false,
		{"c==c", "b==d", "x!=z"}: true,
	} {
		if value := equationsPossible(*input); value != output {
			t.Errorf("Expected equationsPossible(%v) == %t, But got %t", input, output, value)
		}
	}
}
