package syncing

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()
		assertCounter(t, counter, 3)
	})

	t.Run("It should work safely concurrently", func(t *testing.T) {
		want := 1000
		counter := NewCounter()
		var waitGroup sync.WaitGroup
		waitGroup.Add(want)

		for i := 0; i < want; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Increment()
				wg.Done()
			}(&waitGroup)
		}
		waitGroup.Wait()
		assertCounter(t, counter, want)
	})
}

func assertCounter(t *testing.T, counter *Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("Wanted %d, got %d", want, counter.Value())
	}
}
