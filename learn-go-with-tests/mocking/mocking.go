package main

import (
	"time"
	"fmt"
	"os"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3
const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

type CountdownOperationsSpy struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(w, finalWord)
}

func main() {
	sleeper := ConfigurableSleeper{duration: 5 *time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, &sleeper)
}