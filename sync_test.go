package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter twice works", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		assertCounter(t, &counter, 2)
	})

	t.Run("It runs safely concurrently", func(t *testing.T) {
		count := 1000
		counter := Counter{}
		wg := sync.WaitGroup{}
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}

		wg.Wait()
		assertCounter(t, &counter, count)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
