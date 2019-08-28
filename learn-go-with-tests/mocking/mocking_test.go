package main

import (
	"time"
	"reflect"
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("Test correct output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &CountdownOperationsSpy{}
		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("Test correct order of commands", func(t *testing.T) {
		countdownOperationsSpy := &CountdownOperationsSpy{}
		Countdown(countdownOperationsSpy, countdownOperationsSpy)

		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(countdownOperationsSpy.Calls, want) {
			t.Errorf("Got %v, want %v", countdownOperationsSpy.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := SpyTime{}
	configurableSleeper := ConfigurableSleeper{duration: sleepTime, sleep: spyTime.Sleep}
	configurableSleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("Should have slept for %v, slept for %v instead", sleepTime, spyTime.durationSlept)
	}
}