package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Arjun")

	got := buffer.String()
	want := "Hello, Arjun"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
