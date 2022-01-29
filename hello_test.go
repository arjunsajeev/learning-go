package main

import "testing"

func TestHello(t *testing.T) {

	assert := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Arjun")
		want := "Hello, Arjun"

		assert(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assert(t, got, want)
	})

}
