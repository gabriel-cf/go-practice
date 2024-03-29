package selector

import (
	"time"
	"fmt"
	"net/http"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(a string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(a)
		ch <- true
	}()
	return ch
}