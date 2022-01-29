package main

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected %d received %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
