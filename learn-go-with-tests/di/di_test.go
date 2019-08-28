package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gabriel")

	got := buffer.String()
	want := "Hello, Gabriel"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}