package main

import "testing"

func TestRepeat(t *testing.T) {
	output := Repeat("a", 3)
	expected := "aaa"

	if output != expected {
		t.Errorf("Expected %s but received %s", expected, output)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
