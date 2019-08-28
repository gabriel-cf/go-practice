package main

import "testing"

func TestHello(t *testing.T) {
	
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Gabriel")
		want := "Hello Gabriel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saw Hello World when no name is passed", func(t *testing.T){
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

}